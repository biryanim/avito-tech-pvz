package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
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
