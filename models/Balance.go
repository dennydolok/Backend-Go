package models

import "time"

type Balance struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Balance   int       `json:"balance"`
	CategoryID uint 	`json:"categoryid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
}
