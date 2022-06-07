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
