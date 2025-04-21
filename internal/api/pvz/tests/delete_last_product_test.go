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
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteLastProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	type pvzServiceMockFunc func(mc *minimock.Controller) service.PVZService

	type args struct {
		pvzId string
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		validPVZId = uuid.New()
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
			wantResponse: gin.H{
				"message": "Товар удален",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				mock := serviceMock.NewPVZServiceMock(mc)
				mock.DeleteLastProductInReceptionMock.Expect(ctx, validPVZId).Return(nil)
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
			name: "error case - no such PVZ",
			args: args{
				pvzId: validPVZId.String(),
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: dto.ErrorResponse{
				Message: model.ErrNoSuchPvz.Error(),
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				mock := serviceMock.NewPVZServiceMock(mc)
				mock.DeleteLastProductInReceptionMock.Expect(ctx, validPVZId).Return(model.ErrNoSuchPvz)
				return mock
			},
		},
		{
			name: "error case - no open receptions",
			args: args{
				pvzId: validPVZId.String(),
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: dto.ErrorResponse{
				Message: model.ErrNoOpenReceptions.Error(),
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				mock := serviceMock.NewPVZServiceMock(mc)
				mock.DeleteLastProductInReceptionMock.Expect(ctx, validPVZId).Return(model.ErrNoOpenReceptions)
				return mock
			},
		},
		{
			name: "error case - no products",
			args: args{
				pvzId: validPVZId.String(),
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: dto.ErrorResponse{
				Message: model.ErrNoProducts.Error(),
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				mock := serviceMock.NewPVZServiceMock(mc)
				mock.DeleteLastProductInReceptionMock.Expect(ctx, validPVZId).Return(model.ErrNoProducts)
				return mock
			},
		},
		{
			name: "error case - service error",
			args: args{
				pvzId: validPVZId.String(),
			},
			wantStatus: http.StatusInternalServerError,
			wantResponse: dto.ErrorResponse{
				Message: "internal server error",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				mock := serviceMock.NewPVZServiceMock(mc)
				mock.DeleteLastProductInReceptionMock.Expect(ctx, validPVZId).Return(errors.New("database error"))
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pvzServiceMock := tt.pvzServiceMock(mc)
			pvzImpl := pvz.NewImplementation(pvzServiceMock)

			router := gin.New()
			router.POST("/pvz/:pvzId/delete_last_product", pvzImpl.DeleteLastProduct)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/pvz/"+tt.args.pvzId+"/delete_last_product", nil)
			router.ServeHTTP(w, req)

			require.Equal(t, tt.wantStatus, w.Code)
			require.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

			switch expected := tt.wantResponse.(type) {
			case gin.H:
				var actual gin.H
				err := json.Unmarshal(w.Body.Bytes(), &actual)
				require.NoError(t, err)
				require.Equal(t, expected["message"], actual["message"])

			case dto.ErrorResponse:
				var actual dto.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &actual)
				require.NoError(t, err)
				require.Equal(t, expected.Message, actual.Message)
			}
		})
	}
}
