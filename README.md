# Parking Lot Service

A HTTP service to manage parking lots, vehicles, and fees using Go (Echo framework), PostgreSQL, and Docker.

## Features

- **CRUD Operations**: Manage parking lots with Create, Read, Update, and Delete operations.
- **Vehicle Parking**: Park vehicles with validation for vehicle types and generation of entry tickets.
- **Vehicle Exit**: Handle vehicle exits, generate receipts with calculated fees based on parking duration and tariff models.
- **Real-Time Slot Management**: Track and display the number of available slots in each parking lot.
- **Tariff Models**: Support different tariff models for motorcycles/scooters, cars/SUVs, and buses/trucks.
- **Documentation**: Includes Swagger documentation for API endpoints.

## Hosted on AWS Instance

The Parking Lot Service is hosted on an AWS EC2 instance.

- **Live URL**: [parkinglot.bibinvinod.online/swagger/index.html](https://parkinglot.bibinvinod.online/swagger/index.html)


## Getting Started

### Prerequisites

- Go installed on your machine.
- Docker installed and running.
- PostgreSQL database configured.

### Installation

1. Clone the repository:

   ```bash
   git clone [<repository-url>](https://github.com/bibin-zoz/parking-lot-service-go-echo.git)
   cd parking-lot-service
   ```

2. Set up environment variables:

   Create a `.env` file in the root directory with the following variables:

   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=parkinglotdb
   ```

3. Build and run the application using Docker:

   ```bash
   docker-compose up --build
   ```

   This command will build the Docker image and start the application.

4. Access the API documentation:

   Open your web browser and go to `http://localhost:8081/swagger/index.html` to view Swagger documentation.

## Usage

- **Create a Parking Lot**: Use API endpoint to create a new parking lot with specified details.
- **Park a Vehicle**: POST request to `/park-vehicle` with vehicle details to park it and generate a ticket.
- **Exit a Vehicle**: DELETE request to `/park-vehicle` with ticket ID to generate a receipt and calculate fees.
- **Get Free Slots**: GET request to `/parking-lots/free-slots/{parkingLotID}` to get the number of available slots in a parking lot.


