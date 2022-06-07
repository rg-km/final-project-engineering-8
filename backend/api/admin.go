package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

//inputan user
// type Credentials struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// //jwt token
// var jwtKey = []byte("key")

// type Claims struct {
// 	Username string
// 	Role     string
// 	jwt.StandardClaims
// }

func (api *API) LoginAdmin(c *gin.Context) {
	var cred Credentials
	err := json.NewDecoder(c.Request.Body).Decode(&cred)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid request body",
		})
		return
	}

	resp, err := api.adminRepo.LoginAdmin(cred.Username, cred.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	dataUser := *resp

	if dataUser.Password != cred.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "user credential invalid",
		})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: cred.Username,
		Role:     "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "login success",
	})
}