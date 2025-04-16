package model

import "github.com/dgrijalva/jwt-go"

type UserInfo struct {
	Email string
	Role  string
}
type UserRegistration struct {
	Info     UserInfo
	Password string
}

type User struct {
	ID       string
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

/*
 User:
      type: object
      properties:
        id:
          type: string
          format: uuid
        email:
          type: string
          format: email
        role:
          type: string
          enum: [employee, moderator]
      required: [email, role]

/register:
	 properties:
                email:
                  type: string
                  format: email
                password:
                  type: string
                role:
                  type: string
                  enum: [employee, moderator]
              required: [email, password, role]
*/
