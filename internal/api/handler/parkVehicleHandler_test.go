package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	handler "parking-lot-service/internal/api/handler"
	"parking-lot-service/internal/models"
	mock "parking-lot-service/internal/usecase/mock"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParkVehicleHandler_ParkVehicle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock.NewMockParkVehicleUseCase(ctrl)
	handler := handler.NewParkVehicleHandler(mockUseCase)

	e := echo.New()

	t.Run("success", func(t *testing.T) {
		req := `{"vehicle_number":"AB14G6123", "vehicle_type_id":1, "parking_lot_id":1}`

		parkReq := models.ParkReq{
			VehicleNumber: "AB14G6123",
			VehicleTypeID: 1,
			ParkingLotID:  1,
		}

		createdTicket := &models.Ticket{
			ID:            1,
			VehicleNumber: "AB14G6123",
			VehicleType:   "car",
			ParkingLotID:  1,
			EntryTime:     time.Now(),
			IsParked:      true,
		}

		mockUseCase.EXPECT().ParkVehicle(parkReq).Return(createdTicket, nil)

		reqBody := strings.NewReader(req)
		request := httptest.NewRequest(http.MethodPost, "/park-vehicle", reqBody)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(request, rec)

		if assert.NoError(t, handler.ParkVehicle(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)

			var response models.Ticket
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.Equal(t, createdTicket.ID, response.ID)
			assert.Equal(t, createdTicket.VehicleNumber, response.VehicleNumber)
			assert.Equal(t, createdTicket.VehicleType, response.VehicleType)
			assert.Equal(t, createdTicket.ParkingLotID, response.ParkingLotID)
			assert.WithinDuration(t, createdTicket.EntryTime, response.EntryTime, time.Second)
			assert.True(t, response.IsParked)
		}
	})
}
func TestParkVehicleHandler_ParkExit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock.NewMockParkVehicleUseCase(ctrl)
	handler := handler.NewParkVehicleHandler(mockUseCase)

	e := echo.New()

	t.Run("success", func(t *testing.T) {
		req := `{"ticket_id":1, "exit_time":"2024-07-15T10:06:57.8758156+05:30"}`
		receipt := &models.Receipt{
			ID:           1,
			VehicleType:  "car",
			ParkingLotID: 1,
			EntryTime:    time.Date(2024, 7, 15, 9, 0, 0, 0, time.UTC),
			ExitTime:     time.Date(2024, 7, 15, 10, 6, 57, 8758156, time.UTC),
			Rate:         10,
			RateType:     "hourly",
			BillAmount:   10.0,
		}

		// Mock the expected call to ParkExit
		mockUseCase.EXPECT().ParkExit(1, gomock.Any()).Return(receipt, nil)

		reqBody := strings.NewReader(req)
		request := httptest.NewRequest(http.MethodDelete, "/park-vehicle", reqBody) // Use DELETE method
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(request, rec)

		if assert.NoError(t, handler.ParkExit(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			// Marshal the receipt into JSON for comparison
			expectedBody, err := json.Marshal(receipt)
			require.NoError(t, err)

			assert.JSONEq(t, string(expectedBody), rec.Body.String())
		}
	})

	t.Run("invalid_request_format", func(t *testing.T) {
		req := `{"ticket_id": -10}` // Provide a negative value as a string

		reqBody := strings.NewReader(req)
		request := httptest.NewRequest(http.MethodDelete, "/park-vehicle", reqBody) // Use DELETE method
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(request, rec)

		// Assertions
		handler.ParkExit(c)
		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Contains(t, response["error"], "invalid ticket id")
	})

	t.Run("already_checked_out", func(t *testing.T) {
		req := `{"ticket_id": 1}`

		// Set up mock expectation for ParkExit to return an error indicating vehicle already checked out
		mockUseCase.EXPECT().ParkExit(1, gomock.Any()).Return(nil, fmt.Errorf("vehicle already checked out, invalid ID"))

		reqBody := strings.NewReader(req)
		request := httptest.NewRequest(http.MethodDelete, "/park-vehicle", reqBody) // Use DELETE method
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(request, rec)

		// Call the handler function
		err := handler.ParkExit(c)

		// Check the response body for error details
		var response map[string]string
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Contains(t, response["error"], "vehicle already checked out") // Check for specific error message
	})
}
