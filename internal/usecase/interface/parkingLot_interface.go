package interfaces

import "parking-lot-service/internal/models"

type ParkingLotUseCase interface {
	GetAllParkingLots() ([]models.ParkingLot, error)
	GetParkingLotByID(id uint) (*models.ParkingLot, error)
	CreateParkingLot(parkingLot *models.ParkingLot) error
	UpdateParkingLot(parkingLot *models.ParkingLot) error
	DeleteParkingLot(id uint) error
	GetFreeParkingLots(parkingLotID uint) (models.FreeSlots, error)
}
