package repository

import "database/sql"

type AdminRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (u *AdminRepository) LoginAdmin(username string, password string) (*Admin, error) {
	sqlStatement := `SELECT * FROM admin WHERE username = ? AND password = ?;`

	rows, err := u.db.Query(sqlStatement, username, password)
	if err != nil {
		return nil, err
	}

	var admin Admin
	for rows.Next() {
		err = rows.Scan(&admin.AdminID, &admin.Username, &admin.Password, &admin.Nama, &admin.RoleID)
		if err != nil {
			return nil, err
		}
	}

	return &admin, nil
}
