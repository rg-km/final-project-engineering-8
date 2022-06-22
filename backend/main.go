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

	userRepo := repo.NewUserRepository(db)
	mainApi := api.NewAPI(*userRepo)
	mainApi.Start()
}
