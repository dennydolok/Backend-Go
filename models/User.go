package models

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Nama         string    `json:"nama" form:"nama"`
	Email        string    `json:"email" form:"email"`
	Password     string    `json:"password" form:"password"`
	NomorHP      string    `json:"nomor_handphone" form:"nomor_handphone"`
	Kode         string    `json:"kode" form:"kode"`
	Verifikasi   bool      `json:"verifikasi"`
	DiBuatPada   time.Time `json:"dibuat_pada"`
	DiUpdatePada time.Time `json:"diupdate_pada"`
	RoleID       uint      `json:"role_id"`
	Role         Role      `json:"-" gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
}
