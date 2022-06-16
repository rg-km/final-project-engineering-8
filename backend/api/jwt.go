package api

import "github.com/golang-jwt/jwt/v4"

//inputan user
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nama     string `json:"nama"`
	Alamat   string `json:"alamat"`
	NoHp     string `json:"no_hp"`
	Role     string `json:"role"`
}

//jwt token
var jwtKey = []byte("key")

type Claims struct {
	ID       int64 `json:"id"`
	Username string
	Role     string
	jwt.StandardClaims
}
