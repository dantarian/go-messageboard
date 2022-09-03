package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"pencethren/go-messageboard/entities"
	"pencethren/go-messageboard/operations"
	"pencethren/go-messageboard/util"

	"github.com/gin-gonic/gin"
)

type BoardController struct {
	ops    operations.IBoardOperations
	logger *util.ApplicationLogger
}

type CreateBoardRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func NewBoardController(ops operations.IBoardOperations) BoardController {
	return BoardController{ops: ops, logger: util.NewLogger()}
}

func (bc *BoardController) PostBoard(ctx *gin.Context) {
	var json CreateBoardRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		bc.logger.Error().Err(err).Msg("failed to parse request body")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "failed to parse request body",
		})
		return
	}

	result, err := bc.ops.CreateBoard(json.Name, json.Description)
	if err != nil {
		var validationError *entities.ValidationError
		var businessRuleError *operations.BusinessRuleError
		switch {
		case errors.As(err, &validationError):
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": fmt.Sprintf("%v", validationError),
			})
		case errors.As(err, &businessRuleError):
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": fmt.Sprintf("%v", businessRuleError),
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": result,
	})
}

func (bc *BoardController) ApplyRoutes(router gin.IRoutes) {
	router.POST("/", bc.PostBoard)
}
