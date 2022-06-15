package api_test

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rg-km/final-project-engineering-8/backend/api"
	repo "github.com/rg-km/final-project-engineering-8/backend/repository"
)

var _ = Describe("Api", func() {
	var mainApi *api.API

	BeforeEach(func() {
		db, err := sql.Open("sqlite3", "../database/final_project.db")
		if err != nil {
			panic(err)
		}

		userRepo := repo.NewUserRepository(db)
		mainApi = api.NewAPI(*userRepo)
	})

	Describe("/login", func() {
		When("login with valid credentials", func() {
			It("should return 200", func() {
				user := `{"username":"ucup","password":"ucup123"}`
				resp := httptest.NewRecorder()
				req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(user))
				if err != nil {
					panic(err)
				}
				mainApi.Handler().ServeHTTP(resp, req)
				req.Header.Set("Content-Type", "application/json")
				Expect(resp.Code).To(Equal(http.StatusOK))
			})
		})
		When("login with invalid credentials", func() {
			It("should return 401 because username is empty", func() {
				user := `{"username":"","password":"ucup123"}`
				resp := httptest.NewRecorder()
				req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(user))
				if err != nil {
					panic(err)
				}
				mainApi.Handler().ServeHTTP(resp, req)
				req.Header.Set("Content-Type", "application/json")
				Expect(resp.Code).To(Equal(http.StatusUnauthorized))
			})
			It("should return 401 because password is empty", func() {
				user := `{"username":"ucup","password":""}`
				resp := httptest.NewRecorder()
				req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(user))
				if err != nil {
					panic(err)
				}
				mainApi.Handler().ServeHTTP(resp, req)
				req.Header.Set("Content-Type", "application/json")
				Expect(resp.Code).To(Equal(http.StatusUnauthorized))
			})
			It("should return 401 because username and password is empty", func() {
				user := `{"username":"","password":""}`
				resp := httptest.NewRecorder()
				req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(user))
				if err != nil {
					panic(err)
				}
				mainApi.Handler().ServeHTTP(resp, req)
				req.Header.Set("Content-Type", "application/json")
				Expect(resp.Code).To(Equal(http.StatusUnauthorized))
			})
			It("should return 401 because password is different with data in database", func() {
				user := `{"username":"ucup","password":"ucup"}`
				resp := httptest.NewRecorder()
				req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(user))
				if err != nil {
					panic(err)
				}
				mainApi.Handler().ServeHTTP(resp, req)
				req.Header.Set("Content-Type", "application/json")
				Expect(resp.Code).To(Equal(http.StatusUnauthorized))
			})
			It("should return 401 because username is different with data in database", func() {
				user := `{"username":"ucup123","password":"ucup123"}`
				resp := httptest.NewRecorder()
				req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(user))
				if err != nil {
					panic(err)
				}
				mainApi.Handler().ServeHTTP(resp, req)
				req.Header.Set("Content-Type", "application/json")
				Expect(resp.Code).To(Equal(http.StatusUnauthorized))
			})
		})
	})
})
