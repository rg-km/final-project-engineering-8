package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	repo "github.com/rg-km/final-project-engineering-8/backend/repository"
)

type API struct {
	userRepo repo.UserRepository
	gin      *gin.Engine
}

func NewAPI(userRepo repo.UserRepository) *API {
	gin := gin.Default()
	api := &API{
		userRepo: userRepo,
		gin:      gin,
	}

	// gin.POST("/login/siswa", api.LoginSiswa)

	user := gin.Group("/user")
	{
		user.POST("/login", api.LoginUser)
	}

	return api
}

func (api *API) Handler() *gin.Engine {
	return api.gin
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080/")
	api.Handler().Run(":8080")
}
