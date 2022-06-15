package models

import "time"

type Saldo struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Saldo        int       `json:"saldo"`
	KategoriID   uint      `json:"category_id"`
	DibuatPada   time.Time `json:"dibuat_pada"`
	DiupdatePada time.Time `json:"diupdate_pada"`
	Kategory     Kategory  `json:"kategory" gorm:"foreignKey:KategoriID;constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
}
