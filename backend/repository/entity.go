package repository

type Siswa struct {
	SiswaID  int    `json:"siswa_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nama     string `json:"nama"`
	Alamat   string `json:"alamat"`
	RoleID   int    `json:"role_id"`
}

type Admin struct {
	AdminID  int    `json:"admin_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nama     string `json:"nama"`
	RoleID   int    `json:"role_id"`
}

type Guru struct {
	GuruID      int    `json:"guru_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Nama        string `json:"nama"`
	Alamat      string `json:"alamat"`
	NoHp        string `json:"no_hp"`
	Deskripsi   string `json:"deskripsi"`
	Biaya       string `json:"biaya"`
	RoleID      int    `json:"role_id"`
	JenjangID   int    `json:"jenjang_id"`
	PelajaranID int    `json:"pelajaran_id"`
	KategoriID  int    `json:"kategori_id"`
}
