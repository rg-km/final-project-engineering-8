package api

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	repo "github.com/rg-km/final-project-engineering-8/backend/repository"
)

type API struct {
	userRepo repo.UserRepository
	gin      *gin.Engine
}

func NewAPI(userRepo repo.UserRepository) *API {
	gin := gin.Default()
	gin.Use(CORSMiddleware())
	api := &API{
		userRepo: userRepo,
		gin:      gin,
	}

	gin.POST("/login", api.LoginUser)
	gin.POST("/register/student", api.StudentRegister)
	gin.POST("/register/teacher", api.TeacherRegister)
	gin.POST("/logout", api.AuthMiddleWare(api.Logout))
	gin.PUT("v1/teacher/:id", api.AuthMiddleWare(api.UpdateTeacherById))
	gin.GET("/v1/teachers", api.AuthMiddleWare(api.MiddlewareSiswa(api.GetTeachers)))

	return api
}

func (api *API) Handler() *gin.Engine {
	return api.gin
}

func (api *API) Start() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file. ERROR:", err)
	}
	// port := os.Getenv("PORT")

	// if port == "" {
	// 	log.Fatal("$PORT must be set")
	// }

	HOST := os.Getenv("API_HOST")
	PORT := os.Getenv("API_PORT")

	fmt.Printf("starting web server at http://%v:%v/", HOST, PORT)
	api.Handler().Run(":" + PORT)
}
