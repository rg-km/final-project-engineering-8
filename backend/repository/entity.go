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

type InfoGuru struct {
	InfoGuruID  int    `json:"info_guru_id"`
	Deskripsi   string `json:"deskripsi"`
	Biaya       string `json:"biaya"`
	Ratting     string `json:"ratting"`
	UserID      int    `json:"user_id"`
	JenjangID   int    `json:"jenjang_id"`
	PelajaranID int    `json:"pelajaran_id"`
	KategoriID  int    `json:"kategori_id"`
}

type Teacher struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	NoHp             string `json:"no_hp"`
	Description      string `json:"description"`
	Fee              string `json:"fee"`
	Rating           string `json:"rating"`
	TeachingCategory string `json:"teaching_category"`
	TeachingLevel    string `json:"teaching_level"`
	TeachingSubject  string `json:"teaching_subject"`
}
