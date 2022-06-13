package domains

import "WallE/models"

type UserDomain interface {
	Register(user models.User) error
	GetByEmail(email string) (models.User, error)
	Verifikasi(id uint) error
	CreateResetPassword(reset models.ResetPassword) error
	GetResetPassword(email string) (models.ResetPassword, error)
	UpdatePassword(email, password string) error
}

type UserService interface {
	Register(user models.User) error
	VerifikasiRegister(email, kode string) error
	Login(email, password string) (string, int)
	CreateResetPassword(email string) error
	UpdatePassword(email, password, code string) error
}
