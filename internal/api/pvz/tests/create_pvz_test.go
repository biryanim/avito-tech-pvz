package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/api/pvz"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/biryanim/avito-tech-pvz/internal/service"
	serviceMock "github.com/biryanim/avito-tech-pvz/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreatePVZ(t *testing.T) {
	gin.SetMode(gin.TestMode)
	type pvzServiceMockFunc func(mc *minimock.Controller) service.PVZService

	type args struct {
		req *http.Request
	}

	var (
		ctx       = context.Background()
		mc        = minimock.NewController(t)
		pvzID     = uuid.New()
		validCity = "Москва"
	)
	defer mc.Finish()

	tests := []struct {
		name           string
		args           args
		wantStatus     int
		wantResponse   interface{}
		pvzServiceMock pvzServiceMockFunc
	}{
		{
			name: "success case - valid request",
			args: args{
				req: func() *http.Request {
					reqBody := fmt.Sprintf(`{"city":"%s"}`, validCity)
					req, _ := http.NewRequest("POST", "/pvz", bytes.NewBuffer([]byte(reqBody)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusCreated,
			wantResponse: &dto.PVZResponse{
				ID:               pvzID.String(),
				City:             validCity,
				RegistrationDate: time.Now().Format(time.RFC3339),
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				mock := serviceMock.NewPVZServiceMock(mc)
				mock.CreatePVZMock.Expect(ctx, &model.PVZInfo{
					City: model.City(validCity),
				}).Return(&model.PVZ{
					ID: pvzID,
					Info: model.PVZInfo{
						City: model.City(validCity),
					},
					RegistrationDate: time.Now(),
				}, nil)
				return mock
			},
		},
		{
			name: "error case - invalid JSON",
			args: args{
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/pvz", bytes.NewBuffer([]byte(`invalid json`)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: &dto.ErrorResponse{
				Message: "invalid request",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				return serviceMock.NewPVZServiceMock(mc)
			},
		},
		{
			name: "error case - missing city",
			args: args{
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/pvz", bytes.NewBuffer([]byte(`{}`)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: &dto.ErrorResponse{
				Message: "invalid request",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				return serviceMock.NewPVZServiceMock(mc)
			},
		},
		{
			name: "error case - invalid city",
			args: args{
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/pvz", bytes.NewBuffer([]byte(`{"city":"invalid"}`)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: &dto.ErrorResponse{
				Message: "invalid city",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				return serviceMock.NewPVZServiceMock(mc)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pvzServiceMock := tt.pvzServiceMock(mc)
			pvzImpl := pvz.NewImplementation(pvzServiceMock)

			router := gin.New()
			router.POST("/pvz", pvzImpl.CreatePvz)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, tt.args.req)

			require.Equal(t, tt.wantStatus, w.Code)
			require.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

			switch expected := tt.wantResponse.(type) {
			case *dto.PVZResponse:
				var actual dto.PVZResponse
				err := json.Unmarshal(w.Body.Bytes(), &actual)
				require.NoError(t, err)
				require.Equal(t, expected.ID, actual.ID)
				require.Equal(t, expected.City, actual.City)

			case *dto.ErrorResponse:
				var errResp dto.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &errResp)
				require.NoError(t, err)
				require.Equal(t, expected.Message, errResp.Message)

			case gin.H:
				var actual gin.H
				err := json.Unmarshal(w.Body.Bytes(), &actual)
				require.NoError(t, err)
				require.Equal(t, expected["error"], actual["error"])
			}
		})
	}
}
