package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/biryanim/avito-tech-pvz/internal/api/auth"
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/biryanim/avito-tech-pvz/internal/service"
	serviceMock "github.com/biryanim/avito-tech-pvz/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDummyLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *http.Request
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		moderatorRole model.Role = "moderator"
		employeeRole  model.Role = "employee"
		token                    = "expected-token"

		reqBodyWithModeratorRole = `{"role":"moderator"}`
		reqBodyWithEmployeeRole  = `{"role":"employee"}`
		reqBodyWithWrongRole     = `{"role":"qwerty"}`
	)

	tests := []struct {
		name            string
		args            args
		wantStatus      int
		wantResponse    interface{}
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case - moderator role",
			args: args{
				ctx: ctx,
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/dummyLogin", bytes.NewBuffer([]byte(reqBodyWithModeratorRole)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusOK,
			wantResponse: &dto.LoginResponse{
				Token: token,
			},
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(mc)
				mock.DummyLoginMock.Expect(ctx, moderatorRole).Return(token, nil)

				return mock
			},
		},
		{
			name: "success case - employee role",
			args: args{
				ctx: ctx,
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/dummyLogin", bytes.NewBuffer([]byte(reqBodyWithEmployeeRole)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusOK,
			wantResponse: &dto.LoginResponse{
				Token: token,
			},
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(mc)
				mock.DummyLoginMock.Expect(ctx, employeeRole).Return(token, nil)

				return mock
			},
		},
		{
			name: "error case - invalid role",
			args: args{
				ctx: ctx,
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/dummyLogin", bytes.NewBuffer([]byte(reqBodyWithWrongRole)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: &dto.ErrorResponse{
				Message: "invalid role",
			},
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				return serviceMock.NewAuthServiceMock(mc)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authServiceMock := tt.authServiceMock(mc)
			authImpl := auth.NewImplementation(authServiceMock)

			router := gin.New()
			router.POST("/dummyLogin", authImpl.DummyLogin)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, tt.args.req)

			require.Equal(t, tt.wantStatus, w.Code)

			switch expected := tt.wantResponse.(type) {
			case *dto.LoginResponse:
				var actual dto.LoginResponse
				if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
					t.Errorf("unmarshal response error: %s", err)
				}
				require.Equal(t, expected.Token, actual.Token)

			case *dto.ErrorResponse:
				var errResp dto.ErrorResponse
				if err := json.Unmarshal(w.Body.Bytes(), &errResp); err != nil {
					t.Errorf("unmarshal response error: %s", err)
				}
				require.Equal(t, expected.Message, errResp.Message)
			}
		})
	}
}
