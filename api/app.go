package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/yonymo/simplechat/api/config"
	"github.com/yonymo/simplechat/pkg/app"
	"github.com/yonymo/simplechat/pkg/log"
	"net/http"
)

type APP struct {
	*gin.Engine
	s    *http.Server
	port int

	trans ut.Translator
}

func NewAPP(basename string) *app.APP {

	cfg := config.NewConfig()
	app := app.NewApp("api", basename,
		app.WithOptions(cfg),
		app.WithRunFunc(run(cfg)),
		app.WithNoConfig(),
	)

	return app
}

func run(cfg *config.Config) app.RunFunc {
	return func(base string) error {
		gin.SetMode(gin.DebugMode)

		gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
			log.Infof("%-6s %-s --> %s(%d handlers)\n", httpMethod, absolutePath, handlerName, nuHandlers)
		}
		a := &APP{
			Engine: gin.Default(),
			port:   cfg.Server.Port,
		}
		log.Infof("rest server is running on port: %d\n", cfg.Server.Port)

		address := fmt.Sprintf(":%d", cfg.Server.Port)

		a.s = &http.Server{
			Addr:    address,
			Handler: a.Engine,
		}

		err := a.initTrans("zh")
		if err != nil {
			panic(err)
		}
		RegisteMobile(a.trans)

		initRouter(a.Engine, cfg, a.trans)

		err = a.s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
		return nil
	}

}
