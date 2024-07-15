package models

// ParkReq represents a request to park a vehicle
// @description ParkReq is the request structure for parking a vehicle
type ParkReq struct {
	VehicleTypeID uint   `json:"vehicle_type_id" example:"1"`
	VehicleNumber string `json:"vehicle_number" example:"ABC123"`
	ParkingLotID  uint   `json:"parking_lot_id" example:"1"`
}
type ExitRequest struct {
	TicketID int `json:"ticket_id" example:"1"`
}
