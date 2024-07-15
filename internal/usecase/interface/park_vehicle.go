package interfaces

import (
	domain "parking-lot-service/internal/Domain"
	"parking-lot-service/internal/models"
	"time"
)

type ParkVehicleUseCase interface {
	ParkVehicle(parkReq models.ParkReq) (*models.Ticket, error)
	ParkExit(ticketID int, exitTime time.Time) (*domain.Receipt, error)
}
