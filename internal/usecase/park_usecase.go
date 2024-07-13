package usecase

import (
	repository "parking-lot-service/internal/repository/interface"
	usecase "parking-lot-service/internal/usecase/interface"
)

type parkUseCase struct {
	parkingLotRepo repository.ParkingLotRepository
}

func NewParkUseCase(parkingLotRepo repository.ParkingLotRepository) usecase.ParkUseCase {
	return &parkUseCase{
		parkingLotRepo: parkingLotRepo,
	}

}
