package interfaces

import (
	domain "parking-lot-service/internal/Domain"
	"time"
)

type ParkVehicleRepository interface {
	ParkVehicle(vehicleType string, vehicleNumber string, parkingLotID uint) (*domain.Ticket, error)
	GenerateReceipt(ticketID string, exitTime time.Time) (*domain.Receipt, error)
	ParkExit(string, time.Time) (*domain.Receipt, error)
	GetVehicleDetails(vehicleID uint) (*domain.VehicleType, error)
	
}
