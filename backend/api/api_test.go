package api_test

import (
	"bytes"
	"database/sql"
	"fmt"
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
		db, err := sql.Open("sqlite3", "../database/final_project2.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS user (
				UserID integer not null primary key AUTOINCREMENT,
				username varchar(255) not null,
				password varchar(255) not null,
				nama varchar(255) not null,
				alamat varchar(255) not null,
				noHp varchar(255) not null,
				role varchar(255) not null,
				profilePict TEXT NOT NULL DEFAULT "https://halloguru.herokuapp.com/bucket-image/halloguru/1655885462662157252.png"
			);
		
			CREATE TABLE IF NOT EXISTS jenjang (
				JenjangID integer not null primary key AUTOINCREMENT,
				jenjang varchar(255) not null
			);
		
			CREATE TABLE IF NOT EXISTS pelajaran (
				PelajaranID integer not null primary key AUTOINCREMENT,
				pelajaran varchar(255) not null
			);
		
			CREATE TABLE IF NOT EXISTS kategori (
				KategoriID integer not null primary key AUTOINCREMENT,
				kategori varchar(255) not null
			);
		
			CREATE TABLE IF NOT EXISTS info_guru (
				InfoGuruID integer not null primary key AUTOINCREMENT,
				deskripsi varchar(255) not null,
				biaya varchar(255) not null,
				ratting varchar(255) not null,
				UserID integer not null,
				JenjangID integer not null,
				PelajaranID integer not null,
				KategoriID integer not null,
				FOREIGN KEY (UserID) REFERENCES user(UserID),
				FOREIGN KEY (JenjangID) REFERENCES jenjang(JenjangID),
				FOREIGN KEY (PelajaranID) REFERENCES pelajaran(PelajaranID),
				FOREIGN KEY (KategoriID) REFERENCES kategori(KategoriID)
			);

			INSERT INTO user (username, password, nama, alamat, noHP,role) 
			VALUES
			("ruli", "$2a$10$B/ccuPXMZ3/rLdHQiW7M8.nKFa4Rcd9n6JPjavrPNhH59moghfb2G", "ruliansyah", "Jl Bojong Soang No.1, Kota Bandung", "0812341234", "siswa");
		`)

		if err != nil {
			panic(err)
		}

		userRepo := repo.NewUserRepository(db)
		mainApi = api.NewAPI(*userRepo)

	})

	AfterEach(func() {
		db, err := sql.Open("sqlite3", "../database/final_project2.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
		DROP TABLE IF EXISTS user;
		DROP TABLE IF EXISTS jenjang;
		DROP TABLE IF EXISTS kategori;
		DROP TABLE IF EXISTS info_guru;`)

		if err != nil {
			panic(err)
		}
	})

	Describe("/login", func() {
		When("username and password are correct", func() {
			It("should return 200", func() {
				user := `
				{
					"username": "ruli",
					"password": "ruli123"
				}
				`
				resp := httptest.NewRecorder()
				req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(user))
				if err != nil {
					panic(err)
				}
				fmt.Println("RESP", resp)
				mainApi.Handler().ServeHTTP(resp, req)
				req.Header.Set("Content-Type", "application/json")
				Expect(resp.Code).To(Equal(http.StatusOK))
			})
		})

		When("username is empty but password is correct", func() {
			It("should return 401 because username is empty", func() {
				user := `{"username":"","password":"ruli123"}`
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

		When("username is correct but password is empty", func() {
			It("should return 401 because password is empty", func() {
				user := `{"username":"ruli","password":""}`
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

		When("username and password is empty", func() {
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
		})

		When("username is correct but password is incorrect", func() {
			It("should return 401 because password is different with data in database", func() {
				user := `{"username":"ruli","password":"ruli"}`
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

		When("username is empty but password is correct", func() {
			It("should return 401 because username is different with data in database", func() {
				user := `{"username":"ruli123","password":"ruli123"}`
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
