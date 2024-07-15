package models

import (
	"time"
)

// ParkingLot represents a parking lot
// @description ParkingLot holds details about the parking lot such as available spots and tariffs
type ParkingLot struct {
	ID               uint    `gorm:"primaryKey" json:"id" example:"1"`
	Name             string  `json:"name" example:"Main Parking Lot"`
	MotorcycleSpots  int     `json:"motorcycle_spots" example:"10"`
	CarSpots         int     `json:"car_spots" example:"20"`
	BusSpots         int     `json:"bus_spots" example:"5"`
	MotorcycleTariff float64 `json:"motorcycle_tariff" example:"10.0"`
	CarTariff        float64 `json:"car_tariff" example:"20.0"`
	BusTariffDaily   float64 `json:"bus_tariff_daily" example:"100.0"`
	BusTariffHourly  float64 `json:"bus_tariff_hourly" example:"15.0"`
}

// VehicleType represents a type of vehicle
// @description VehicleType holds the vehicle type information
type VehicleType struct {
	ID          uint   `gorm:"primaryKey" json:"id" example:"1"`
	VehicleType string `json:"vehicle_type" example:"Car"`
}

// Ticket represents a parking ticket
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

// Receipt represents a receipt for a parking exit
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

// FreeSlots represents the available parking slots
// @description FreeSlots holds information about the available parking slots for different types of vehicles
type FreeSlots struct {
	TwoWheel      int `json:"two_wheel" example:"5"`
	FourWheel     int `json:"four_wheel" example:"15"`
	HeavyVehicles int `json:"heavy_vehicles" example:"2"`
}


