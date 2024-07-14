package repository

import (
	"fmt"
	domain "parking-lot-service/internal/Domain"
	repository "parking-lot-service/internal/repository/interface"
	"time"

	"gorm.io/gorm"
)

type parkVehicleRepo struct {
	db *gorm.DB
}

func NewParkVehicleRepo(psqlDB *gorm.DB) repository.ParkVehicleRepository {
	return &parkVehicleRepo{
		db: psqlDB,
	}

}
func (repo *parkVehicleRepo) ParkVehicle(vehicleType string, vehicleNumber string, parkingLotID uint) (*domain.Ticket, error) {
	ticket := &domain.Ticket{
		VehicleType:   vehicleType,
		VehicleNumber: vehicleNumber,
		ParkingLotID:  parkingLotID,
		EntryTime:     time.Now(),
	}

	if err := repo.db.Create(ticket).Error; err != nil {
		return nil, fmt.Errorf("failed to create ticket: %w", err)
	}

	return ticket, nil
}

func (repo *parkVehicleRepo) GenerateReceipt(ticketID string, exitTime time.Time) (*domain.Receipt, error) {
	var ticket domain.Ticket
	if err := repo.db.First(&ticket, "id = ?", ticketID).Error; err != nil {
		return nil, fmt.Errorf("failed to find ticket: %w", err)
	}

	receipt := &domain.Receipt{
		VehicleType:  ticket.VehicleType,
		ParkingLotID: ticket.ParkingLotID,
		EntryTime:    ticket.EntryTime,
		ExitTime:     exitTime,
		RateType:     "hourly", // or determine rate type based on business logic
	}

	// Calculate bill amount using Receipt's CalculateBill method

	if err := repo.db.Create(receipt).Error; err != nil {
		return nil, fmt.Errorf("failed to generate receipt: %w", err)
	}

	return receipt, nil
}

func (repo *parkVehicleRepo) ParkExit(ticketID string, exitTime time.Time) (*domain.Receipt, error) {
	var ticket domain.Ticket
	if err := repo.db.First(&ticket, "id = ?", ticketID).Error; err != nil {
		return nil, fmt.Errorf("failed to find ticket: %w", err)
	}

	ticket.ExitTime = &exitTime
	ticket.IsParked = false

	tx := repo.db.Begin()
	if err := tx.Save(&ticket).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to update ticket: %w", err)
	}

	// Generate receipt
	receipt, err := repo.GenerateReceipt(ticketID, exitTime)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return receipt, nil
}

func (repo *parkVehicleRepo) GetVehicleDetails(vehicleID uint) (*domain.VehicleType, error) {
	var vehicleType domain.VehicleType
	result := repo.db.First(&vehicleType, vehicleID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("invalid vehicle ID: %d", vehicleID)
		}
		return nil, result.Error
	}
	return &vehicleType, nil
}
