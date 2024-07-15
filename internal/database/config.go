package db

import (
	"fmt"
	"log"
	"os"
	domain "parking-lot-service/internal/Domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadConfig() error {
	return godotenv.Load()
}

func ConnectDB() (*gorm.DB, error) {
	err := LoadConfig()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&domain.ParkingLot{}, &domain.Receipt{}, &domain.Ticket{}, &domain.VehicleType{})
	if err != nil {
		log.Println("failed to create table")
		return nil, err
	}

	err = CreateSampleParkingLots(db)
	if err != nil {
		log.Println("failed to create sample parking lots")
		return nil, err
	}

	err = CreateSampleVehicleTypes(db)
	if err != nil {
		log.Println("failed to create sample vehicle types")
		return nil, err
	}

	return db, nil
}

func CreateSampleParkingLots(db *gorm.DB) error {
	parkingLots := []domain.ParkingLot{
		{
			Name:             "Main Parking Lot",
			MotorcycleSpots:  10,
			CarSpots:         20,
			BusSpots:         5,
			MotorcycleTariff: 10.0,
			CarTariff:        20.0,
			BusTariffDaily:   100.0,
			BusTariffHourly:  15.0,
		},
	}

	for _, lot := range parkingLots {
		err := db.Create(&lot).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSampleVehicleTypes(db *gorm.DB) error {
	vehicleTypes := []domain.VehicleType{
		{
			VehicleType: "Motorcycle",
		},
		{
			VehicleType: "Car",
		},
		{
			VehicleType: "Bus",
		},
	}

	for _, vType := range vehicleTypes {
		err := db.Create(&vType).Error
		if err != nil {
			return err
		}
	}
	return nil
}
