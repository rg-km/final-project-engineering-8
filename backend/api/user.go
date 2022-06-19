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
	api.AllowOrigin(c)
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
		ID:       int64(dataUser.UserID),
		Username: cred.Username,
		Role:     dataUser.Role,
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
		"status":  "true",
		"code":    http.StatusOK,
		"message": "login success",
		"token":   tokenString,
	})
}

func (api *API) Register(c *gin.Context) {
	api.AllowOrigin(c)
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
	api.AllowOrigin(c)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
	})

	c.JSON(http.StatusOK, gin.H{
		"status":  "true",
		"code":    http.StatusOK,
		"message": "logout success",
	})
}

func (api *API) GetTeachers(c *gin.Context) {
	api.AllowOrigin(c)

	var (
		page    int
		perPage int
		offset  int
		total   int
		message string
		isError bool
	)

	params := c.Request.URL.Query()

	//Convert string to int
	_, err := fmt.Sscan(params.Get("per_page"), &perPage)
	_, err = fmt.Sscan(params.Get("page"), &page)

	if err != nil && err.Error() != "EOF" { // chcek jika tidak error dan jika errornya karena mengirim param yg tidak bisa di convert ke int
		c.JSON(http.StatusBadRequest, Result{
			Status:  false,
			Code:    http.StatusBadRequest,
			Message: "Throw a param with the value convertible to a number, ERROR: " + err.Error(),
			Data:    []string{},
		})
		return
	}

	// set default value for (optional) params
	if perPage == 0 {
		perPage = 50
	}

	if page == 0 {
		page = 1
	}

	offset = (page - 1) * perPage

	defer func() {
		if isError {
			c.JSON(http.StatusInternalServerError, Result{
				Status:  false,
				Code:    http.StatusInternalServerError,
				Message: "Failed to fetch teachers, ERROR: " + message,
				Data:    nil,
			})
			return
		}
	}()

	teachers, err := api.userRepo.FetchAllTeachers(perPage, offset)
	if err != nil {
		isError = true
		message = err.Error()
		return
	}

	total, err = api.userRepo.GetNumberofTeacherRow()
	if err != nil {
		isError = true
		message = err.Error()
		return
	}

	c.JSON(http.StatusOK, Result{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Success",
		Data:    teachers,
		Pagination: &Pagination{
			Total:     total,
			Page:      page,
			PerPage:   perPage,
			TotalPage: 10,
		},
	})
}

func (api *API) UpdateTeacherById(c *gin.Context) {
	api.AllowOrigin(c)
	id := c.Param("id")

	var (
		isError bool
		message string
	)

	defer func() {
		if isError {
			c.JSON(http.StatusInternalServerError, Result{
				Status:  false,
				Code:    http.StatusInternalServerError,
				Message: "Failed to update teacher, ERROR: " + message,
				Data:    nil,
			})
			return
		}
	}()

	payload, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, Result{
			Status:  false,
			Code:    http.StatusInternalServerError,
			Message: "Bad Request: " + message,
			Data:    nil,
		})

		return
	}

	var teacherData map[string]interface{}
	json.Unmarshal(payload, &teacherData)

	err = api.userRepo.UpdateTeacher(id, teacherData)
	if err != nil {
		isError = true
		message = err.Error()
		return
	}

	updatedTeacher, err := api.userRepo.GetTeacherByID(id)
	if err != nil {
		isError = true
		message = err.Error()
		return
	}

	c.JSON(http.StatusOK, Result{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Success",
		Data:    updatedTeacher,
	})
}

func (api *API) DeleteTeacher(c *gin.Context) {
	id := c.Param("id")

	var (
		isError bool
		message string
		code    int
	)

	defer func() {
		if isError {
			c.JSON(code, Result{
				Status:  false,
				Code:    code,
				Message: message,
				Data:    nil,
			})
		}
	}()

	code, err := api.userRepo.DeleteTeacherByUserID(id)

	if err != nil {
		isError = true
		message = err.Error()
		return
	}

	c.JSON(http.StatusOK, Result{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}
