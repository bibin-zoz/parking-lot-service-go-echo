package interfaces

import (
	"parking-lot-service/internal/models"
	"time"
)

type ParkVehicleUseCase interface {
	ParkVehicle(parkReq models.ParkReq) (*models.Ticket, error)
	ParkExit(ticketID int, exitTime time.Time) (*models.Receipt, error)
}
