package domains

import "WallE/models"

type UserDomain interface {
	Register(user models.User) error
	GetByEmail(email string) (models.User, error)
	Verifikasi(id uint) error
}

type UserService interface {
	Register(user models.User) error
	VerifikasiRegister(email, kode string) error
}
