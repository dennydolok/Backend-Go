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

	"github.com/mailjet/mailjet-apiv3-go/v3"
)

type serviceUser struct {
	repo   domains.UserDomain
	config config.Config
}

func (s *serviceUser) Register(user models.User) error {
	user.Password = base64.StdEncoding.EncodeToString([]byte(user.Password))
	user.Code = GenerateCode()
	err := SendMail(user.Code, user.Email, user.Name, "registrasi")
	if err != nil {
		fmt.Println(err)
		return errors.New("Gagal kirim email verifikasi")
	}
	return s.repo.Register(user)
}

func (s *serviceUser) VerifikasiRegister(email, kode string) error {
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

func SendMail(code, email, name, context string) error {
	publicKey := "b5cd4a33c4ea6788fbdc347067b3a35b"
	secretKey := "4fbc6b06490243458fe04b3182f3f818"
	mj := mailjet.NewMailjetClient(publicKey, secretKey)
	messageInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "bearuang0816@gmail.com",
				Name:  "WallE",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: email,
					Name:  name,
				},
			},
			Subject:  "Verifikasi Kode",
			TextPart: "Berikut verifikasi kode anda untuk " + context + "! \n " + code,
			HTMLPart: "<h3>Berikut verifikasi kode anda untuk " + context + "!</h3> <br /><center><strong>" + code + "</strong></center>",
		},
	}
	messages := mailjet.MessagesV31{Info: messageInfo}
	_, err := mj.SendMailV31(&messages)
	if err != nil {
		return err
	}
	return nil
}
