package models

type ParkReq struct {
	VehicleTypeID uint   `json:"vehicle_type_id"`
	VehicleNumber string `json:"vehicle_number"`
	ParkingLotID  uint   `json:"parking_lot_id"`
}
