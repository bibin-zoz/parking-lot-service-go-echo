package handlers

import usecase "parking-lot-service/internal/usecase/interface"

type Handler struct {
	parkUseCase usecase.ParkUseCase
}

func NewHandler(parkUseCase usecase.ParkUseCase) *Handler {
	return &Handler{parkUseCase: parkUseCase}
}
