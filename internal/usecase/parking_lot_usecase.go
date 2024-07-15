package usecase

import (
	"errors"
	"fmt"

	"parking-lot-service/internal/models"
	repository "parking-lot-service/internal/repository/interface"
	usecase "parking-lot-service/internal/usecase/interface"

	"github.com/rs/zerolog/log"
)

type parkingLotUseCase struct {
	parkingLotRepo repository.ParkingLotRepository
}

func NewParkingLotUseCase(parkingLotRepo repository.ParkingLotRepository) usecase.ParkingLotUseCase {
	return &parkingLotUseCase{
		parkingLotRepo: parkingLotRepo,
	}
}

func (u *parkingLotUseCase) GetAllParkingLots() ([]models.ParkingLot, error) {
	lots, err := u.parkingLotRepo.GetAllParkingLots()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get all parking lots")
		return nil, err
	}

	if len(lots) == 0 {
		return nil, fmt.Errorf("no parking lots available")
	}

	return lots, nil
}

func (u *parkingLotUseCase) GetParkingLotByID(id uint) (*models.ParkingLot, error) {
	lot, err := u.parkingLotRepo.GetParkingLotByID(id)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get parking lot by ID: %d", id)
		return nil, err
	}
	if lot == nil {
		err = errors.New("parking lot not found")
		log.Warn().Msgf("Parking lot with ID: %d not found", id)
		return nil, err
	}
	return lot, nil
}

func (u *parkingLotUseCase) CreateParkingLot(parkingLot *models.ParkingLot) error {

	err := u.parkingLotRepo.CreateParkingLot(parkingLot)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create parking lot")
		return err
	}

	log.Info().Msg("Parking lot created successfully")
	return nil
}

func (u *parkingLotUseCase) UpdateParkingLot(parkingLot *models.ParkingLot) error {

	err := u.parkingLotRepo.UpdateParkingLot(parkingLot)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update parking lot")
		return err
	}

	log.Info().Msgf("Parking lot with ID: %d updated successfully", parkingLot.ID)
	return nil
}

func (u *parkingLotUseCase) DeleteParkingLot(id uint) error {
	err := u.parkingLotRepo.DeleteParkingLot(id)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to delete parking lot with ID: %d", id)
		return err
	}

	log.Info().Msgf("Parking lot with ID: %d deleted successfully", id)
	return nil
}
func (u *parkingLotUseCase) GetFreeParkingLots(parkingLotID uint) (models.FreeSlots, error) {
	freeSlots := models.FreeSlots{}

	// Get parking lot details
	parkingLot, err := u.parkingLotRepo.GetParkingLotByID(parkingLotID)
	if err != nil {
		return freeSlots, err
	}

	// Get vehicle counts by type
	counts, err := u.parkingLotRepo.GetVehicleCountsByType(parkingLotID)
	if err != nil {
		return freeSlots, err
	}

	// Calculate free slots
	freeSlots.TwoWheel = parkingLot.MotorcycleSpots - counts[1] // Assuming 1 corresponds to TwoWheel
	freeSlots.FourWheel = parkingLot.CarSpots - counts[2]       // Assuming 2 corresponds to FourWheel
	freeSlots.HeavyVehicles = parkingLot.BusSpots - counts[3]   // Assuming 3 corresponds to HeavyVehicles

	return freeSlots, nil
}
