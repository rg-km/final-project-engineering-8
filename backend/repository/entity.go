package repository

type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nama     string `json:"nama"`
	Alamat   string `json:"alamat"`
	NoHp     string `json:"no_hp"`
	Role     string `json:"role"`
}
