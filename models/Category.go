package models

import "time"

type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      string    `json:"category"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
