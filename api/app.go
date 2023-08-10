package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type APP struct {
	*gin.Engine
	s    *http.Server
	port int
}

func NewAPP(port int) *APP {
	gin.SetMode(gin.DebugMode)

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("%-6s %-s --> %s(%d handlers)\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	app := &APP{
		Engine: gin.Default(),
		port:   port,
	}
	log.Printf("rest server is running on port: %d\n", port)

	address := fmt.Sprintf(":%d", port)

	app.s = &http.Server{
		Addr:    address,
		Handler: app.Engine,
	}

	initRouter(app.Engine)

	return app
}

func (a *APP) Run() {
	err := a.s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
