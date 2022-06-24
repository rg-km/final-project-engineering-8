package api

import (
	"net/http"
	"strings"

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
	// return func(c *gin.Context) {
	// 	err := GetAuthentication(c)

	// 	if err != nil {
	// 		c.JSON(401, gin.H{
	// 			"status":  401,
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	next(c)
	// }
	// return func(c *gin.Context) {
	// 	api.AllowOrigin(c)
	// 	token, err := c.Request.Cookie("token")
	// 	if err != nil {
	// 		if err == http.ErrNoCookie {
	// 			c.Writer.WriteHeader(http.StatusUnauthorized)
	// 			c.JSON(http.StatusUnauthorized, gin.H{
	// 				"status":  "false",
	// 				"code":    http.StatusUnauthorized,
	// 				"message": "anda belum login",
	// 			})
	// 			return
	// 		}
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"status":  "false",
	// 			"code":    http.StatusBadRequest,
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	tknStr := token.Value

	// 	claims := &Claims{}

	// 	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
	// 		return jwtKey, nil
	// 	})

	// 	if err != nil {
	// 		if err == jwt.ErrSignatureInvalid {
	// 			c.Writer.WriteHeader(http.StatusUnauthorized)
	// 			c.JSON(http.StatusUnauthorized, gin.H{
	// 				"status":  "false",
	// 				"code":    http.StatusUnauthorized,
	// 				"message": err.Error(),
	// 			})
	// 			return
	// 		}
	// 		c.Writer.WriteHeader(http.StatusBadRequest)
	// 		c.JSON(http.StatusUnauthorized, gin.H{
	// 			"status":  "false",
	// 			"code":    http.StatusUnauthorized,
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	if !tkn.Valid {
	// 		c.Writer.WriteHeader(http.StatusUnauthorized)
	// 		c.JSON(http.StatusUnauthorized, gin.H{
	// 			"status":  "false",
	// 			"code":    http.StatusUnauthorized,
	// 			"message": "token invalid!",
	// 		})
	// 		return
	// 	}

	// 	ctx := context.WithValue(c, "username", claims.Username)
	// 	ctx = context.WithValue(ctx, "role", claims.Role)
	// 	next(c)
	// }
	return func(c *gin.Context) {
		api.AllowOrigin(c)
		var token string
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.Writer.WriteHeader(http.StatusUnauthorized)

			c.JSON(http.StatusUnauthorized, Result{
				Status:  false,
				Code:    http.StatusUnauthorized,
				Message: "anda belum login",
			})
			return
		}
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			token = bearerToken[1]
		} else {
			token = ""
		}

		if token == "" {
			c.JSON(401, gin.H{
				"status":  false,
				"code":    401,
				"message": "Token Not Valid",
			})
			return
		}
		claims := &Claims{}

		parseTkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Writer.WriteHeader(http.StatusUnauthorized)

				c.JSON(http.StatusUnauthorized, Result{
					Status:  false,
					Code:    http.StatusUnauthorized,
					Message: err.Error(),
				})
				return
			}
			c.Writer.WriteHeader(http.StatusBadRequest)

			c.JSON(http.StatusUnauthorized, Result{
				Status:  false,
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}

		if !parseTkn.Valid {
			c.Writer.WriteHeader(http.StatusUnauthorized)

			c.JSON(http.StatusUnauthorized, Result{
				Status:  false,
				Code:    http.StatusUnauthorized,
				Message: "token invalid!",
			})
			return
		}

		c.Set("claims", claims)

		next(c)
	}
}

func (api *API) MiddlewareSiswa(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		api.AllowOrigin(c)
		// token, _ := c.Request.Cookie("token")

		// tknStr := token.Value

		// claims := &Claims{}

		// _, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		// 	return jwtKey, nil
		// })

		// if err != nil {
		// 	if err == jwt.ErrSignatureInvalid {
		// 		c.Writer.WriteHeader(http.StatusUnauthorized)
		// 		c.JSON(http.StatusUnauthorized, gin.H{
		// 			"status":  "false",
		// 			"code":    http.StatusUnauthorized,
		// 			"message": err.Error(),
		// 		})
		// 		return
		// 	}
		// 	c.Writer.WriteHeader(http.StatusBadRequest)
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"status":  "false",
		// 		"code":    http.StatusUnauthorized,
		// 		"message": err.Error(),
		// 	})
		// 	return
		// }

		var token string
		authHeader := c.Request.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			token = bearerToken[1]
		} else {
			token = ""
		}

		if token == "" {
			c.JSON(401, gin.H{
				"status":  false,
				"code":    401,
				"message": "Token Not Valid",
			})
			return
		}
		claims := &Claims{}

		parseTkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Writer.WriteHeader(http.StatusUnauthorized)

				c.JSON(http.StatusUnauthorized, Result{
					Status:  false,
					Code:    http.StatusUnauthorized,
					Message: err.Error(),
				})
				return
			}
			c.Writer.WriteHeader(http.StatusBadRequest)

			c.JSON(http.StatusUnauthorized, Result{
				Status:  false,
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}

		if !parseTkn.Valid {
			c.Writer.WriteHeader(http.StatusUnauthorized)

			c.JSON(http.StatusUnauthorized, Result{
				Status:  false,
				Code:    http.StatusUnauthorized,
				Message: "token invalid",
			})
			return
		}

		if claims.Role != "siswa" {
			c.Writer.WriteHeader(http.StatusForbidden)

			c.JSON(http.StatusForbidden, Result{
				Status:  false,
				Code:    http.StatusForbidden,
				Message: "forbidden access",
			})
			return
		}
		next(c)
	}

}

func (api *API) MiddlewareGuru(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		api.AllowOrigin(c)

		var token string
		authHeader := c.Request.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			token = bearerToken[1]
		} else {
			token = ""
		}

		if token == "" {
			c.JSON(401, gin.H{
				"status":  false,
				"code":    401,
				"message": "Token Not Valid",
			})
			return
		}
		claims := &Claims{}

		parseTkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Writer.WriteHeader(http.StatusUnauthorized)

				c.JSON(http.StatusUnauthorized, Result{
					Status:  false,
					Code:    http.StatusUnauthorized,
					Message: err.Error(),
				})
				return
			}
			c.Writer.WriteHeader(http.StatusBadRequest)

			c.JSON(http.StatusUnauthorized, Result{
				Status:  false,
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}

		if !parseTkn.Valid {
			c.Writer.WriteHeader(http.StatusUnauthorized)

			c.JSON(http.StatusUnauthorized, Result{
				Status:  false,
				Code:    http.StatusUnauthorized,
				Message: "token invalid!",
			})
			return
		}

		if claims.Role != "guru" {
			c.Writer.WriteHeader(http.StatusForbidden)

			c.JSON(http.StatusForbidden, Result{
				Status:  false,
				Code:    http.StatusForbidden,
				Message: "forbidden access",
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
