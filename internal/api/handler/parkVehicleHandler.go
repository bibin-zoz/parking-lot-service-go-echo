package handlers

import (
	"net/http"
	"parking-lot-service/internal/models"
	usecase "parking-lot-service/internal/usecase/interface"

	"github.com/labstack/echo/v4"
)

type ParkVehicleHandler struct {
	parkingVehicleUseCase usecase.ParkVehicleUseCase
}

func NewParkVehicleHandler(ParkVehicleUseCase usecase.ParkVehicleUseCase) *ParkVehicleHandler {
	return &ParkVehicleHandler{parkingVehicleUseCase: ParkVehicleUseCase}
}

func (h *ParkVehicleHandler) ParkVehicle(c echo.Context) error {
	req := new(models.ParkReq)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Call use case to park vehicle
	ticket, err := h.parkingVehicleUseCase.ParkVehicle(*req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, ticket)
}

// func (h *ParkVehicleHandler) ParkExit(c echo.Context) error {
// 	var request struct {
// 		TicketID string `json:"ticket_id"`
// 		ExitTime string `json:"exit_time"`
// 	}

// 	if err := c.Bind(&request); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
// 	}

// 	exitTime, err := time.Parse(time.RFC3339, request.ExitTime)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid exit time format"})
// 	}

// 	// Call use case to handle parking exit
// 	if err := h.parkingVehicleUseCase.ParkExit(request.TicketID, exitTime); err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, map[string]string{"message": "vehicle exited successfully"})
// }
