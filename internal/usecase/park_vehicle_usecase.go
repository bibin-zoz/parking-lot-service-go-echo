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
	// Fetch parking lot details
	parkingLotDetails, err := uc.parkingLotRepo.GetParkingLotByID(parkReq.ParkingLotID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch parking lot details: %w", err)
	}

	// Fetch vehicle details
	vehicleDetails, err := uc.parkVehicleRepo.GetVehicleDetails(parkReq.VehicleTypeID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch vehicle details: %w", err)
	}

	// Check if the vehicle is already parked
	existingTicket, err := uc.parkVehicleRepo.GetParkingDetailsByVehicleNumber(parkReq.VehicleNumber)
	if err == nil && existingTicket != nil {
		return nil, fmt.Errorf("vehicle with number %s is already parked", parkReq.VehicleNumber)
	}

	// Fetch vehicle counts by type
	counts, err := uc.parkingLotRepo.GetVehicleCountsByType(parkingLotDetails.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch vehicle counts: %w", err)
	}

	// Determine the number of free slots
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

	// Check if there are free slots
	if freeSlots <= 0 {
		return nil, fmt.Errorf("parking lot is full for %s", vehicleDetails.VehicleType)
	}

	// Generate a new ticket
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

	// Return the generated ticket
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

func (uc *ParkVehicleUseCase) ParkExit(ticketID int, exitTime time.Time) (*models.Receipt, error) {
	ticketDetails, err := uc.parkVehicleRepo.GetTicketDetailsByID(ticketID)
	if err != nil {
		return nil, fmt.Errorf("invalid ticket id: %w", err)
	}
	if !ticketDetails.IsParked {
		return &models.Receipt{}, fmt.Errorf("vehicle already checked out, invalid ID")
	}

	// Fetch parking lot details
	parkingLot, err := uc.parkingLotRepo.GetParkingLotByID(ticketDetails.ParkingLotID)
	if err != nil {
		return &models.Receipt{}, fmt.Errorf("failed to fetch parking lot details: %w", err)
	}

	receipt := &domain.Receipt{
		VehicleTypeID: int(ticketDetails.VehicleTypeID),
		VehicleType:   ticketDetails.VehicleType,
		ParkingLotID:  ticketDetails.ParkingLotID,
		EntryTime:     ticketDetails.EntryTime,
		ExitTime:      exitTime,
		RateType:      "hourly", // or set this based on your logic
	}

	// Calculate bill amount using Receipt's CalculateBill method
	receipt.CalculateBill(*parkingLot)

	// Update ticket details
	ticketDetails.ExitTime = &exitTime
	ticketDetails.IsParked = false

	// Delegate database operations to the repository
	genRecipt, err := uc.parkVehicleRepo.SaveExitDetails(ticketDetails, receipt)
	if err != nil {
		return &models.Receipt{}, fmt.Errorf("failed to save exit details: %w", err)
	}

	result := &models.Receipt{
		ID:           genRecipt.ID, // Assuming genRecipt has the ID field
		VehicleType:  genRecipt.VehicleType,
		ParkingLotID: genRecipt.ParkingLotID,
		EntryTime:    genRecipt.EntryTime,
		ExitTime:     genRecipt.ExitTime,
		Rate:         genRecipt.Rate,
		RateType:     genRecipt.RateType,
		BillAmount:   genRecipt.BillAmount,
	}

	return result, nil
}
func (uc *ParkVehicleUseCase) GetVehicleTypes() ([]domain.VehicleType, error) {
	vehicleTypes, err := uc.parkVehicleRepo.GetVehicleTypes()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch details: %w", err)
	}
	return vehicleTypes, nil
}
