package models

import "time"

type Kategory struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Nama         string    `json:"nama"`
	DibuatPada   time.Time `json:"dibuat_pada"`
	DiupdatePada time.Time `json:"diupdate_pada"`
}
