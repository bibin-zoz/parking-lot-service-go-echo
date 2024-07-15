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

// CreateParkingLot godoc
// @Summary Create a new parking lot
// @Description Create a new parking lot with specified details
// @Tags parkinglot
// @Accept  json
// @Produce  json
// @Param parkingLot body models.ParkingLot true "Parking Lot object to be created"
// @Success 201 {object} models.ParkingLot
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /parking-lots [post]
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

// GetParkingLotByID godoc
// @Summary Get a parking lot by ID
// @Description Retrieve details of a parking lot identified by its ID
// @Tags parkinglot
// @Produce  json
// @Param id path int true "Parking Lot ID"
// @Success 200 {object} models.ParkingLot
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /parking-lots/{id} [get]
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

// GetAllParkingLots godoc
// @Summary Get all parking lots
// @Description Retrieve all existing parking lots
// @Tags parkinglot
// @Produce  json
// @Success 200 {array} models.ParkingLot
// @Failure 500 {object} map[string]string
// @Router /parking-lots [get]
func (h *ParkingLotHandler) GetAllParkingLots(c echo.Context) error {
	lots, err := h.parkingLotUseCase.GetAllParkingLots()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get all parking lots")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, lots)
}

// UpdateParkingLot godoc
// @Summary Update a parking lot
// @Description Update details of a parking lot identified by its ID
// @Tags parkinglot
// @Accept  json
// @Produce  json
// @Param id path int true "Parking Lot ID"
// @Param parkingLot body models.ParkingLot true "Updated Parking Lot object"
// @Success 200 {object} models.ParkingLot
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /parking-lots/{id} [put]
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

// DeleteParkingLot godoc
// @Summary Delete a parking lot
// @Description Delete a parking lot identified by its ID
// @Tags parkinglot
// @Param id path int true "Parking Lot ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /parking-lots/{id} [delete]
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

// GetFreeSlots godoc
// @Summary Get free slots in a parking lot
// @Description Retrieve the number of available slots in a parking lot identified by its ID
// @Tags parkinglot
// @Produce  json
// @Param parkingLotID path int true "Parking Lot ID"
// @Success 200 {object} models.FreeSlots
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /parking-lots/free-slots/{parkingLotID} [get]
func (h *ParkingLotHandler) GetFreeSlots(c echo.Context) error {
	parkingLotID, err := strconv.Atoi(c.Param("parkingLotID"))
	if err != nil {
		log.Error().Err(err).Msg("Invalid parking lot ID format")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid parking lot ID format"})
	}

	freeSlots, err := h.parkingLotUseCase.GetFreeParkingLots(uint(parkingLotID))
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get free slots for parking lot ID: %d", parkingLotID)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, freeSlots)
}
