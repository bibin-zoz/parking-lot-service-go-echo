package domain

import (
	"gorm.io/gorm"
)

type ParkingLot struct {
	gorm.Model
	ID               uint    `gorm:"primaryKey"`
	Name             string  `json:"name"`
	MotorcycleSpots  int     `json:"motorcycle_spots"`
	CarSpots         int     `json:"car_spots"`
	BusSpots         int     `json:"bus_spots"`
	MotorcycleTariff float64 `json:"motorcycle_tariff"`
	CarTariff        float64 `json:"car_tariff"`
	BusTariffDaily   float64 `json:"bus_tariff_daily"`
	BusTariffHourly  float64 `json:"bus_tariff_hourly"`
}
