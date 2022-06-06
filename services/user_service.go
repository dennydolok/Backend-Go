package services

import (
	"WallE/domains"
	"WallE/models"
	"encoding/base64"
)

type serviceUser struct {
	repo domains.UserDomain
}

func (s *serviceUser) Register(user models.User) error {
	user.Password = base64.StdEncoding.EncodeToString([]byte(user.Password))
	return s.repo.Register(user)
}

func NewUserService(repo domains.UserDomain) domains.UserService {
	return &serviceUser{
		repo: repo,
	}
}
