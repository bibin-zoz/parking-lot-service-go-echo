package routes

import (
	"net/http"
	handlers "parking-lot-service/internal/api/handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, parkingLotHandler *handlers.ParkingLotHandler) {
	e.GET("/home", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "test")
	})
	e.POST("/parking-lots", parkingLotHandler.CreateParkingLot)
	e.GET("/parking-lots/:id", parkingLotHandler.GetParkingLotByID)
	e.GET("/parking-lots", parkingLotHandler.GetAllParkingLots)
	e.PUT("/parking-lots/:id", parkingLotHandler.UpdateParkingLot)
	e.DELETE("/parking-lots/:id", parkingLotHandler.DeleteParkingLot)

	e.GET("/parkinglots/freeslots/:parkingLotID", parkingLotHandler.GetFreeSlots)
}
