package services

import (
	"WallE/config"
	"WallE/domains"
	"WallE/helper"
	"WallE/models"
	"encoding/base64"
	"errors"
	"math/rand"
	"net/http"
	"time"
)

type serviceUser struct {
	repo   domains.UserDomain
	config config.Config
}

func (s *serviceUser) Register(user models.User) error {
	userExist, check := s.repo.GetByEmail(user.Email)
	user.DiBuatPada = time.Now()
	user.DiUpdatePada = time.Now()
	user.Kode = GenerateCode()
	user.Password = base64.StdEncoding.EncodeToString([]byte(user.Password))
	if check == nil {
		if userExist.Verifikasi == false {
			err := helper.SendMail(userExist.Kode, userExist.Email, userExist.Nama, "Registrasi")
			if err != nil {
				return errors.New("Sistem Error")
			}
			return errors.New("resend")
		}
		return errors.New("Email sudah terdaftar")
	}
	err := helper.SendMail(user.Kode, user.Email, user.Nama, "Registrasi")
	if err != nil {
		return errors.New("Gagal kirim email verifikasi")
	}
	return s.repo.Register(user)
}

func (s *serviceUser) VerifikasiRegister(email, kode string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if kode != user.Kode {
		return "", errors.New("Kode Salah")
	}
	err = s.repo.Verifikasi(user.ID)
	if err != nil {
		return "", err
	}
	token, err := helper.CreateToken(user.ID, user.RoleID, s.config.SECRET_KEY)
	return token, nil
}

func (s *serviceUser) GetUserDataById(id uint) (models.User, error) {
	user, err := s.repo.GetUserDataById(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *serviceUser) Login(email, password string) (string, int) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "Email tidak terdaftar", http.StatusNotFound
	}
	if base64.StdEncoding.EncodeToString([]byte(password)) != user.Password {
		return "Password Salah", http.StatusUnauthorized
	}
	if user.Verifikasi != true {
		return "belum verifikasi", http.StatusNotAcceptable
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
	reset.Kode = GenerateCode()
	reset.DiBuatPada = time.Now()
	reset.DiUpdatePada = time.Now()
	err = s.repo.CreateResetPassword(reset)
	if err != nil {
		return errors.New("Kesalahan database")
	}
	err = helper.SendMail(reset.Kode, email, user.Nama, "Hilang Password")
	if err != nil {
		return errors.New("Gagal")
	}
	return nil
}

func (s *serviceUser) UpdatePassword(email, password, code string) error {
	user, err := s.repo.GetResetPassword(email)
	if err != nil {
		return err
	}
	if user.Kode != code {
		return errors.New("Kode Salah")
	}
	err = s.repo.UpdatePassword(email, base64.StdEncoding.EncodeToString([]byte(password)))
	if err != nil {
		return err
	}
	err = s.repo.UpdateResetTable(email)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceUser) UpdateUserData(id uint, user models.User) error {
	return s.repo.UpdateUserData(id, user)
}

func NewUserService(repo domains.UserDomain, conf config.Config) domains.UserService {
	return &serviceUser{
		repo:   repo,
		config: conf,
	}
}

func GenerateCode() string {
	var letters = []rune("1234567890")
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
