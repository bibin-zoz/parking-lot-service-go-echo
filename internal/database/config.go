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

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbPort)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if dbErr != nil {
		log.Fatal(dbErr)
		return nil, dbErr
	}

	var exists bool
	// Check if the database exists
	err = db.Raw("SELECT EXISTS (SELECT FROM pg_database WHERE datname = ?)", dbName).Scan(&exists).Error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Create the database if it does not exist
	if !exists {
		err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		log.Println("Created database " + dbName)
	}

	// Connect to the newly created database
	db, err = gorm.Open(postgres.Open(psqlInfo+" dbname="+dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// AutoMigrate the domain models
	err = db.AutoMigrate(&domain.ParkingLot{}, &domain.Receipt{}, &domain.Ticket{}, &domain.VehicleType{})
	if err != nil {
		log.Fatal("Failed to create tables:", err)
		return nil, err
	}

	err = CreateSampleParkingLots(db)
	if err != nil {
		log.Println("Failed to create sample parking lots")
		return nil, err
	}

	err = CreateSampleVehicleTypes(db)
	if err != nil {
		log.Println("Failed to create sample vehicle types")
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
		err := db.Save(&lot).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSampleVehicleTypes(db *gorm.DB) error {
	vehicleTypes := []domain.VehicleType{
		{
			ID:          1,
			VehicleType: "Motorcycle",
		},
		{
			ID:          2,
			VehicleType: "Car",
		},
		{
			ID:          3,
			VehicleType: "Bus",
		},
	}

	for _, vType := range vehicleTypes {
		err := db.Save(&vType).Error
		if err != nil {
			return err
		}
	}
	return nil
}
