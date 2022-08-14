package api

import (
	"pencethren/go-messageboard/controllers"
	"pencethren/go-messageboard/entities"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(router *gin.Engine, pingRepo entities.IPingRepository) {
	pings := router.Group("pings")
	{
		pingController := controllers.NewPingsController(pingRepo)
		pingController.ApplyRoutes(pings)
	}
}
