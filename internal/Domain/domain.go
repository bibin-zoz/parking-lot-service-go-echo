package domain

import (
	"time"

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

type VehicleType struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	VehicleType string `json:"vehicle_type"`
}

type Ticket struct {
	gorm.Model
	ID            string `gorm:"primaryKey"`
	VehicleTypeID uint   `json:"vehicle_type_id"`
	VehicleType   string
	VehicleNumber string     `json:"vehicle_number"`
	ParkingLotID  uint       `json:"parking_lot_id"`
	EntryTime     time.Time  `json:"entry_time"`
	ExitTime      *time.Time `json:"exit_time,omitempty"`
	IsParked      bool       `gorm:"default:true" json:"is_parked"`
}

type Receipt struct {
	gorm.Model
	ID           string    `gorm:"primaryKey"`
	VehicleType  string    `json:"vehicle_type"`
	ParkingLotID uint      `json:"parking_lot_id"`
	EntryTime    time.Time `json:"entry_time"`
	ExitTime     time.Time `json:"exit_time"`
	Rate         float64   `json:"rate"`
	RateType     string    `gorm:"type:enum('hourly', 'daily')"`
	BillAmount   float64   `json:"bill_amount"`
}

func (r *Receipt) CalculateBill(parkingLot ParkingLot) {
	duration := r.ExitTime.Sub(r.EntryTime)
	hours := duration.Hours()
	days := duration.Hours() / 24

	switch r.VehicleType {
	case "motorcycle":
		r.Rate = parkingLot.MotorcycleTariff
	case "car":
		r.Rate = parkingLot.CarTariff
	case "bus":
		if r.RateType == "daily" {
			r.Rate = parkingLot.BusTariffDaily
		} else {
			r.Rate = parkingLot.BusTariffHourly
		}
	default:
		r.Rate = 0
	}

	if r.RateType == "daily" {
		r.BillAmount = days * r.Rate
	} else {
		r.BillAmount = hours * r.Rate
	}
}
