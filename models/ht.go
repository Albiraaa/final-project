package models

type HT struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	Serial     string   `json:"serial"`
	Merk       string   `json:"merk"`
	Status     string   `json:"status"`
	KategoriID uint     `json:"kategori_id"`
	LokasiID   uint     `json:"lokasi_id"`
	Kategori   Kategori `gorm:"foreignKey:KategoriID"`
	Lokasi     Lokasi   `gorm:"foreignKey:LokasiID"`
}
