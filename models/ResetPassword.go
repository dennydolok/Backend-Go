package models

import "time"

type ResetPassword struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Code      string    `json:"code"`
	Email     string    `json:"email"`
	UserID    uint      `json:"userid"`
	IsDone    bool      `json:"isdone"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	User      User      `json:"resetuser" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
}
