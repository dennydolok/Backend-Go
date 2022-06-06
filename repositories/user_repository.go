package repositories

import (
	"WallE/domains"
	"WallE/models"

	"gorm.io/gorm"
)

type repositoryUser struct {
	DB *gorm.DB
}

func (r *repositoryUser) Register(user models.User) error {
	err := r.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *gorm.DB) domains.UserDomain {
	return &repositoryUser{
		DB: db,
	}
}
