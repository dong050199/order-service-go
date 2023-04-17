package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"order-service/di"
	"order-service/pkg/config"
	"order-service/pkg/errormap"
	"order-service/pkg/ginutils"
	"order-service/pkg/infra"
	"order-service/pkg/middleware"
	"order-service/pkg/swagger"
	"order-service/pkg/tracing"
	"order-service/pkg/tracingfx"
	"order-service/service/http/router"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func registerSwaggerHandler(g *gin.Engine) {
	swaggerAPI := g.Group("/swagger")
	swag := swagger.NewSwagger()
	swaggerAPI.Use(swag.SwaggerHandler(false))
	swag.Register(swaggerAPI)
}

func startServer(ginEngine *gin.Engine, lifecycle fx.Lifecycle) {
	port := viper.GetString("PORT")
	server := http.Server{
		Addr:    ":" + port,
		Handler: ginEngine,
	}
	ginEngine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("run on port:", port)
			go func() {
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					fmt.Errorf("failed to listen and serve from server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}

func registerService(ginEngine *gin.Engine,
	userRouter router.UserRouter,
	productRouter router.ProductRouter,
	cartRouter router.CartRouter,
	tracer tracing.Tracer,
) {
	ginEngine.Use(middleware.Recover())
	ginEngine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "OPTIONS", "GET"},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	gGroup := ginEngine.Group("api/v1")
	userRouter.Register(gGroup)
	productRouter.Register(gGroup)
	cartRouter.Register(gGroup)
	gGroup.Use(ginutils.InjectTraceID, tracer.TracingHandler)
}

func main() {
	fx.New(
		fx.Invoke(errormap.Initialize),
		fx.Invoke(config.InitConfig),
		fx.Invoke(config.SetConfig),
		fx.Invoke(infra.InitMySQL),
		di.Module,
		tracingfx.Module,
		fx.Provide(gin.Default),
		fx.Invoke(
			registerSwaggerHandler,
			registerService,
			startServer,
		),
	).Run()
}
