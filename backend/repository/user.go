package repository

import "database/sql"

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) LoginUser(username string, password string) (*User, error) {
	sqlStatement := `SELECT * FROM user WHERE username = ? AND password = ?;`

	rows, err := u.db.Query(sqlStatement, username, password)
	if err != nil {
		return nil, err
	}

	var user User
	for rows.Next() {
		err = rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Nama, &user.Alamat, &user.NoHp, &user.Role)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}
