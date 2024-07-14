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

type VehicleType struct {
	ID          uint   `gorm:"primaryKey"`
	VehicleType string `json:"vehicle_type"`
}
type Ticket struct {
	ID            string      `gorm:"primaryKey"`
	VehicleTypeID uint        `json:"vehicle_type_id"`
	VehicleType   VehicleType `gorm:"foreignKey:VehicleTypeID"`
	VehicleNumber string      `json:"vehicle_number"`
	ParkingLotID  uint        `json:"parking_lot_id"`
	EntryTime     time.Time   `json:"entry_time"`
	ExitTime      *time.Time  `json:"exit_time,omitempty"`
	Status        string      `gorm:"type:enum('parked', 'exited');default:'parked'" json:"status"`
}

type Receipt struct {
	ID           string    `gorm:"primaryKey"`
	VehicleType  string    `json:"vehicle_type"`
	ParkingLotID uint      `json:"parking_lot_id"`
	EntryTime    time.Time `json:"entry_time"`
	ExitTime     time.Time `json:"exit_time"`
	Rate         float64   `json:"rate"`
	RateType     string    `gorm:"type:enum('hourly', 'daily')"`
	BillAmount   float64   `json:"bill_amount"`
}

type FreeSlots struct {
	TwoWheel      int `json:"two_wheel"`
	FourWheel     int `json:"four_wheel"`
	HeavyVehicles int `json:"heavy_vehicles"`
}
