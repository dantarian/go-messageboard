package ping

import (
	"github.com/gin-gonic/gin"
)

func Build(router *gin.Engine) {
	router.GET("/ping", Get)
}
