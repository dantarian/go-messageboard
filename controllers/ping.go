package controllers

import (
	"net/http"
	"pencethren/go-messageboard/entities"
	"pencethren/go-messageboard/operations"

	"github.com/gin-gonic/gin"
)

type PingsController struct {
	repository entities.IPingRepository
}

func NewPingsController(repository entities.IPingRepository) PingsController {
	return PingsController{repository: repository}
}

func (p *PingsController) PostPing(ctx *gin.Context) {
	result := operations.RespondToPing(p.repository)
	ctx.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func (p *PingsController) GetPings(ctx *gin.Context) {
	total := operations.CountPingsReceived(p.repository)
	ctx.JSON(http.StatusOK, gin.H{
		"total": total,
	})
}

func (p *PingsController) ApplyRoutes(router gin.IRoutes) {
	router.POST("/", p.PostPing)
	router.GET("/", p.GetPings)
}
