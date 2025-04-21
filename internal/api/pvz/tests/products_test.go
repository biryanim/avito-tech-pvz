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

func TestProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	type pvzServiceMockFunc func(mc *minimock.Controller) service.PVZService

	type args struct {
		reqBody string
	}

	var (
		ctx       = context.Background()
		mc        = minimock.NewController(t)
		pvzID     = uuid.New()
		productID = uuid.New()
		now       = time.Now()
		validReq  = fmt.Sprintf(`{"type":"электроника","pvzId":"%s"}`, pvzID.String())
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
			wantResponse: &dto.ProductResponse{
				ID:       productID.String(),
				Type:     "электроника",
				DateTime: now.Format(time.RFC3339),
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				mock := serviceMock.NewPVZServiceMock(mc)
				mock.AddProductToReceptionMock.Expect(ctx, &model.ProductPVZ{
					Type:  "электроника",
					PvzId: pvzID,
				}).Return(&model.Product{
					ID:        productID,
					Info:      model.ProductInfo{Type: "электроника"},
					CreatedAt: now,
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
				Message: "invalid request body",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				return serviceMock.NewPVZServiceMock(mc)
			},
		},
		{
			name: "error case - missing type",
			args: args{
				reqBody: fmt.Sprintf(`{"pvzId":"%s"}`, pvzID.String()),
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: dto.ErrorResponse{
				Message: "invalid request body",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				return serviceMock.NewPVZServiceMock(mc)
			},
		},
		{
			name: "error case - invalid PVZ ID",
			args: args{
				reqBody: `{"type":"электроника","pvzId":"invalid"}`,
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: dto.ErrorResponse{
				Message: "invalid product parameters",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				return serviceMock.NewPVZServiceMock(mc)
			},
		},
		{
			name: "error case - no open receptions",
			args: args{
				reqBody: validReq,
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: dto.ErrorResponse{
				Message: model.ErrNoOpenReceptions.Error(),
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				mock := serviceMock.NewPVZServiceMock(mc)
				mock.AddProductToReceptionMock.Expect(ctx, &model.ProductPVZ{
					Type:  "электроника",
					PvzId: pvzID,
				}).Return(nil, model.ErrNoOpenReceptions)
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
				mock.AddProductToReceptionMock.Expect(ctx, &model.ProductPVZ{
					Type:  "электроника",
					PvzId: pvzID,
				}).Return(nil, errors.New("database error"))
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pvzServiceMock := tt.pvzServiceMock(mc)
			pvzImpl := pvz.NewImplementation(pvzServiceMock)

			router := gin.New()
			router.POST("/products", pvzImpl.Products)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer([]byte(tt.args.reqBody)))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			require.Equal(t, tt.wantStatus, w.Code)
			require.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

			switch expected := tt.wantResponse.(type) {
			case *dto.ProductResponse:
				var actual dto.ProductResponse
				err := json.Unmarshal(w.Body.Bytes(), &actual)
				require.NoError(t, err)
				require.Equal(t, expected.ID, actual.ID)
				require.Equal(t, expected.Type, actual.Type)
			case dto.ErrorResponse:
				var actual dto.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &actual)
				require.NoError(t, err)
				require.Equal(t, expected.Message, actual.Message)
			}
		})
	}
}
