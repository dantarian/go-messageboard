package controllers

import (
	"net/http"
	"pencethren/go-messageboard/operations"
	"pencethren/go-messageboard/repositories"

	"github.com/gin-gonic/gin"
)

type PingsController struct {
	repository repositories.IPingRepository
}

func NewPingsController(repository repositories.IPingRepository) PingsController {
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
