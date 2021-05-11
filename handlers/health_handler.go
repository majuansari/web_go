package handlers

import (
	"net/http"
	"web/app"
	error2 "web/pkg/error"

	"github.com/labstack/echo/v4"
)

// HealthHandler ..
type HealthHandler struct {
	server *app.App
}

// NewHealthHandler ..
func NewHealthHandler(server *app.App) HealthHandler {
	return HealthHandler{server: server}
}

// Health ...
func (handler HealthHandler) Health(c echo.Context) error {
	return error2.NewRequestError("Some error", 500, "somecode")
	return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

	c.Logger().Info("Request reached health handler")
	return c.String(200, "All good")
}
