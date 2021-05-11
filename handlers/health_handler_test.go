package handlers

import (
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
	"web/app"
	"web/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func BenchmarkHealthCheck(b *testing.B) {
	e := echo.New()
	config := config.NewEnvConfig()
	server := app.NewMockServer(config)

	server.Echo.Logger.SetLevel(log.DEBUG)
	server.Echo.Logger.SetHeader("${time_rfc3339} ${level}")

	healthHandler := NewHealthHandler(server)
	echo := server.Echo
	echo.Use(middleware.Logger())
	echo.GET("/health", healthHandler.Health)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/health", nil)

	for i := 0; i < b.N; i++ {
		e.ServeHTTP(w, r)
	}
}
