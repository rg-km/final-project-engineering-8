package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (api *API) LoginUser(c *gin.Context) {
	var cred Credentials
	err := json.NewDecoder(c.Request.Body).Decode(&cred)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid request body",
		})
		return
	}

	if cred.Username == "" && cred.Password == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "username dan password tidak boleh kosong",
		})
		return
	} else if cred.Username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "username tidak boleh kosong",
		})
		return
	} else if cred.Password == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "password tidak boleh kosong",
		})
		return
	}

	resp, err := api.userRepo.LoginUser(cred.Username, cred.Password)
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
	} else if dataUser.Username != cred.Username {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "user credential invalid",
		})
		return
	}

	expirationTime := time.Now().Local().Add((5 * time.Minute))

	claims := &Claims{
		Username: cred.Username,
		Role:     "user",
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
		"data":    dataUser,
	})
}

func (api *API) Register(c *gin.Context) {
	var register Register
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// register.Username = c.PostForm("username")
	data, err := api.userRepo.Register(register.Username, register.Password, register.Nama, register.Alamat, register.NoHp, register.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": data,
	})
}

func (api *API) Logout(c *gin.Context) {
	//logout
	_, err := c.Request.Cookie("token")
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "anda belum login",
		})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
	})

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "logout success",
	})
}