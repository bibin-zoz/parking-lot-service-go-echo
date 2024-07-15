package handlers

import (
	"net/http"
	"parking-lot-service/internal/models"
	usecase "parking-lot-service/internal/usecase/interface"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type ParkVehicleHandler struct {
	parkingVehicleUseCase usecase.ParkVehicleUseCase
}

func NewParkVehicleHandler(ParkVehicleUseCase usecase.ParkVehicleUseCase) *ParkVehicleHandler {
	return &ParkVehicleHandler{parkingVehicleUseCase: ParkVehicleUseCase}
}

// ParkVehicle godoc
// @Summary Park a vehicle
// @Description Park a vehicle in the parking lot
// @Tags parking
// @Accept  json
// @Produce  json
// @Param parkReq body models.ParkReq true "Parking Request object"
// @Success 201 {object} models.Ticket
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /park-vehicle [post]
func (h *ParkVehicleHandler) ParkVehicle(c echo.Context) error {
	req := new(models.ParkReq)

	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}

	err := models.ValidateStruct(req)
	if err != nil {
		log.Error().Err(err).Msg("Validation failed for parking request")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error(), "message": "Validation failed for creating parking lot, recheck the data"})
	}

	ticket, err := h.parkingVehicleUseCase.ParkVehicle(*req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to park vehicle")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, ticket)
}

// ParkExit godoc
// @Summary Exit a parked vehicle
// @Description Exit a parked vehicle from the parking lot
// @Tags parking
// @Accept  json
// @Produce  json
// @Param exitRequest body models.ExitRequest true "Exit Request object"
// @Success 200 {object} models.Receipt
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /park-vehicle [delete]
func (h *ParkVehicleHandler) ParkExit(c echo.Context) error {
	var request models.ExitRequest

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
