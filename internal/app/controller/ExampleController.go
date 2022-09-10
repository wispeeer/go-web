package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ExampleController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    1,
		"result":  time.Now(),
		"message": "this is example controller",
	})
}
