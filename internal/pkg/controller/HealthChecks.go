package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wispeeer/go-web/pkg/info"
)

func HealthChecks(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"service":      info.AppName,
			"healthy":      info.SysInfo.Get("healthy"),
			"timestamp":    time.Now().Format(time.ANSIC),
			"dependencies": []gin.H{{}},
		},
	)
}
