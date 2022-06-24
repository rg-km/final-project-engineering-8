package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../final_project.db")
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
	CREATE TABLE IF NOT EXISTS user (
		UserID integer not null primary key AUTOINCREMENT,
		username varchar(255) not null,
		password varchar(255) not null,
		nama varchar(255) not null,
		alamat varchar(255) not null,
		noHp varchar(255) not null,
		role varchar(255) not null,
		profilePict TEXT NOT NULL DEFAULT "https://bucket-halloguru.online/bucket-image/halloguru/default.png")
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

	CREATE TABLE IF NOT EXISTS info_guru (
		InfoGuruID integer not null primary key AUTOINCREMENT,
		deskripsi varchar(255) not null,
		biaya varchar(255) not null,
		ratting varchar(255) not null,
		UserID integer not null,
		JenjangID integer not null,
		PelajaranID integer not null,
		KategoriID integer not null,
		FOREIGN KEY (UserID) REFERENCES user(UserID),
		FOREIGN KEY (JenjangID) REFERENCES jenjang(JenjangID),
		FOREIGN KEY (PelajaranID) REFERENCES pelajaran(PelajaranID),
		FOREIGN KEY (KategoriID) REFERENCES kategori(KategoriID)
	);

	INSERT INTO jenjang (jenjang) VALUES ("SD"), ("SMP"), ("SMA"), ("Semua Kategori");
	INSERT INTO pelajaran (pelajaran) VALUES ("Matematika"),("Bahasa Inggris"),("Kimia"),("Biologi"),("Fisika"),("Semua Mata Pelajaran");
	INSERT INTO kategori (kategori) VALUES ("Anak Berkebutuhan Khusus"), ("Anak Non Berkebutuhan Khusus"),("Semua Kategori");

	INSERT INTO user (username, password, nama, alamat, noHP,role) 
	VALUES 
	("admin", "admin", "admin", "admin", "admin", "admin"),
	("ucup", "ucup123", "Muhammad Ucup", "Jl Bojong Soang No.1, Kota Bandung", "0812341234", "siswa"),
	("Rendi", "Rendi123", "Muhammad Rendi", "Jl Rajawali No.1, Jakarta Timur", "0812341234", "siswa"),
	("Subabjo", "Subabjo123", "Muhammad Subabjo", "Jl Dago, Kota Bandung", "0812341234", "guru"),
	("Kurniawan", "Kurniawan123", "Muhammad Kurniawan", "Jl Jendral Sudirman, Jakarta Timur", "0812341234", "guru");

	INSERT INTO info_guru (deskripsi, biaya, ratting, UserID, JenjangID, PelajaranID, KategoriID)
	VALUES
	("Saya memiliki keahlian dalam bidang matematika", "Rp. 100.000", "4.5", 3, 1, 1, 1),
	("Saya memiliki keahlian dalam bidang bahasa inggris", "Rp. 100.000", "4.5", 4, 1, 2, 2);
	`)
	if err != nil {
		return "Tabel gagal dibuat", err
	}
	return "Tabel berhasil dibuat", nil
}
