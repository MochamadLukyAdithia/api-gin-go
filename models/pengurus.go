package models

type Pengurus struct {
	IdPengurus int64  `json:"id_pengurus" gorm:"primaryKey;autoIncrement"`
	Nama       string `json:"nama" binding:"required"`
	Gambar     string `json:"gambar" binding:"required"`
	Divisi     string `json:"divisi" binding:"required"`
}
