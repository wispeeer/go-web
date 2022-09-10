package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/wispeeer/go-web/internal/app/router/v1"
	"github.com/wispeeer/go-web/internal/pkg/controller"
)

// define root routes
func rootRoutes(engine *gin.Engine) {
	apiRoutes := engine.Group("/")
	apiRoutes.GET("health", controller.HealthChecks)
}

// registry all routes
func SetRouter(engine *gin.Engine) {
	rootRoutes(engine)

	// api v1
	v1.ExampleRoutes(engine)
}
