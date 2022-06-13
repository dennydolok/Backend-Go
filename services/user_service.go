package services

import (
	"WallE/config"
	"WallE/domains"
	"WallE/helper"
	"WallE/models"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
)

type serviceUser struct {
	repo   domains.UserDomain
	config config.Config
}

func (s *serviceUser) Register(user models.User) error {
	_, check := s.repo.GetByEmail(user.Email)
	if check == nil {
		return errors.New("Email sudah terdaftar")
	}
	user.Password = base64.StdEncoding.EncodeToString([]byte(user.Password))
	user.Code = GenerateCode()
	err := helper.SendMail(user.Code, user.Email, user.Name, "Registrasi")
	if err != nil {
		fmt.Println(err)
		return errors.New("Gagal kirim email verifikasi")
	}
	return s.repo.Register(user)
}

func (s *serviceUser) VerifikasiRegister(email, kode string) error {
	fmt.Println(email)
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return err
	}
	fmt.Println(user)
	fmt.Println(kode)
	if kode != user.Code {
		return errors.New("Kode Salah")
	}
	err = s.repo.Verifikasi(user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceUser) Login(email, password string) (string, int) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "Email tidak terdaftar", http.StatusNotFound
	}
	if base64.StdEncoding.EncodeToString([]byte(password)) != user.Password {
		return "Password Salah", http.StatusUnauthorized
	}
	token, err := helper.CreateToken(user.ID, user.RoleID, s.config.SECRET_KEY)
	return token, http.StatusAccepted
}

func (s *serviceUser) CreateResetPassword(email string) error {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return err
	}
	reset := models.ResetPassword{}
	reset.Email = user.Email
	reset.UserID = user.ID
	reset.Code = GenerateCode()
	err = s.repo.CreateResetPassword(reset)
	if err != nil {
		fmt.Println(err)
		return errors.New("Kesalahan database")
	}
	err = helper.SendMail(reset.Code, email, user.Name, "Hilang Password")
	if err != nil {
		fmt.Println(err)
		return errors.New("Gagal")
	}
	return nil
}

func (s *serviceUser) UpdatePassword(email, password, code string) error {
	user, err := s.repo.GetResetPassword(email)
	if err != nil {
		return err
	}
	if user.Code != code {
		return errors.New("Kode Salah")
	}
	err = s.repo.UpdatePassword(email, base64.StdEncoding.EncodeToString([]byte(password)))
	if err != nil {
		return err
	}
	return nil
}

func NewUserService(repo domains.UserDomain, conf config.Config) domains.UserService {
	return &serviceUser{
		repo:   repo,
		config: conf,
	}
}

func GenerateCode() string {
	var letters = []rune("1234567890")
	b := make([]rune, 5)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
