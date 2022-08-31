package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"pencethren/go-messageboard/entities"
	"pencethren/go-messageboard/operations"
	"pencethren/go-messageboard/repositories"
	"pencethren/go-messageboard/util"

	"github.com/gin-gonic/gin"
)

type BoardController struct {
	ops    operations.BoardOperations
	logger *util.ApplicationLogger
}

type CreateBoardRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewBoardsController(repository repositories.IBoardRepository) BoardController {
	return BoardController{ops: operations.NewBoardOperations(repository), logger: util.NewLogger()}
}

func (bc *BoardController) PostBoard(ctx *gin.Context) {
	var json CreateBoardRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to parse request body",
		})
		bc.logger.Error().Err(err).Msg("failed to parse request body")
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
		"message": result,
	})
}

func (bc *BoardController) ApplyRoutes(router gin.IRoutes) {
	router.POST("/", bc.PostBoard)
}
