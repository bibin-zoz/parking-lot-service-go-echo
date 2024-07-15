package handlers

import (
	"net/http"
	"parking-lot-service/internal/models"
	usecase "parking-lot-service/internal/usecase/interface"
	"time"

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

	ticket, err := h.parkingVehicleUseCase.ParkVehicle(*req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, ticket)
}

func (h *ParkVehicleHandler) ParkExit(c echo.Context) error {
	var request struct {
		TicketID int `json:"ticket_id"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if request.TicketID < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ticket id"})
	}

	exitTime := time.Now()

	// Call use case to handle parking exit
	receipt, err := h.parkingVehicleUseCase.ParkExit(request.TicketID, exitTime)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, receipt)
}
