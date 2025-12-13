package models

type Kategori struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Nama string `json:"nama"`
}
