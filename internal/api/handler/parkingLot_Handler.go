package handlers

import (
	"net/http"
	"parking-lot-service/internal/models"
	usecase "parking-lot-service/internal/usecase/interface"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type ParkingLotHandler struct {
	parkingLotUseCase usecase.ParkingLotUseCase
}

func NewHandler(parkingLotUseCase usecase.ParkingLotUseCase) *ParkingLotHandler {
	return &ParkingLotHandler{parkingLotUseCase: parkingLotUseCase}
}

func (h *ParkingLotHandler) CreateParkingLot(c echo.Context) error {
	var req models.ParkingLot
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("Invalid request format")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request format"})
	}

	err := models.ValidateStruct(&req)
	if err != nil {
		log.Error().Err(err).Msg("Validation failed for creating parking lot")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = h.parkingLotUseCase.CreateParkingLot(&req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create parking lot")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	log.Info().Msg("Parking lot created successfully")
	return c.JSON(http.StatusCreated, req)
}

func (h *ParkingLotHandler) GetParkingLotByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error().Err(err).Msg("Invalid ID format")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID format"})
	}

	lot, err := h.parkingLotUseCase.GetParkingLotByID(uint(id))
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get parking lot by ID: %d", id)
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, lot)
}

func (h *ParkingLotHandler) GetAllParkingLots(c echo.Context) error {
	lots, err := h.parkingLotUseCase.GetAllParkingLots()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get all parking lots")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, lots)
}

func (h *ParkingLotHandler) UpdateParkingLot(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error().Err(err).Msg("Invalid ID format")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID format"})
	}

	var req models.ParkingLot
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("Invalid request format")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request format"})
	}
	req.ID = uint(id)

	err = models.ValidateStruct(&req)
	if err != nil {
		log.Error().Err(err).Msg("Validation failed for updating parking lot")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = h.parkingLotUseCase.UpdateParkingLot(&req)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to update parking lot with ID: %d", id)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	log.Info().Msgf("Parking lot with ID: %d updated successfully", id)
	return c.JSON(http.StatusOK, req)
}

func (h *ParkingLotHandler) DeleteParkingLot(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error().Err(err).Msg("Invalid ID format")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID format"})
	}

	err = h.parkingLotUseCase.DeleteParkingLot(uint(id))
	if err != nil {
		log.Error().Err(err).Msgf("Failed to delete parking lot with ID: %d", id)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	log.Info().Msgf("Parking lot with ID: %d deleted successfully", id)
	return c.NoContent(http.StatusNoContent)
}
func (h *ParkingLotHandler) GetFreeSlots(c echo.Context) error {
	// Extract parkingLotID from route parameter
	parkingLotID, err := strconv.Atoi(c.Param("parkingLotID"))
	if err != nil {
		log.Error().Err(err).Msg("Invalid parking lot ID format")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid parking lot ID format"})
	}

	// Call use case to get free parking slots
	freeSlots, err := h.parkingLotUseCase.GetFreeParkingLots(uint(parkingLotID))
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get free slots for parking lot ID: %d", parkingLotID)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, freeSlots)
}
