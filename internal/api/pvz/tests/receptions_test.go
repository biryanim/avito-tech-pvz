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
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestReceptions(t *testing.T) {
	gin.SetMode(gin.TestMode)
	type pvzServiceMockFunc func(mc *minimock.Controller) service.PVZService

	type args struct {
		reqBody string
	}

	var (
		ctx         = context.Background()
		mc          = minimock.NewController(t)
		pvzID       = uuid.New()
		receptionID = uuid.New()
		now         = time.Now()
		validReq    = fmt.Sprintf(`{"pvzId":"%s"}`, pvzID.String())
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
				reqBody: validReq,
			},
			wantStatus: http.StatusCreated,
			wantResponse: &dto.ReceptionResponse{
				ID:       receptionID.String(),
				DateTime: now.Format(time.RFC3339),
				Status:   "in_progress",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				mock := serviceMock.NewPVZServiceMock(mc)
				mock.CreateReceptionMock.Expect(ctx, pvzID).Return(&model.Reception{
					ID:        receptionID,
					CreatedAt: now,
					Status:    "in_progress",
				}, nil)
				return mock
			},
		},
		{
			name: "error case - invalid JSON",
			args: args{
				reqBody: `invalid json`,
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: dto.ErrorResponse{
				Message: "invalid request",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				return serviceMock.NewPVZServiceMock(mc)
			},
		},
		{
			name: "error case - missing pvzId",
			args: args{
				reqBody: `{}`,
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: dto.ErrorResponse{
				Message: "invalid request",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				return serviceMock.NewPVZServiceMock(mc)
			},
		},
		{
			name: "error case - invalid pvzId format",
			args: args{
				reqBody: `{"pvzId":"invalid"}`,
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
			name: "error case - last reception not closed",
			args: args{
				reqBody: validReq,
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: dto.ErrorResponse{
				Message: model.ErrLastReceptionNotClosed.Error(),
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				mock := serviceMock.NewPVZServiceMock(mc)
				mock.CreateReceptionMock.Expect(ctx, pvzID).Return(nil, model.ErrLastReceptionNotClosed)
				return mock
			},
		},
		{
			name: "error case - service error",
			args: args{
				reqBody: validReq,
			},
			wantStatus: http.StatusInternalServerError,
			wantResponse: dto.ErrorResponse{
				Message: "internal server error",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				mock := serviceMock.NewPVZServiceMock(mc)
				mock.CreateReceptionMock.Expect(ctx, pvzID).Return(nil, errors.New("database error"))
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pvzServiceMock := tt.pvzServiceMock(mc)
			pvzImpl := pvz.NewImplementation(pvzServiceMock)

			router := gin.New()
			router.POST("/receptions", pvzImpl.Receptions)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/receptions", bytes.NewBuffer([]byte(tt.args.reqBody)))
			req.Header.Set("Content-Type", "application/json")
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
			case dto.ErrorResponse:
				var actual dto.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &actual)
				require.NoError(t, err)
				require.Equal(t, expected.Message, actual.Message)
			}
		})
	}
}
