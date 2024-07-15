package models

type ParkReq struct {
	VehicleTypeID uint   `json:"vehicle_type_id" validate:"nonZeroPositive"`
	VehicleNumber string `json:"vehicle_number" validate:"vehicleNumber"`
	ParkingLotID  uint   `json:"parking_lot_id" validate:"nonZeroPositive"`
}

type ExitRequest struct {
	TicketID int `json:"ticket_id" validate:"min=1"`
}
