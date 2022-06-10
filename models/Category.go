package models

import "time"

type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
