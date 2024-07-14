package interfaces

import "parking-lot-service/internal/models"

type ParkingLotRepository interface {
	GetAllParkingLots() ([]models.ParkingLot, error)
	GetParkingLotByID(id uint) (*models.ParkingLot, error)
	CreateParkingLot(parkingLot *models.ParkingLot) error
	UpdateParkingLot(parkingLot *models.ParkingLot) error
	DeleteParkingLot(id uint) error
	GetVehicleCountsByType(parkingLotID uint) (map[uint]int, error)
}
