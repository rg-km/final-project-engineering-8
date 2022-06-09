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

func (u *SiswaRepository) FetchAllUsers() ([]Siswa, error) {
	sqlStmt := `SELECT * FROM siswa`
	rows, err := u.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	siswae := []Siswa{}
	for rows.Next() {
		var siswa Siswa
		err := rows.Scan(&siswa.SiswaID, &siswa.Username, &siswa.Password, &siswa.Nama, &siswa.Alamat, &siswa.RoleID)
		if err != nil {
			return nil, err
		}
		siswae = append(siswae, siswa)
	}
	return siswae, nil
	// return []User{}, nil // TODO: replace this
}
