package routes

import (
	"net/http"
	handlers "parking-lot-service/internal/api/handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, handler *handlers.Handler) {
	e.GET("/home", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "test")
	})
}
