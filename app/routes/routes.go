package routes

import (
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"web/app"
	"web/config"
	"web/handlers"
)

func RegisterRoutes(server *app.App, cfg *config.EnvConfig) {
	healthHandler := handlers.NewHealthHandler(server)
	userHandler := handlers.NewUserHandler(server)
	testHandler := handlers.NewTestHandler(server)
	e := server.Echo
	//pprof.Register(e)
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(otelecho.Middleware(cfg.Tracer.ServiceName))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/health", healthHandler.Health)
	e.POST("/user/create", userHandler.Create)
	e.GET("/user/list", userHandler.List)
	e.GET("/user/index", userHandler.Index)
	e.GET("/user/:id", userHandler.Details)
	e.POST("/user/test/:id", userHandler.RestAPITest)
	e.GET("/test", testHandler.Test)

	e.Static("/", "static")

}
