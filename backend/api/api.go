package api

import (
	"github.com/gin-gonic/gin"
	repo "github.com/rg-km/final-project-engineering-8/backend/repository"
)

type API struct {
	userRepo repo.UserRepository
	gin      *gin.Engine
}

func NewAPI(userRepo repo.UserRepository) *API {
	// gin.SetMode(gin.ReleaseMode)
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
	gin.GET("/v1/teachers", api.GetTeachers)
	gin.DELETE("/v1/teacher/:id", api.AuthMiddleWare(api.DeleteTeacher))
	gin.GET("/v1/teacher/:id", api.AuthMiddleWare(api.GetTeacherByUserID))

	gin.GET("/v1/student/update", api.AuthMiddleWare(api.MiddlewareSiswa(api.GetStudentLogin)))
	gin.PUT("/v1/student/update/:id", api.AuthMiddleWare(api.MiddlewareSiswa(api.UpdateStudentById)))
	gin.DELETE("/v1/student", api.AuthMiddleWare(api.MiddlewareSiswa(api.DeleteStudent)))

	return api
}

func (api *API) Handler() *gin.Engine {
	return api.gin
}

func (api *API) Start() {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file. ERROR:", err)
	// }
	// // port := os.Getenv("PORT")

	// // if port == "" {
	// // 	log.Fatal("$PORT must be set")
	// // }

	// HOST := os.Getenv("API_HOST")
	// PORT := os.Getenv("API_PORT")

	// fmt.Printf("starting web server at http://%v:%v/", HOST, PORT)
	api.Handler().Run(":" + "8080")
}
