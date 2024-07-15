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
func (repo *parkVehicleRepo) ParkVehicle(ticket *domain.Ticket) (*domain.Ticket, error) {
	ticket.EntryTime = time.Now()
	tx := repo.db.Begin()

	if err := tx.Create(ticket).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create ticket: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
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
func (repo *parkVehicleRepo) GetTicketDetailsByID(ticketID int) (*domain.Ticket, error) {
	var ticket domain.Ticket
	result := repo.db.First(&ticket, "id = ?", ticketID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("invalid ticket ID: %d", ticketID)
		}
		return nil, result.Error
	}
	return &ticket, nil
}
func (repo *parkVehicleRepo) SaveExitDetails(ticket *domain.Ticket, receipt *domain.Receipt) (*domain.Receipt, error) {
	tx := repo.db.Begin()

	if err := tx.Save(ticket).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to update ticket: %w", err)
	}

	if err := tx.Create(&receipt).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create receipt: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	fmt.Println(receipt)
	return receipt, nil
}
func (repo *parkVehicleRepo) GetParkingDetailsByVehicleNumber(vehicleNumber string) (*domain.Ticket, error) {
	var ticket domain.Ticket
	if err := repo.db.First(&ticket, "vehicle_number = ? and is_parked=true", vehicleNumber).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("vehicle number not found: %s", vehicleNumber)
		}
		return nil, fmt.Errorf("failed to get parking details: %w", err)
	}
	return &ticket, nil
}
