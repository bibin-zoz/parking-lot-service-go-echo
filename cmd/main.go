package main

import (
	"log"
	routes "parking-lot-service/internal/api"
	handlers "parking-lot-service/internal/api/handler"
	config "parking-lot-service/internal/database"
	"parking-lot-service/internal/repository"
	usecase "parking-lot-service/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	parkingLotRepo := repository.NewParkingLotRepository(db)

	parkUseCase := usecase.NewParkUseCase(parkingLotRepo)

	// Initialize HTTP handler
	parkhandler := handlers.NewHandler(parkUseCase)
	routes.SetupRoutes(e, parkhandler)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
