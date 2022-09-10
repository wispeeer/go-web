package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wispeeer/go-web/internal/app/controller"
)

// define example routes
func ExampleRoutes(engine *gin.Engine) {
	apiRoutes := engine.Group("/api/v1")
	apiRoutes.GET("example", controller.ExampleController)
}
