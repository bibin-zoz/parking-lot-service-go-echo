package interfaces

import (
	domain "parking-lot-service/internal/Domain"
	"time"
)

type ParkVehicleRepository interface {
	ParkVehicle(ticket *domain.Ticket) (*domain.Ticket, error)
	GenerateReceipt(ticketID string, exitTime time.Time) (*domain.Receipt, error)
	ParkExit(string, time.Time) (*domain.Receipt, error)
	GetVehicleDetails(vehicleID uint) (*domain.VehicleType, error)
	GetTicketDetailsByID(ticketID int) (*domain.Ticket, error)
	SaveExitDetails(ticket *domain.Ticket, receipt *domain.Receipt) (*domain.Receipt, error)
	GetParkingDetailsByVehicleNumber(vehicleNumber string) (*domain.Ticket, error)
}
