package models

import "time"

type Provider struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      int       `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Category  Category  `json:"category" gorm:"foreignKey:BalanceID;constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
}
