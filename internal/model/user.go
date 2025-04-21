package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrorInvalidToken    = errors.New("invalid token")
	ErrorInvalidRole     = errors.New("invalid role")
	ErrorInvalidPassword = errors.New("invalid password")
	ErrorUserNotFound    = errors.New("user not found")
)

type UserInfo struct {
	Email string
	Role  Role
}
type UserRegistration struct {
	Info     UserInfo
	Password string
}

type User struct {
	ID       uuid.UUID
	Info     UserInfo
	Password string
}

type UserLoginInfo struct {
	Email    string
	Password string
}

type UserClaims struct {
	jwt.StandardClaims
	Role Role
}

type Role string

const (
	RoleEmployee  Role = "employee"
	RoleModerator Role = "moderator"
)

func (r Role) IsValid() bool {
	switch r {
	case RoleEmployee, RoleModerator:
		return true
	}
	return false
}
