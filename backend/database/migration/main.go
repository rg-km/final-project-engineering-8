package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	con "github.com/rg-km/final-project-engineering-8/backend/database/connection"
)

func main() {
	db, err := con.ConnectSQLite()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	asd, err := CreateTableTest(db)
	if err != nil {
		panic(err)
	}
	println(asd)
}

func CreateTableTest(db *sql.DB) (string, error) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS siswa (
		SiswaID integer not null primary key AUTOINCREMENT,
		username varchar(255) not null,
		password varchar(255) not null,
		nama varchar(255) not null,
		alamat varchar(255) not null,
		RoleID varchar(255) not null,
		FOREIGN KEY (RoleID) REFERENCES role(RoleID)
	);

	CREATE TABLE IF NOT EXISTS role (
		RoleID integer not null primary key AUTOINCREMENT,
		role varchar(255) not null
	);

	CREATE TABLE IF NOT EXISTS admin (
		AdminID integer not null primary key AUTOINCREMENT,
		username varchar(255) not null,
		password varchar(255) not null,
		nama varchar(255) not null,
		RoleID varchar(255) not null,
		FOREIGN KEY (RoleID) REFERENCES role(RoleID)
	);

	CREATE TABLE IF NOT EXISTS jenjang (
		JenjangID integer not null primary key AUTOINCREMENT,
		jenjang varchar(255) not null
	);

	CREATE TABLE IF NOT EXISTS pelajaran (
		PelajaranID integer not null primary key AUTOINCREMENT,
		pelajaran varchar(255) not null
	);

	CREATE TABLE IF NOT EXISTS kategori (
		KategoriID integer not null primary key AUTOINCREMENT,
		kategori varchar(255) not null
	);

	CREATE TABLE IF NOT EXISTS Guru (
		GuruID integer not null primary key AUTOINCREMENT,
		username varchar(255) not null,
		password varchar(255) not null,
		nama varchar(255) not null,
		alamat varchar(255) not null,
		noHp varchar(255) not null,
		deskripsi varchar(255) not null,
		biaya varchar(255) not null,
		RoleID varchar(255) not null,
		JenjangID varchar(255) not null,
		PelajaranID varchar(255) not null,
		KategoriID varchar(255) not null,
		FOREIGN KEY (RoleID) REFERENCES role(RoleID),
		FOREIGN KEY (JenjangID) REFERENCES jenjang(JenjangID),
		FOREIGN KEY (PelajaranID) REFERENCES pelajaran(PelajaranID),
		FOREIGN KEY (KategoriID) REFERENCES kategori(KategoriID)
	);

	CREATE TABLE IF NOT EXISTS siswa_guru (
		SiswaID integer not null,
		GuruID integer not null,
		FOREIGN KEY (SiswaID) REFERENCES siswa(SiswaID),
		FOREIGN KEY (GuruID) REFERENCES guru(GuruID)
	);

	INSERT INTO role (role) VALUES ("admin"), ("siswa"), ("guru");
	INSERT INTO admin (username, password, nama, RoleID) VALUES ("admin", "admin", "admin", 1);
	INSERT INTO jenjang (jenjang) VALUES ("SD"), ("SMP"), ("SMA");
	INSERT INTO pelajaran (pelajaran) VALUES ("Matematika"),("Bahasa Inggris");
	INSERT INTO kategori (kategori) VALUES ("Anak Berkebutuhan Khusus"), ("Anak Non Berkebutuhan Khusus");

	INSERT INTO siswa (username, password, nama, alamat, RoleID) 
	VALUES 
	("ucup", "ucup123", "Muhammad Ucup", "Jl Bojong Soang No.1, Kota Bandung", 2),
	("Rendi", "Rendi123", "Muhammad Rendi", "Jl Rajawali No.1, Jakarta Timur", 2);

	INSERT INTO guru (username, password, nama, alamat, noHp, deskripsi, biaya, RoleID, JenjangID, PelajaranID, KategoriID)
	VALUES
	("Subabjo", "Subabjo123", "Muhammad Subabjo", "Jl Dago, Kota Bandung", "081212121212", "Saya guru yang mengajar Matematika", "10000", 2, 1, 1, 1),
	("Kurniawan", "Kurniawan123", "Muhammad Kurniawan", "Jl Jendral Sudirman, Jakarta Timur", "081212121212", "Saya guru yang mengajar Bahasa Inggris", "20000", 2, 1, 2, 2);
	
	INSERT INTO siswa_guru (SiswaID, GuruID)
	VALUES
	(1, 1),
	(2, 2);

	`)
	if err != nil {
		return "Tabel gagal dibuat", err
	}
	return "Tabel berhasil dibuat", nil
}
