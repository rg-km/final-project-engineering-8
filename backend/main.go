package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/rg-km/final-project-engineering-8/backend/api"
	con "github.com/rg-km/final-project-engineering-8/backend/database/connection"
	repo "github.com/rg-km/final-project-engineering-8/backend/repository"
)

func main() {
	db, err := con.ConnectSQLite()
	if err != nil {
		panic(err)
	}

	siswaRepo := repo.NewSiswaRepository(db)
	adminRepo := repo.NewAdminRepository(db)
	guruRepo := repo.NewGuruRepository(db)
	mainApi := api.NewAPI(*siswaRepo, *adminRepo, *guruRepo)
	mainApi.Start()
}
