package repositories

import (
	"WallE/domains"
	"WallE/models"
	"errors"

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

func (r *repositoryUser) GetByEmail(email string) (models.User, error) {
	user := models.User{}
	data := r.DB.Find(&user).Where("email =?", email)
	if data.RowsAffected < 1 {
		return user, errors.New("Email tidak terdaftar")
	}
	return user, nil
}

func (r *repositoryUser) Verifikasi(id uint) error{
	user := models.User{}
	err := r.DB.Find(&user).Where("id = ?", id).Update("Verified", true).Error
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
