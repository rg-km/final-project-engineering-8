package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	repo "github.com/rg-km/final-project-engineering-8/backend/repository"
)

type API struct {
	siswaRepo repo.SiswaRepository
	gin       *gin.Engine
}

func NewAPI(siswaRepo repo.SiswaRepository) *API {
	gin := gin.Default()
	api := &API{
		siswaRepo: siswaRepo,
		gin:       gin,
	}

	gin.POST("/login/siswa", api.LoginSiswa)

	return api
}

func (api *API) Handler() *gin.Engine {
	return api.gin
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080/")
	api.Handler().Run(":8080")
}
