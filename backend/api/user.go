package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (api *API) LoginUser(c *gin.Context) {
	api.AllowOrigin(c)
	var cred Credentials
	err := json.NewDecoder(c.Request.Body).Decode(&cred)

	if err != nil {
		c.JSON(http.StatusUnauthorized, Result{
			Status:  false,
			Code:    http.StatusUnauthorized,
			Message: "Invalid request body",
		})
		return
	}

	if cred.Username == "" && cred.Password == "" {
		c.JSON(http.StatusUnauthorized, Result{
			Status:  false,
			Code:    http.StatusUnauthorized,
			Message: "username dan password tidak boleh kosong",
		})
		return
	} else if cred.Username == "" {
		c.JSON(http.StatusUnauthorized, Result{
			Status:  false,
			Code:    http.StatusUnauthorized,
			Message: "username tidak boleh kosong",
		})
		return
	} else if cred.Password == "" {
		c.JSON(http.StatusUnauthorized, Result{
			Status:  false,
			Code:    http.StatusUnauthorized,
			Message: "password tidak boleh kosong",
		})
		return
	}

	resp, err := api.userRepo.LoginUser(cred.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Result{
			Status:  false,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	dataUser := *resp

	if err := bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(cred.Password)); err != nil {

		c.JSON(http.StatusUnauthorized, Result{
			Status:  false,
			Code:    http.StatusUnauthorized,
			Message: "password salah",
		})
		return
	} else if dataUser.Username != cred.Username {
		c.JSON(http.StatusUnauthorized, Result{
			Status:  false,
			Code:    http.StatusUnauthorized,
			Message: "user credential invalid",
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
		c.JSON(http.StatusInternalServerError, Result{
			Status:  false,
			Code:    http.StatusInternalServerError,
			Message: "password tidak boleh kosong",
		})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	c.Header("Authorization", "Bearer "+tokenString)
	c.JSON(http.StatusOK, gin.H{
		"status":       true,
		"code":         http.StatusOK,
		"id":           dataUser.UserID,
		"name":         dataUser.Nama,
		"role":         dataUser.Role,
		"profile_pict": dataUser.ProfilePict,
		"message":      "login success",
		"token":        tokenString,
	})
}

func (api *API) StudentRegister(c *gin.Context) {
	api.AllowOrigin(c)
	var register Register
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// register.Username = c.PostForm("username")
	password, _ := bcrypt.GenerateFromPassword([]byte(register.Password), 10)
	strPassword := string(password)

	data, err := api.userRepo.StudentRegister(register.Username, strPassword, register.Nama, register.Alamat, register.NoHp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Result{
			Status:  false,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Result{
		Status:  true,
		Code:    http.StatusOK,
		Message: "registration success",
		Data:    data,
	})
}

func (api *API) TeacherRegister(c *gin.Context) {
	api.AllowOrigin(c)
	var register Register
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// register.Username = c.PostForm("username")
	password, err := HashPassword(register.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Result{
			Status:  false,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	register.Password = password

	data, err := api.userRepo.TeacherRegister(register.Username, register.Password, register.Nama, register.Alamat, register.NoHp, register.Deskripsi, register.Biaya, register.JenjangID, register.PelajaranID, register.KategoriID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, Result{
			Status:  false,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Result{
		Status:  true,
		Code:    http.StatusOK,
		Message: "registration success",
		Data:    data,
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (api *API) Logout(c *gin.Context) {
	//logout
	api.AllowOrigin(c)

	c.Header("Authorization", "")

	c.JSON(http.StatusOK, Result{
		Status:  true,
		Code:    http.StatusOK,
		Message: "logout success",
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

	totalPage := 1
	if total > perPage {
		totalPage = int(math.Ceil(float64(total) / float64(perPage)))
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
			TotalPage: totalPage,
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

	teacherData := make(map[string]interface{})
	json.Unmarshal(payload, &teacherData)

	if teacherData["profile_pict"] != nil {

		isLinkFormat := CheckIsLinkFormat(teacherData["profile_pict"].(string))

		if !isLinkFormat {
			// Jika bukan link, berarti base64 string
			// Maka store base64 ke image bucket untuk generate link
			generatedLink, err := StoreBase64ToBucketImage(teacherData["profile_pict"].(string))
			fmt.Println("Generated Link ==>", generatedLink)

			if err != nil {
				isError = true
				message = err.Error()
				return
			}

			// pass generated link to payload
			teacherData["profile_pict"] = generatedLink
		}

	}

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

func (api *API) GetTeacherByUserID(c *gin.Context) {
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
				Message: message,
				Data:    nil,
			})
		}
	}()

	teacher, err := api.userRepo.GetTeacherByID(id)

	if err != nil {
		isError = true
		message = err.Error()
		return
	}

	c.JSON(http.StatusOK, Result{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Success",
		Data:    teacher,
	})
}

func (api *API) GetStudentLogin(c *gin.Context) {
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

	username := claims.Username
	student, err := api.userRepo.GetStudentProfile(username)
	if err != nil {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, Result{
			Status:  false,
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Result{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Success",
		Data:    student,
	})
}

func (api *API) UpdateStudentById(c *gin.Context) {
	api.AllowOrigin(c)
	var user Users
	id := c.Param("id")

	convertIDInt, _ := strconv.Atoi(id)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	strPassword := string(password)
	data, err := api.userRepo.UpdateStudentById(convertIDInt, user.Username, strPassword, user.Nama, user.Alamat, user.NoHp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Result{
			Status:  false,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Result{
		Status:  true,
		Code:    http.StatusOK,
		Message: "update successfully",
		Data:    data,
	})
}

func CheckIsLinkFormat(text string) bool {
	regex, _ := regexp.Compile(`http(s*)://`)
	return regex.MatchString(text)
}

func StoreBase64ToBucketImage(base64String string) (generatedLink string, err error) {
	client := &http.Client{}

	var jsonPayload = []byte(fmt.Sprintf("{\"image\": \"%s\"}", base64String))
	r, fail := http.NewRequest("POST", "https://halloguru.herokuapp.com/v1/image", bytes.NewBuffer(jsonPayload))
	r.Header.Set("Content-Type", "application/json")

	resp, _ := client.Do(r)
	if fail != nil {
		return generatedLink, fail
	}

	defer resp.Body.Close()

	var body map[string]interface{}
	fail = json.NewDecoder(resp.Body).Decode(&body)

	if fail != nil {
		return generatedLink, fail
	}

	if resp.StatusCode != 200 {
		return generatedLink, errors.New(fmt.Sprint("Failed store image to bucket, ERROR: ", body["message"]))
	}

	fmt.Println("Response ==>", body)

	return body["link"].(string), nil
}
