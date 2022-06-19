package conec

import "database/sql"

func ConnectSQLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../final_project.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}
