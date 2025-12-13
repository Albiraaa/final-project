package models

type Lokasi struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Nama string `json:"nama"`
}
