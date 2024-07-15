package interfaces

import (
	"parking-lot-service/internal/models"
)

type ParkVehicleUseCase interface {
	ParkVehicle(parkReq models.ParkReq) (*models.Ticket, error)
}
