package models

type ParkReq struct {
	VehicleTypeID uint   `json:"vehicle_type_id" validate:"nonZeroPositive" example:"1"`
	VehicleNumber string `json:"vehicle_number" validate:"vehicleNumber" example:"AB01C1234"`
	ParkingLotID  uint   `json:"parking_lot_id" validate:"nonZeroPositive" example:"1"`
}

type ExitRequest struct {
	TicketID int `json:"ticket_id" validate:"min=1" example:"1"`
}
