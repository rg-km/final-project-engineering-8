package api

import "github.com/golang-jwt/jwt/v4"

//inputan user
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//jwt token
var jwtKey = []byte("key")

type Claims struct {
	Username string
	Role     string
	jwt.StandardClaims
}
