package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/biryanim/avito-tech-pvz/internal/api/auth"
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/biryanim/avito-tech-pvz/internal/service"
	serviceMock "github.com/biryanim/avito-tech-pvz/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/gojuno/minimock/v3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *http.Request
	}

	var (
		ctx         = context.Background()
		validEmail  = "test@example.com"
		validPass   = "password123"
		invalidPass = "wrongpassword"
		token       = "test-token"
	)

	tests := []struct {
		name            string
		args            args
		wantStatus      int
		wantResponse    interface{}
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case - valid credentials",
			args: args{
				req: func() *http.Request {
					reqBody := fmt.Sprintf(`{"email":"%s","password":"%s"}`, validEmail, validPass)
					req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(reqBody)))
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
				mock.LoginMock.Expect(ctx, &model.UserLoginInfo{
					Email:    validEmail,
					Password: validPass,
				}).Return(token, nil)
				return mock
			},
		},
		{
			name: "error case - invalid request format",
			args: args{
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(`invalid json`)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: &dto.ErrorResponse{
				Message: "invalid request",
			},
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				return serviceMock.NewAuthServiceMock(mc)
			},
		},
		{
			name: "error case - invalid credentials",
			args: args{
				req: func() *http.Request {
					reqBody := fmt.Sprintf(`{"email":"%s","password":"%s"}`, validEmail, invalidPass)
					req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(reqBody)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusUnauthorized,
			wantResponse: gin.H{
				"error": "invalid credentials",
			},
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(mc)
				mock.LoginMock.Expect(ctx, &model.UserLoginInfo{
					Email:    validEmail,
					Password: invalidPass,
				}).Return("", errors.New("invalid credentials"))
				return mock
			},
		},
		{
			name: "error case - missing email",
			args: args{
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(`{"password":"password"}`)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusBadRequest,
			wantResponse: &dto.ErrorResponse{
				Message: "invalid request",
			},
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				return serviceMock.NewAuthServiceMock(mc)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := minimock.NewController(t)
			defer mc.Finish()

			authServiceMock := tt.authServiceMock(mc)
			authImpl := auth.NewImplementation(authServiceMock)

			router := gin.New()
			router.POST("/login", authImpl.Login)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, tt.args.req)

			require.Equal(t, tt.wantStatus, w.Code)
			require.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

			switch expected := tt.wantResponse.(type) {
			case *dto.LoginResponse:
				var actual dto.LoginResponse
				err := json.Unmarshal(w.Body.Bytes(), &actual)
				require.NoError(t, err)
				require.Equal(t, expected.Token, actual.Token)

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
