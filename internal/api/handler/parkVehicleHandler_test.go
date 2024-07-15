package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	domain "parking-lot-service/internal/Domain"
	handler "parking-lot-service/internal/api/handler"
	mock "parking-lot-service/internal/usecase/mock"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestParkVehicleHandler_ParkExit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock.NewMockParkVehicleUseCase(ctrl)
	handler := handler.NewParkVehicleHandler(mockUseCase)

	e := echo.New()

	t.Run("success", func(t *testing.T) {
		req := `{"ticket_id":"1", "exit_time":"2024-07-15T10:06:57.8758156+05:30"}`
		receipt := &domain.Receipt{
			ID:           "1",
			VehicleType:  "car",
			ParkingLotID: 1,
			EntryTime:    time.Date(2024, 7, 15, 9, 0, 0, 0, time.UTC),
			ExitTime:     time.Date(2024, 7, 15, 10, 6, 57, 8758156, time.UTC),
			Rate:         10,
			RateType:     "hourly",
			BillAmount:   10.0,
		}

		// exitTime, _ := time.Parse(time.RFC3339, "2024-07-15T10:06:57.8758156+05:30")
		mockUseCase.EXPECT().ParkExit(gomock.Eq(1), gomock.Any()).Return(receipt, nil)

		reqBody := strings.NewReader(req)
		request := httptest.NewRequest(http.MethodPost, "/parkexit", reqBody)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(request, rec)

		if assert.NoError(t, handler.ParkExit(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			expectedBody := `{"ID":"1","vehicle_type":"car","parking_lot_id":1,"entry_time":"2024-07-15T09:00:00Z","exit_time":"2024-07-15T10:06:57.008758156Z","rate":10,"RateType":"hourly","bill_amount":10}`
			assert.JSONEq(t, expectedBody, rec.Body.String())
		}
	})

}
