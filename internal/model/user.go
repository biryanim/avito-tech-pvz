package model

import (
	"github.com/google/uuid"
	"github.com/dgrijalva/jwt-go"
)

type UserInfo struct {
	Email string
	Role  string
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
	Email string
}
