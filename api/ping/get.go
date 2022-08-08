package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
