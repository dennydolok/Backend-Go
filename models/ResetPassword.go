package models

import "time"

type ResetPassword struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Code      int       `json:"code"`
	UserID    uint      `json:"userid"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	User      User      `json:"resetuser" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
}
