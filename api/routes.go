package api

import (
	"pencethren/go-messageboard/api/ping"

	"github.com/gin-gonic/gin"
)

type RouteBuilder func(*gin.Engine)

func Build(router *gin.Engine) {
	ping.Build(router)
}
