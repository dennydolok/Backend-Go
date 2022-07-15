package models

type Kategori struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Nama string `json:"nama"`
}
