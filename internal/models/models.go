package models

import (
	"time"
)

// @description ParkingLot holds details about the parking lot such as available spots and tariffs
// ParkingLot represents a parking lot entity
type ParkingLot struct {
	ID               uint    `gorm:"primaryKey" json:"id" validate:"nonZeroPositive"`
	Name             string  `json:"name" validate:"nameLength"`
	MotorcycleSpots  int     `json:"motorcycle_spots" validate:"parkingSpots"`
	CarSpots         int     `json:"car_spots" validate:"parkingSpots"`
	BusSpots         int     `json:"bus_spots" validate:"parkingSpots"`
	MotorcycleTariff float64 `json:"motorcycle_tariff"`
	CarTariff        float64 `json:"car_tariff"`
	BusTariffDaily   float64 `json:"bus_tariff_daily"`
	BusTariffHourly  float64 `json:"bus_tariff_hourly"`
}


// @description VehicleType holds the vehicle type information
type VehicleType struct {
	ID          uint   `gorm:"primaryKey" json:"id" example:"1"`
	VehicleType string `json:"vehicle_type" example:"Car"`
}

// @description Ticket is issued when a vehicle is parked
type Ticket struct {
	ID            int       `gorm:"primaryKey" json:"id" example:"1"`
	VehicleTypeID uint      `json:"vehicle_type_id" example:"1"`
	VehicleType   string    `json:"vehicle_type" example:"Car"`
	VehicleNumber string    `json:"vehicle_number" example:"ABC123"`
	ParkingLotID  uint      `json:"parking_lot_id" example:"1"`
	EntryTime     time.Time `json:"entry_time" example:"2023-07-15T19:20:30+01:00"`
	IsParked      bool      `gorm:"default:true" json:"is_parked" example:"true"`
}


// @description Receipt is issued when a vehicle exits the parking lot
type Receipt struct {
	ID           int       `gorm:"primaryKey" json:"id" example:"1"`
	VehicleType  string    `json:"vehicle_type" example:"Car"`
	ParkingLotID uint      `json:"parking_lot_id" example:"1"`
	EntryTime    time.Time `json:"entry_time" example:"2023-07-15T19:20:30+01:00"`
	ExitTime     time.Time `json:"exit_time" example:"2023-07-15T21:20:30+01:00"`
	Rate         float64   `json:"rate" example:"20.0"`
	RateType     string    `json:"rate_type" example:"Hourly"`
	BillAmount   float64   `json:"bill_amount" example:"40.0"`
}

// @description FreeSlots holds information about the available parking slots for different types of vehicles
type FreeSlots struct {
	TwoWheel      int `json:"two_wheel" example:"5"`
	FourWheel     int `json:"four_wheel" example:"15"`
	HeavyVehicles int `json:"heavy_vehicles" example:"2"`
}
