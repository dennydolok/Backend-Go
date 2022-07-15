package models

type Provider struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Nama string `json:"nama"`
}
