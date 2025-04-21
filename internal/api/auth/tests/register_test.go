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
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *http.Request
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		validEmail = "test@example.com"
		validPass  = "securePassword123"
		validRole  = "employee"
		userID     = uuid.New()
	)

	defer mc.Finish()

	tests := []struct {
		name            string
		args            args
		wantStatus      int
		wantResponse    interface{}
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case - valid registration",
			args: args{
				ctx: ctx,
				req: func() *http.Request {
					reqBody := fmt.Sprintf(`{"email":"%s","password":"%s","role":"%s"}`, validEmail, validPass, validRole)
					req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(reqBody)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusCreated,
			wantResponse: &dto.RegisterResponse{
				ID:    userID.String(),
				Email: validEmail,
				Role:  validRole,
			},
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(mc)
				mock.RegisterMock.Expect(ctx, &model.UserRegistration{
					Info: model.UserInfo{
						Email: validEmail,
						Role:  model.Role(validRole),
					},
					Password: validPass,
				}).Return(&model.User{
					ID: userID,
					Info: model.UserInfo{
						Email: validEmail,
						Role:  model.Role(validRole),
					},
				}, nil)
				return mock
			},
		},
		{
			name: "error case - invalid JSON",
			args: args{
				ctx: ctx,
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(`invalid json`)))
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
			name: "error case - missing email",
			args: args{
				ctx: ctx,
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(`{"password":"pass","role":"employee"}`)))
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
			name: "error case - invalid email format",
			args: args{
				ctx: ctx,
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(`{"email":"invalid","password":"pass","role":"employee"}`)))
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
			name: "error case - service internal error",
			args: args{
				ctx: ctx,
				req: func() *http.Request {
					reqBody := fmt.Sprintf(`{"email":"%s","password":"%s","role":"%s"}`, validEmail, validPass, validRole)
					req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(reqBody)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusInternalServerError,
			wantResponse: &dto.ErrorResponse{
				Message: "internal server error",
			},
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(mc)
				mock.RegisterMock.Expect(ctx, &model.UserRegistration{
					Info: model.UserInfo{
						Email: validEmail,
						Role:  model.Role(validRole),
					},
					Password: validPass,
				}).Return(nil, errors.New("database error"))
				return mock
			},
		},
		{
			name: "error case - invalid role",
			args: args{
				ctx: ctx,
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(`{"email":"test@test.com","password":"pass","role":"invalid"}`)))
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
			authServiceMock := tt.authServiceMock(mc)
			authImpl := auth.NewImplementation(authServiceMock)

			router := gin.New()
			router.POST("/register", authImpl.Register)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, tt.args.req)

			require.Equal(t, tt.wantStatus, w.Code)
			require.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

			switch expected := tt.wantResponse.(type) {
			case *dto.RegisterResponse:
				var actual dto.RegisterResponse
				err := json.Unmarshal(w.Body.Bytes(), &actual)
				require.NoError(t, err)
				require.Equal(t, expected.ID, actual.ID)
				require.Equal(t, expected.Email, actual.Email)
				require.Equal(t, expected.Role, actual.Role)

			case *dto.ErrorResponse:
				var errResp dto.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &errResp)
				require.NoError(t, err)
				require.Equal(t, expected.Message, errResp.Message)
			}
		})
	}
}
