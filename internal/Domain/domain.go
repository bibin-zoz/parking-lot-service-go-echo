package domain

import (
	"parking-lot-service/internal/models"
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
	ID            int  `gorm:"primaryKey"`
	VehicleTypeID uint `json:"vehicle_type_id"`
	VehicleType   string
	VehicleNumber string     `json:"vehicle_number"`
	ParkingLotID  uint       `json:"parking_lot_id"`
	EntryTime     time.Time  `json:"entry_time"`
	ExitTime      *time.Time `json:"exit_time,omitempty"`
	IsParked      bool       `gorm:"default:true" json:"is_parked"`
}

type Receipt struct {
	gorm.Model
	ID            int       `gorm:"primaryKey"`
	VehicleTypeID int       `json:"vehicle_typeID"`
	VehicleType   string    `json:"vehicle_type"`
	ParkingLotID  uint      `json:"parking_lot_id"`
	EntryTime     time.Time `json:"entry_time"`
	ExitTime      time.Time `json:"exit_time"`
	Rate          float64   `json:"rate"`
	RateType      string    `json:"RateType"`
	BillAmount    float64   `json:"bill_amount"`
}

func (r *Receipt) CalculateBill(parkingLot models.ParkingLot) {
	duration := r.ExitTime.Sub(r.EntryTime)
	hours := duration.Hours()
	days := duration.Hours() / 24

	switch r.VehicleTypeID {
	case 1:
		r.Rate = parkingLot.MotorcycleTariff
	case 2:
		r.Rate = parkingLot.CarTariff
	case 3:
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
