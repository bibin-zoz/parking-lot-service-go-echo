package models

import (
	"time"
)

type ParkingLot struct {
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

// type Vehicle struct {
// 	ID           string `gorm:"primaryKey"`
// 	Type         VehicleType
// 	EntryTime    time.Time
// 	ExitTime     *time.Time
// 	ParkingLotID uint
// }

type Ticket struct {
	ID           string `gorm:"primaryKey"`
	VehicleID    string
	ParkingLotID uint
	EntryTime    time.Time
}
