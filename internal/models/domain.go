package domain

import "time"

type VehicleType string

const (
	Motorcycle VehicleType = "motorcycle"
	Car        VehicleType = "car"
	Bus        VehicleType = "bus"
)

type ParkingLot struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"unique"`
	MotorbikeSpots int
	CarSpots       int
	BusSpots       int
}

type Vehicle struct {
	ID           string `gorm:"primaryKey"`
	Type         VehicleType
	EntryTime    time.Time
	ExitTime     *time.Time
	ParkingLotID uint
}

type Ticket struct {
	ID           string `gorm:"primaryKey"`
	VehicleID    string
	ParkingLotID uint
	EntryTime    time.Time
}
