package repository

import "database/sql"

type GuruRepository struct {
	db *sql.DB
}

func NewGuruRepository(db *sql.DB) *GuruRepository {
	return &GuruRepository{db: db}
}

func (u *GuruRepository) LoginGuru(username string, password string) (*Guru, error) {
	sqlStatement := `SELECT * FROM guru WHERE username = ? AND password = ?;`

	rows, err := u.db.Query(sqlStatement, username, password)
	if err != nil {
		return nil, err
	}

	var guru Guru
	for rows.Next() {
		err = rows.Scan(&guru.GuruID, &guru.Username, &guru.Password, &guru.Nama, &guru.Alamat, &guru.NoHp, &guru.Deskripsi, &guru.Biaya, &guru.RoleID, &guru.JenjangID, &guru.PelajaranID, &guru.KategoriID)
		if err != nil {
			return nil, err
		}
	}

	return &guru, nil
}
