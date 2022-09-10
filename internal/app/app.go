package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wispeeer/go-web/internal/app/router"
)

func init() {

}

func flags() error {
	var argv = os.Args
	var argc = len(os.Args)

	// 处理 argc
	fmt.Printf("argc: %d - argv: %v", argc, argv)

	return nil
}

func App() error {
	var err error

	// process app flags
	err = flags()
	if err != nil {
		return err
	}

	// connect data source: like database

	// new server & init router
	engine := gin.New()
	engine.SetTrustedProxies(nil)

	router.SetRouter(engine)

	err = engine.Run(":2333")
	if err != nil {
		return err
	}
	return nil
}
