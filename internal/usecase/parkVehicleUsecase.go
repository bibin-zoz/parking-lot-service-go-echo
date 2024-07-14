// internal/usecase/park_vehicle_uc.go
package usecase

import (
	repository "parking-lot-service/internal/repository/interface"
)

type ParkVehicleUseCase struct {
	parkVehicleRepo repository.ParkVehicleRepository
	parkingLotRepo  repository.ParkingLotRepository
}

func NewParkVehicleUseCase(parkVehicleRepo repository.ParkVehicleRepository) *ParkVehicleUseCase {
	return &ParkVehicleUseCase{
		parkVehicleRepo: parkVehicleRepo,
	}
}

// func (uc *ParkVehicleUseCase) ParkVehicle(parkReq models.ParkReq) (*domain.Ticket, error) {
// 	parkingLotDetails, err := uc.parkingLotRepo.GetParkingLotByID(parkReq.ParkingLotID)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to fetch parkinglotdetails", err)
// 	}
// 	vehicleDetails, err := uc.parkVehicleRepo.GetVehicleDetails(parkReq.VehicleTypeID)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to fetch parkinglotdetails", err)
// 	}
// }

// func (uc *ParkVehicleUseCase) ParkExit(ticketID string, exitTime time.Time) error {
// 	// Implement any validation or business logic before calling repository
// 	return uc.repo.ParkExit(ticketID, exitTime)
// }
