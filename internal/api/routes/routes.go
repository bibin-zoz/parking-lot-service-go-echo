package routes

import (
	"net/http"
	handlers "parking-lot-service/internal/api/handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, parkingLotHandler *handlers.ParkingLotHandler, parkVehicleHandler *handlers.ParkVehicleHandler) {
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "ci-cd test")
	})

	// Parking Lots
	e.POST("/parking-lots", parkingLotHandler.CreateParkingLot)
	e.GET("/parking-lots/:id", parkingLotHandler.GetParkingLotByID)
	e.GET("/parking-lots", parkingLotHandler.GetAllParkingLots)
	e.PUT("/parking-lots/:id", parkingLotHandler.UpdateParkingLot)
	e.DELETE("/parking-lots/:id", parkingLotHandler.DeleteParkingLot)

	// Free Slots
	e.GET("/parking-lots/free-slots/:parkingLotID", parkingLotHandler.GetFreeSlots)

	// Parking Vehicle
	e.GET("/vehicle-types", parkVehicleHandler.GetVehicleTypes)
	e.POST("/park-vehicle", parkVehicleHandler.ParkVehicle)
	e.DELETE("/park-vehicle", parkVehicleHandler.ParkExit)
}
