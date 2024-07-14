package repository

import (
	domain "parking-lot-service/internal/Domain"
	"parking-lot-service/internal/models"
	repository "parking-lot-service/internal/repository/interface"

	"gorm.io/gorm"
)

type parkingLotRepo struct {
	db *gorm.DB
}

func NewParkingLotRepository(psqlDB *gorm.DB) repository.ParkingLotRepository {
	return &parkingLotRepo{
		db: psqlDB,
	}

}

func (r *parkingLotRepo) GetAllParkingLots() ([]models.ParkingLot, error) {
	var parkingLots []models.ParkingLot
	result := r.db.Find(&parkingLots)
	return parkingLots, result.Error
}

func (r *parkingLotRepo) GetParkingLotByID(id uint) (*models.ParkingLot, error) {
	var parkingLot models.ParkingLot
	result := r.db.First(&parkingLot, id)
	return &parkingLot, result.Error
}

func (r *parkingLotRepo) CreateParkingLot(parkingLot *models.ParkingLot) error {
	result := r.db.Create(parkingLot)
	return result.Error
}

func (r *parkingLotRepo) UpdateParkingLot(parkingLot *models.ParkingLot) error {
	result := r.db.Save(parkingLot)
	return result.Error
}

func (r *parkingLotRepo) DeleteParkingLot(id uint) error {
	result := r.db.Delete(&models.ParkingLot{}, id)
	return result.Error
}
func (repo *parkingLotRepo) GetVehicleCountsByType(parkingLotID uint) (map[uint]int, error) {
	type CountResult struct {
		VehicleTypeID uint
		Count         int
	}

	var results []CountResult
	counts := make(map[uint]int)

	// Single query to count parked vehicles by type
	err := repo.db.Model(&domain.Ticket{}).
		Select("vehicle_type_id, COUNT(*) as count").
		Where("parking_lot_id = ? AND status = 'parked'", parkingLotID).
		Group("vehicle_type_id").
		Scan(&results).Error

	if err != nil {
		return counts, err
	}

	// Populate the map with the results
	for _, result := range results {
		counts[result.VehicleTypeID] = result.Count
	}

	return counts, nil
}
