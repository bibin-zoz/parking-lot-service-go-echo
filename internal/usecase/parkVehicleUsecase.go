// internal/usecase/park_vehicle_uc.go
package usecase

import (
	"fmt"
	domain "parking-lot-service/internal/Domain"
	"parking-lot-service/internal/models"
	repository "parking-lot-service/internal/repository/interface"
	"time"
)

type ParkVehicleUseCase struct {
	parkVehicleRepo repository.ParkVehicleRepository
	parkingLotRepo  repository.ParkingLotRepository
}

func NewParkVehicleUseCase(parkVehicleRepo repository.ParkVehicleRepository, parkingLotRepo repository.ParkingLotRepository) *ParkVehicleUseCase {
	return &ParkVehicleUseCase{
		parkVehicleRepo: parkVehicleRepo,
		parkingLotRepo:  parkingLotRepo,
	}
}

func (uc *ParkVehicleUseCase) ParkVehicle(parkReq models.ParkReq) (*models.Ticket, error) {

	parkingLotDetails, err := uc.parkingLotRepo.GetParkingLotByID(parkReq.ParkingLotID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch parking lot details: %w", err)
	}

	vehicleDetails, err := uc.parkVehicleRepo.GetVehicleDetails(parkReq.VehicleTypeID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch vehicle details: %w", err)
	}

	counts, err := uc.parkingLotRepo.GetVehicleCountsByType(parkingLotDetails.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch vehicle counts: %w", err)
	}

	var freeSlots int
	switch vehicleDetails.VehicleType {
	case "Motorcycle":
		freeSlots = parkingLotDetails.MotorcycleSpots - counts[1]
	case "Car":
		freeSlots = parkingLotDetails.CarSpots - counts[2]
	case "Bus":
		freeSlots = parkingLotDetails.BusSpots - counts[3]
	default:
		return nil, fmt.Errorf("invalid vehicle type")
	}

	if freeSlots <= 0 {
		return nil, fmt.Errorf("parking lot is full for %s", vehicleDetails.VehicleType)
	}

	ticketGenReq := domain.Ticket{
		VehicleTypeID: vehicleDetails.ID,
		VehicleType:   vehicleDetails.VehicleType,
		VehicleNumber: parkReq.VehicleNumber,
		ParkingLotID:  parkReq.ParkingLotID,
	}
	ticket, err := uc.parkVehicleRepo.ParkVehicle(&ticketGenReq)
	if err != nil {
		return nil, fmt.Errorf("failed to park vehicle: %w", err)
	}

	return &models.Ticket{
		ID:            ticket.ID,
		VehicleTypeID: ticket.VehicleTypeID,
		VehicleType:   ticket.VehicleType,
		VehicleNumber: ticket.VehicleNumber,
		ParkingLotID:  ticket.ParkingLotID,
		EntryTime:     ticket.EntryTime,
		IsParked:      ticket.IsParked,
	}, nil
}

func (uc *ParkVehicleUseCase) ParkExit(ticketID int, exitTime time.Time) (*domain.Receipt, error) {
	ticketDetails, err := uc.parkVehicleRepo.GetTicketDetailsByID(ticketID)
	if err != nil {
		return nil, fmt.Errorf("invalid ticket id: %w", err)
	}

	// Fetch parking lot details
	parkingLot, err := uc.parkingLotRepo.GetParkingLotByID(ticketDetails.ParkingLotID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch parking lot details: %w", err)
	}

	receipt := &domain.Receipt{
		VehicleType:  ticketDetails.VehicleType,
		ParkingLotID: ticketDetails.ParkingLotID,
		EntryTime:    ticketDetails.EntryTime,
		ExitTime:     exitTime,
		RateType:     "hourly", // or set this based on your logic
	}

	// Calculate bill amount using Receipt's CalculateBill method
	receipt.CalculateBill(*parkingLot)

	// Update ticket details
	ticketDetails.ExitTime = &exitTime
	ticketDetails.IsParked = false

	// Delegate database operations to the repository
	receipt, err = uc.parkVehicleRepo.SaveExitDetails(ticketDetails, receipt)
	if err != nil {
		return nil, fmt.Errorf("failed to save exit details: %w", err)
	}

	return receipt, nil
}
