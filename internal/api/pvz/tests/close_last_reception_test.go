package tests

import (
	"context"
	"encoding/json"
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

func TestCloseLastReception(t *testing.T) {
	gin.SetMode(gin.TestMode)
	type pvzServiceMockFunc func(mc *minimock.Controller) service.PVZService

	type args struct {
		pvzId string
	}

	var (
		ctx         = context.Background()
		mc          = minimock.NewController(t)
		validPVZId  = uuid.New()
		receptionID = uuid.New()
		now         = time.Now().UTC()
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
			name: "success case - valid PVZ ID",
			args: args{
				pvzId: validPVZId.String(),
			},
			wantStatus: http.StatusOK,
			wantResponse: &dto.ReceptionResponse{
				ID:       receptionID.String(),
				DateTime: now.Format(time.RFC3339),
				Status:   "close",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				mock := serviceMock.NewPVZServiceMock(mc)
				mock.CloseReceptionMock.Expect(ctx, validPVZId).Return(&model.Reception{
					ID:        receptionID,
					CreatedAt: now,
					Status:    "close",
				}, nil)
				return mock
			},
		},
		{
			name: "error case - invalid PVZ ID format",
			args: args{
				pvzId: "invalid-uuid",
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: dto.ErrorResponse{
				Message: "invalid pvzId",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				return serviceMock.NewPVZServiceMock(mc)
			},
		},
		{
			name: "error case - no open receptions",
			args: args{
				pvzId: validPVZId.String(),
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: dto.ErrorResponse{
				Message: "no open receptions",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				mock := serviceMock.NewPVZServiceMock(mc)
				mock.CloseReceptionMock.Expect(ctx, validPVZId).Return(nil, model.ErrNoOpenReceptions)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pvzServiceMock := tt.pvzServiceMock(mc)
			pvzImpl := pvz.NewImplementation(pvzServiceMock)

			router := gin.New()
			router.POST("/pvz/:pvzId/close_last_reception", pvzImpl.CloseLastReception)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/pvz/"+tt.args.pvzId+"/close_last_reception", nil)
			router.ServeHTTP(w, req)

			require.Equal(t, tt.wantStatus, w.Code)
			require.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

			switch expected := tt.wantResponse.(type) {
			case *dto.ReceptionResponse:
				var actual dto.ReceptionResponse
				err := json.Unmarshal(w.Body.Bytes(), &actual)
				require.NoError(t, err)
				require.Equal(t, expected.ID, actual.ID)
				require.Equal(t, expected.Status, actual.Status)

			case gin.H:
				var actual gin.H
				err := json.Unmarshal(w.Body.Bytes(), &actual)
				require.NoError(t, err)
				require.Equal(t, expected["error"], actual["error"])
			}
		})
	}
}
