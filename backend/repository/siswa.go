package repository

import "database/sql"

type SiswaRepository struct {
	db *sql.DB
}

func NewSiswaRepository(db *sql.DB) *SiswaRepository {
	return &SiswaRepository{db: db}
}

func (u *SiswaRepository) LoginSiswa(username string, password string) (*Siswa, error) {
	sqlStatement := `SELECT * FROM siswa WHERE username = ? AND password = ?;`

	rows, err := u.db.Query(sqlStatement, username, password)
	if err != nil {
		return nil, err
	}

	var user Siswa
	for rows.Next() {
		err = rows.Scan(&user.SiswaID, &user.Username, &user.Password, &user.Nama, &user.Alamat, &user.RoleID)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}
