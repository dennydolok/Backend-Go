package domains

import "WallE/models"

type UserDomain interface {
	Register(user models.User) error
}

type UserService interface {
	Register(user models.User) error
}
