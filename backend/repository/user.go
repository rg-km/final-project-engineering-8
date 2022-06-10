package repository

import (
	"database/sql"
	"errors"
)

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

func (u *UserRepository) Register(username string, password string, nama string, alamat string, noHp string, role string) (*User, error) {
	check, err := u.CheckAccount(username, password)

	//check jika data sudah ada
	if check.UserID != 0 {
		err1 := errors.New("Akun sudah ada")
		return nil, err1
	} else {
		//jika data belum ada
		sqlStatement := `INSERT INTO user (username, password, nama, alamat, noHp, role) VALUES (?, ?, ?, ?, ?, ?);`

		_, err = u.db.Exec(sqlStatement, username, password, nama, alamat, noHp, role)
		if err != nil {
			return nil, err
		}

		return &User{Username: username, Password: password, Nama: nama, Alamat: alamat, NoHp: noHp, Role: role}, nil
	}
}

func (u *UserRepository) CheckAccount(username string, password string) (*User, error) {
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

// func (u *UserRepository) FetchAllUsers() ([]User, error) {
// 	sqlStmt := `SELECT id, username, password, role, loggedin FROM users`
// 	rows, err := u.db.Query(sqlStmt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	users := []User{}
// 	for rows.Next() {
// 		var user User
// 		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin)
// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}
// 	return users, nil
// 	// return []User{}, nil // TODO: replace this
// }
