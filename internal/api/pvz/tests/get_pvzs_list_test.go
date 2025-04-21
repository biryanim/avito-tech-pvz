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
	"time"
)

func TestGetPvzsList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	type pvzServiceMockFunc func(mc *minimock.Controller) service.PVZService

	type args struct {
		url string
	}

	var (
		ctx         = context.Background()
		mc          = minimock.NewController(t)
		pvzID       = uuid.New()
		receptionID = uuid.New()
		productID   = uuid.New()
		now         = time.Now()
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
			name: "success case - with all params",
			args: args{
				url: "/pvz?startDate=2023-01-01T00:00:00Z&endDate=2023-12-31T23:59:59Z&page=1&limit=10",
			},
			wantStatus: http.StatusOK,
			wantResponse: []*dto.PVZListResponse{
				{
					PVZ: dto.PVZResponse{
						ID:               pvzID.String(),
						City:             "Москва",
						RegistrationDate: now.String(),
					},
					Receptions: []dto.ReceptionsWithProducts{
						{
							Reception: dto.ReceptionResponse{
								ID:       receptionID.String(),
								DateTime: now.Format(time.RFC3339),
								Status:   "in_progress",
							},
							Products: []dto.ProductResponse{
								{
									ID:       productID.String(),
									Type:     "электроника",
									DateTime: now.Format(time.RFC3339),
								},
							},
						},
					},
				},
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				mock := serviceMock.NewPVZServiceMock(mc)
				startDate, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
				endDate, _ := time.Parse(time.RFC3339, "2023-12-31T23:59:59Z")
				mock.GetListPVZsMock.Expect(ctx, &model.Filter{
					StartDate: startDate,
					EndDate:   endDate,
					Offset:    0,
					Limit:     10,
				}).Return([]*model.PVZWithReceptions{
					{
						PVZ: model.PVZ{
							ID:               pvzID,
							Info:             model.PVZInfo{City: "Москва"},
							RegistrationDate: now,
						},
						Receptions: []model.ReceptionsWithProducts{
							{
								Reception: model.Reception{
									ID:        receptionID,
									PvzId:     pvzID,
									CreatedAt: now,
									Status:    "in_progress",
								},
								Products: []model.Product{
									{
										ID: productID,
										Info: model.ProductInfo{
											Type:        "электроника",
											ReceptionId: receptionID,
										},
										CreatedAt: now,
									},
								},
							},
						},
					},
				}, nil)
				return mock
			},
		},
		{
			name: "success case - empty result",
			args: args{
				url: "/pvz",
			},
			wantStatus:   http.StatusOK,
			wantResponse: []*dto.PVZListResponse{},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				mock := serviceMock.NewPVZServiceMock(mc)
				mock.GetListPVZsMock.Expect(ctx, &model.Filter{
					Offset: 0,
					Limit:  10,
				}).Return([]*model.PVZWithReceptions{}, nil)
				return mock
			},
		},
		{
			name: "error case - invalid date format",
			args: args{
				url: "/pvz?startDate=invalid",
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: gin.H{
				"error": "invalid date format",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				return serviceMock.NewPVZServiceMock(mc)
			},
		},
		{
			name: "error case - service error",
			args: args{
				url: "/pvz",
			},
			wantStatus: http.StatusInternalServerError,
			wantResponse: gin.H{
				"error": "database error",
			},
			pvzServiceMock: func(mc *minimock.Controller) service.PVZService {
				mock := serviceMock.NewPVZServiceMock(mc)
				mock.GetListPVZsMock.Expect(ctx, &model.Filter{
					Offset: 0,
					Limit:  10,
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
			router.GET("/pvz", pvzImpl.GetPVZs)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tt.args.url, nil)
			router.ServeHTTP(w, req)

			require.Equal(t, tt.wantStatus, w.Code)
			require.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

			switch expected := tt.wantResponse.(type) {
			case []*dto.PVZListResponse:
				var actual []*dto.PVZListResponse
				err := json.Unmarshal(w.Body.Bytes(), &actual)
				require.NoError(t, err)

				require.Equal(t, len(expected), len(actual))
				if len(expected) > 0 {
					require.Equal(t, expected[0].PVZ.ID, actual[0].PVZ.ID)
					require.Equal(t, expected[0].PVZ.City, actual[0].PVZ.City)

					if len(expected[0].Receptions) > 0 {
						require.Equal(t, expected[0].Receptions[0].Reception.ID,
							actual[0].Receptions[0].Reception.ID)

						if len(expected[0].Receptions[0].Products) > 0 {
							require.Equal(t, expected[0].Receptions[0].Products[0].ID,
								actual[0].Receptions[0].Products[0].ID)
						}
					}
				}

			case gin.H:
				var actual gin.H
				err := json.Unmarshal(w.Body.Bytes(), &actual)
				require.NoError(t, err)
				require.Equal(t, expected["error"], actual["error"])
			}
		})
	}
}
