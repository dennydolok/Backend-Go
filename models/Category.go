package models

type Category struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      string `json:"category"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
