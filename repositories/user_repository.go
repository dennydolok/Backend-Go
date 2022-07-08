package repositories

import (
	"WallE/domains"
	"WallE/models"
	"errors"
	"time"

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

func (r *repositoryUser) UpdateUserData(id uint, user models.User) error {
	err := r.DB.Model(&user).Where("id = ?", id).Omit("dibuat_pada").Updates(models.User{
		Nama:         user.Nama,
		DiUpdatePada: time.Now(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryUser) GetByEmail(email string) (models.User, error) {
	user := models.User{}
	data := r.DB.Where("email = ?", email).Find(&user)
	if data.RowsAffected < 1 {
		return user, errors.New("Email tidak terdaftar")
	}
	return user, nil
}

func (r *repositoryUser) GetUserDataById(id uint) (models.User, error) {
	user := models.User{}
	data := r.DB.Where("id = ?", id).Find(&user).Error
	if data != nil {
		return user, errors.New("User tidak ditemukan")
	}
	return user, nil
}

func (r *repositoryUser) Verifikasi(id uint) error {
	user := models.User{}
	err := r.DB.Model(&user).Where("id = ?", id).Update("verifikasi", true).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryUser) CreateResetPassword(reset models.ResetPassword) error {
	err := r.DB.Create(&reset).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryUser) GetResetPassword(email string) (models.ResetPassword, error) {
	reset := models.ResetPassword{}
	data := r.DB.Where("email = ?", email).Where("selesai = ?", false).Find(&reset)
	if data.RowsAffected < 1 {
		return reset, errors.New("Email tidak ditemukan")
	}
	return reset, nil
}

func (r *repositoryUser) GetUserByEmail(email string) (models.User, error) {
	user := models.User{}
	data := r.DB.Model(&user).Where("email = ?", email)
	if data.RowsAffected < 1 {
		return user, errors.New("Email tidak ditemukan")
	}
	return user, nil
}

func (r *repositoryUser) UpdatePassword(email, password string) error {
	user := models.User{}
	err := r.DB.Model(&user).Where("email = ?", email).Update("password", password).Update("di_update_pada", time.Now()).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryUser) UpdateResetTable(email string) error {
	reset := models.ResetPassword{}
	err := r.DB.Model(&reset).Where("email = ?", email).Where("selesai = ?", false).Update("selesai", true).Error
	if err != nil {
		return errors.New("database error")
	}
	return nil
}

func NewUserRepository(db *gorm.DB) domains.UserDomain {
	return &repositoryUser{
		DB: db,
	}
}
