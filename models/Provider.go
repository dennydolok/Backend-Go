package models

import "time"

type Provider struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Nama         int       `json:"nama"`
	DibuatPada   time.Time `json:"dibuat_pada"`
	DiupdatePada time.Time `json:"diupdate_pada"`
}
