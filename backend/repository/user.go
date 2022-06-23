package repository

import (
	"context"
	"database/sql"
	"errors"
	// repo "github.com/rg-km/final-project-engineering-8/backend/repository"
)

var DefaultProfilePict = "https://halloguru.herokuapp.com/bucket-image/halloguru/1655885462662157252.png"

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) LoginUser(username string) (*User, error) {
	sqlStatement := `SELECT * FROM user WHERE username = ?;`

	rows, err := u.db.Query(sqlStatement, username)
	if err != nil {
		return nil, err
	}

	var user User
	for rows.Next() {
		err = rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Nama, &user.Alamat, &user.NoHp, &user.Role, &user.ProfilePict)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (u *UserRepository) StudentRegister(username string, password string, nama string, alamat string, noHp string) (*User, error) {
	check, err := u.CheckAccount(username)

	//check jika data sudah ada
	if check.UserID != 0 {
		err1 := errors.New("Akun sudah ada")
		return nil, err1
	} else {
		//jika data belum ada
		sqlStatement := `INSERT INTO user (username, password, nama, alamat, noHp, role) VALUES (?, ?, ?, ?, ?, ?);`

		_, err = u.db.Exec(sqlStatement, username, password, nama, alamat, noHp, "siswa")
		if err != nil {
			return nil, err
		}

		return &User{Username: username, Password: password, Nama: nama, Alamat: alamat, NoHp: noHp, ProfilePict: DefaultProfilePict, Role: "siswa"}, nil
	}
}

func (u *UserRepository) TeacherRegister(username string, password string, nama string, alamat string, noHp string, deskripsi string, biaya string, jenjangID int, pelajaranID int, kategoriID int) (*TeacherRegister, error) {
	check, _ := u.CheckAccount(username)

	//check jika data sudah ada
	if check.UserID != 0 {
		err1 := errors.New("Akun sudah ada")
		return nil, err1
	} else {
		//jika data belum ada
		sqlStatement := `INSERT INTO user (username, password, nama, alamat, noHp, role) VALUES (?, ?, ?, ?, ?, ?);`

		rows, err := u.db.Exec(sqlStatement, username, password, nama, alamat, noHp, "guru")
		if err != nil {
			return nil, err
		}

		userID, _ := rows.LastInsertId()

		sqlStatement2 := `INSERT INTO info_guru (deskripsi, biaya, ratting, userID, jenjangID, pelajaranID, kategoriID) VALUES (?, ?, ?, ?, ?, ?, ?);`

		rows, err = u.db.Exec(sqlStatement2, deskripsi, biaya, "1", userID, jenjangID, pelajaranID, kategoriID)
		if err != nil {
			return nil, err
		}

		// return &User{Username: username, Password: password, Nama: nama, Alamat: alamat, NoHp: noHp, Role: "guru"}, nil
		return &TeacherRegister{ID: int(userID), Name: nama, Address: alamat, NoHp: noHp, Description: deskripsi, Rating: "1", Fee: biaya, TeachingLevel: jenjangID, TeachingSubject: pelajaranID, TeachingCategory: kategoriID}, nil
	}
}

func (u *UserRepository) CheckAccount(username string) (*User, error) {
	sqlStatement := `SELECT * FROM user WHERE username = ?;`

	rows, err := u.db.Query(sqlStatement, username)
	if err != nil {
		return nil, err
	}

	var user User
	for rows.Next() {
		err = rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Nama, &user.Alamat, &user.NoHp, &user.Role, &user.ProfilePict)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (u *UserRepository) CheckAccountUpdate(username string, userID int) (*User, error) {
	sqlStatement := `SELECT * FROM user WHERE username = ? AND NOT userID = ?;`

	rows, err := u.db.Query(sqlStatement, username, userID)
	if err != nil {
		return nil, err
	}

	var user User
	for rows.Next() {
		err = rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Nama, &user.Alamat, &user.NoHp, &user.Role, &user.ProfilePict)
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

func (u *UserRepository) FetchAllTeachers(limit int, offset int) ([]Teacher, error) {
	sqlStmt := `
	SELECT
		u.UserID AS id,
		u.nama AS name,
		u.alamat AS address,
		u.noHp AS no_hp,
		u.profilePict AS profile_pict,
		g.deskripsi AS description,
		g.biaya AS fee,
		g.ratting as rating,
		p.pelajaran AS teaching_subject,
		k.kategori AS teaching_category,
		j.jenjang AS teaching_level
	FROM
		info_guru AS g
	JOIN USER AS u ON (g.UserID = u.UserID)
	JOIN kategori AS k ON (g.KategoriID = k.KategoriID)
	JOIN pelajaran AS p ON (g.PelajaranID = p.PelajaranID)
	JOIN jenjang AS j ON (g.JenjangID = j.JenjangID)
	LIMIT ?
	OFFSET ?
	`

	rows, err := u.db.Query(sqlStmt, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teachers := []Teacher{}
	for rows.Next() {
		var teacher Teacher
		err := rows.Scan(
			&teacher.ID,
			&teacher.Name,
			&teacher.Address,
			&teacher.NoHp,
			&teacher.ProfilePict,
			&teacher.Description,
			&teacher.Fee,
			&teacher.Rating,
			&teacher.TeachingSubject,
			&teacher.TeachingCategory,
			&teacher.TeachingLevel,
		)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}

func (u *UserRepository) GetNumberofTeacherRow() (int, error) {
	sqlStmt := `SELECT COUNT(*) FROM info_guru`
	var total int
	err := u.db.QueryRow(sqlStmt).Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}

func (u *UserRepository) UpdateTeacher(id string, teacher map[string]interface{}) error {
	ctx := context.Background()
	tx, err := u.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	sqlStmt := `
	UPDATE
		user
	SET
		nama = ?,
		alamat = ?,
		noHp = ?,
		profilePict = ?
	WHERE
		user.UserID = ?;
	`
	_, err = tx.Exec(
		sqlStmt,
		teacher["name"],
		teacher["address"],
		teacher["no_hp"],
		teacher["profile_pict"],
		id,
	)

	if err != nil {
		return err
	}

	sqlStmt = `
	UPDATE
		info_guru
	SET
		deskripsi = ?,
		biaya = ?,
		PelajaranID = ?,
		JenjangID = ?,
		KategoriID = ?
	WHERE
		UserID = ?;
	`

	_, err = tx.Exec(
		sqlStmt,
		teacher["description"],
		teacher["fee"],
		teacher["teaching_subject"],
		teacher["teaching_level"],
		teacher["teaching_category"],
		id,
	)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (u *UserRepository) GetTeacherByID(id string) (Teacher, error) {
	var teacher Teacher

	sqlStatement := `
	SELECT
		u.UserID AS id,
		u.nama AS name,
		u.alamat AS address,
		u.noHp AS no_hp,
		u.profilePict AS profile_pict,
		g.deskripsi AS description,
		g.biaya AS fee,
		g.ratting AS rating,
		p.pelajaran AS teaching_subject,
		k.kategori AS teaching_category,
		j.jenjang AS teaching_level
	FROM
		info_guru AS g
	JOIN USER AS u ON (g.UserID = u.UserID)
	JOIN kategori AS k ON (g.KategoriID = k.KategoriID)
	JOIN pelajaran AS p ON (g.PelajaranID = p.PelajaranID)
	JOIN jenjang AS j ON (g.JenjangID = j.JenjangID)
	WHERE u.UserID = ?
	`
	err := u.db.QueryRow(sqlStatement, id).Scan(
		&teacher.ID,
		&teacher.Name,
		&teacher.Address,
		&teacher.NoHp,
		&teacher.ProfilePict,
		&teacher.Description,
		&teacher.Fee,
		&teacher.Rating,
		&teacher.TeachingSubject,
		&teacher.TeachingCategory,
		&teacher.TeachingLevel,
	)

	if err != nil {
		return teacher, err
	}

	return teacher, nil
}

func (u *UserRepository) DeleteTeacherByUserID(id string) (code int, err error) {

	ctx := context.Background()
	tx, fail := u.db.BeginTx(ctx, nil)
	if fail != nil {
		return 500, fail
	}

	defer tx.Rollback()

	sqlStatement := `DELETE FROM user WHERE user.UserID = ?`

	if _, fail := tx.Exec(sqlStatement, id); fail != nil {
		return 500, fail
	}

	sqlStatement = `DELETE FROM info_guru WHERE info_guru.UserID = ?`

	if _, fail := tx.Exec(sqlStatement, id); fail != nil {
		return 500, fail
	}

	tx.Commit()
	return 200, nil
}

func (u *UserRepository) GetStudentProfile(username string) (*User, error) {
	sqlStatement := `SELECT * FROM user WHERE username = ?;`

	rows, err := u.db.Query(sqlStatement, username)
	if err != nil {
		return nil, err
	}

	var user User
	for rows.Next() {
		err = rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Nama, &user.Alamat, &user.NoHp, &user.Role, &user.ProfilePict)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (u *UserRepository) UpdateStudentById(userId int, username string, password string, nama string, alamat string, noHp string, profil_pict string) (*User, error) {
	check, _ := u.CheckAccountUpdate(username, userId)

	//check jika data sudah ada
	if check.UserID != 0 {
		err1 := errors.New("Akun sudah ada")
		return nil, err1
	} else {

		sqlStatement := `UPDATE user SET username = ?, password = ?, nama = ?, alamat = ?, noHp = ?, profilePict = ? WHERE userId = ?;`

		_, err := u.db.Exec(sqlStatement, username, password, nama, alamat, noHp, profil_pict, userId)
		if err != nil {
			return nil, err
		}

		return u.GetStudentProfile(username)
	}

}

func (u *UserRepository) DeleteUserByID(id int) (code int, err error) {
	sqlStatement := `DELETE FROM user WHERE user.UserID = ?`
	if _, err := u.db.Exec(sqlStatement, id); err != nil {
		return 500, nil
	}
	return 200, nil
}
