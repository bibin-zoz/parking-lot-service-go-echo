package main

import (
	"log"
	handlers "parking-lot-service/internal/api/handler"
	routes "parking-lot-service/internal/api/routes"
	config "parking-lot-service/internal/database"
	"parking-lot-service/internal/repository"
	usecase "parking-lot-service/internal/usecase"

	_ "parking-lot-service/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Parking Lot API
// @version 1.0
// @description This is a sample server for a parking lot service.

// @host localhost:8080
// @BasePath /
func main() {
	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories
	parkingLotRepo := repository.NewParkingLotRepository(db)
	parkVehicelRepo := repository.NewParkVehicleRepo(db)

	// Initialize use cases
	parkUseCase := usecase.NewParkingLotUseCase(parkingLotRepo)
	parkVehicleUseCase := usecase.NewParkVehicleUseCase(parkVehicelRepo, parkingLotRepo)

	// Initialize handlers
	parkhandler := handlers.NewHandler(parkUseCase)
	parkvehiclehandler := handlers.NewParkVehicleHandler(parkVehicleUseCase)

	// Setup routes
	routes.SetupRoutes(e, parkhandler, parkvehiclehandler)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve Swagger UI and API documentation
	e.Static("/swagger-ui", "swagger-ui")
	e.File("/swagger-ui/openapi.yaml", "openApi/openapi.yaml")
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
