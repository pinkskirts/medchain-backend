package login

import "github.com/golang-jwt/jwt/v4"

type responseLogin struct {
	Token string    `json:"token"`
	Error typeError `json:"errror"`
}
type typeError struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

type BodyJWT struct {
	Email       string   `json:"email"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}
