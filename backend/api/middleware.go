package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (api *API) AllowOrigin(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080/")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*") //for development purpose
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
}

func (api *API) AuthMiddleWare(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		api.AllowOrigin(c)
		token, err := c.Request.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				c.Writer.WriteHeader(http.StatusUnauthorized)
				c.JSON(http.StatusUnauthorized, gin.H{
					"status":  "false",
					"code":    http.StatusUnauthorized,
					"message": "anda belum login",
				})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "false",
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		tknStr := token.Value

		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Writer.WriteHeader(http.StatusUnauthorized)
				c.JSON(http.StatusUnauthorized, gin.H{
					"status":  "false",
					"code":    http.StatusUnauthorized,
					"message": err.Error(),
				})
				return
			}
			c.Writer.WriteHeader(http.StatusBadRequest)
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "false",
				"code":    http.StatusUnauthorized,
				"message": err.Error(),
			})
			return
		}

		if !tkn.Valid {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "false",
				"code":    http.StatusUnauthorized,
				"message": "token invalid!",
			})
			return
		}

		ctx := context.WithValue(c, "username", claims.Username)
		ctx = context.WithValue(ctx, "role", claims.Role)
		next(c)
	}
}

func (api *API) MiddlewareSiswa(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		api.AllowOrigin(c)
		token, _ := c.Request.Cookie("token")

		tknStr := token.Value

		claims := &Claims{}

		_, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Writer.WriteHeader(http.StatusUnauthorized)
				c.JSON(http.StatusUnauthorized, gin.H{
					"status":  "false",
					"code":    http.StatusUnauthorized,
					"message": err.Error(),
				})
				return
			}
			c.Writer.WriteHeader(http.StatusBadRequest)
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "false",
				"code":    http.StatusUnauthorized,
				"message": err.Error(),
			})
			return
		}

		if claims.Role != "siswa" {
			c.Writer.WriteHeader(http.StatusForbidden)
			c.JSON(http.StatusForbidden, gin.H{
				"status":  "false",
				"code":    http.StatusForbidden,
				"message": "forbidden access",
			})
			return
		}
		next(c)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
