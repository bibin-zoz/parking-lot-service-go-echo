package repository

import (
	repository "parking-lot-service/internal/repository/interface"

	"gorm.io/gorm"
)

type parkingLotRepo struct {
	db *gorm.DB
}

func NewParkingLotRepository(db *gorm.DB) repository.ParkingLotRepository {
	return &parkingLotRepo{db}
}
