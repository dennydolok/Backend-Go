package models

import "time"

type ResetPassword struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Kode         string    `json:"kode"`
	Email        string    `json:"email"`
	UserID       uint      `json:"userid"`
	Selesai      bool      `json:"selesai"`
	DiBuatPada   time.Time `json:"dibuat_pada"`
	DiUpdatePada time.Time `json:"diupdate_pada"`
	User         User      `json:"resetuser" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
}
