package api

import "github.com/golang-jwt/jwt/v4"

//inputan user
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Nama        string `json:"nama"`
	Alamat      string `json:"alamat"`
	NoHp        string `json:"no_hp"`
	Role        string `json:"role,omitempty"`
	Deskripsi   string `json:"deskripsi,omitempty"`
	Biaya       string `json:"biaya,omitempty"`
	Ratting     string `json:"ratting,omitempty"`
	JenjangID   int    `json:"jenjang_id,omitempty"`
	PelajaranID int    `json:"pelajaran_id,omitempty"`
	KategoriID  int    `json:"kategori_id,omitempty"`
}

//jwt token
var jwtKey = []byte("key")

type Claims struct {
	ID       int64 `json:"id"`
	Username string
	Role     string
	jwt.StandardClaims
}

type Users struct {
	UserID      int    `json:"user_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Nama        string `json:"nama"`
	Alamat      string `json:"alamat"`
	NoHp        string `json:"no_hp"`
	ProfilePict string `json:"profile_pict"`
	Role        string `json:"role"`
}
