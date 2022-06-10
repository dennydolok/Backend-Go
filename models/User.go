package models

import "time"

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phonenumber"`
	Code        string    `json:"code"`
	Verified    bool      `json:"verified"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt   time.Time `json:"updatedat"`
	RoleID      uint      `json:"roleid"`
	Role        Role      `json:"userrole" gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
}
