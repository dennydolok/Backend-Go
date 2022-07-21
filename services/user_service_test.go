package services

import (
	"WallE/config"
	m "WallE/domains/mocks"
	"WallE/models"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository m.UserDomain

var code = "123456"
var falseCode = "6542321"

var user = models.User{
	ID:           1,
	Nama:         "John Doe",
	Email:        "dennydolok12@gmail.com",
	Password:     "MTIz",
	NomorHP:      "083890913399",
	Kode:         code,
	Verifikasi:   false,
	DiBuatPada:   time.Now(),
	DiUpdatePada: time.Now(),
	RoleID:       1,
}
var userVerfified = models.User{
	ID:           1,
	Nama:         "John Doe",
	Email:        "dennydolok12@gmail.com",
	Password:     "MTIz",
	NomorHP:      "083890913399",
	Kode:         code,
	Verifikasi:   true,
	DiBuatPada:   time.Now(),
	DiUpdatePada: time.Now(),
	RoleID:       1,
}
var userReset = models.ResetPassword{
	ID:           1,
	Kode:         code,
	Email:        "dennydolok12@gmail.com",
	UserID:       1,
	Selesai:      false,
	DiBuatPada:   time.Now(),
	DiUpdatePada: time.Now(),
}

func TestRegister(t *testing.T) {
	userService := serviceUser{
		repo:   &userRepository,
		config: config.Config{},
	}
	t.Run("success register", func(t *testing.T) {
		defer func() { SendMail = old }()
		SendMail = func(code, email, name, context string) error {
			return nil
		}
		userRepository.On("GetByEmail", mock.Anything).Return(user, errors.New("error")).Once()
		userRepository.On("Register", mock.Anything).Return(nil).Once()
		err := userService.Register(user)
		assert.NoError(t, err)
	})
	t.Run("success register failed to send email", func(t *testing.T) {
		defer func() { SendMail = old }()
		SendMail = func(code, email, name, context string) error {
			return errors.New("service error")
		}
		userRepository.On("GetByEmail", mock.Anything).Return(user, errors.New("error")).Once()
		userRepository.On("Register", mock.Anything).Return(nil).Once()
		err := userService.Register(user)
		assert.Error(t, err)
	})
	t.Run("error register email already register and verified", func(t *testing.T) {
		defer func() { SendMail = old }()
		SendMail = func(code, email, name, context string) error {
			return errors.New("error test")
		}
		userRepository.On("GetByEmail", mock.Anything).Return(userVerfified, nil).Once()
		userRepository.On("Register", userVerfified).Return(errors.New("error alredy register")).Once()
		err := userService.Register(userVerfified)
		assert.Error(t, err)
	})
	t.Run("error register email already register but not verified", func(t *testing.T) {
		defer func() { SendMail = old }()
		SendMail = func(code, email, name, context string) error {
			return nil
		}
		userRepository.On("GetByEmail", mock.Anything).Return(user, nil).Once()
		userRepository.On("Register", user).Return(errors.New("error alredy register")).Once()
		err := userService.Register(user)
		assert.Error(t, err)
	})
	t.Run("error resend verification email", func(t *testing.T) {
		defer func() { SendMail = old }()
		SendMail = func(code, email, name, context string) error {
			return errors.New("failed to resend")
		}
		userRepository.On("GetByEmail", mock.Anything).Return(user, nil).Once()
		userRepository.On("Register", user).Return(errors.New("error alredy register")).Once()
		err := userService.Register(user)
		assert.Error(t, err)
	})
}

func TestVerifikasi(t *testing.T) {
	userService := serviceUser{
		repo:   &userRepository,
		config: config.Config{},
	}
	t.Run("error user id", func(t *testing.T) {
		userRepository.On("GetByEmail", user.Email).Return(user, nil).Once()
		userRepository.On("Verifikasi", user.ID).Return(errors.New("error")).Once()
		_, err := userService.VerifikasiRegister(user.Email, code)
		assert.Error(t, err)
	})
	t.Run("success verification", func(t *testing.T) {
		userRepository.On("GetByEmail", user.Email).Return(user, nil).Once()
		userRepository.On("Verifikasi", user.ID).Return(nil).Once()
		_, err := userService.VerifikasiRegister(user.Email, code)
		assert.NoError(t, err)
	})
	t.Run("error no email", func(t *testing.T) {
		userRepository.On("GetByEmail", user.Email).Return(user, errors.New("email not found")).Once()
		_, err := userService.VerifikasiRegister(user.Email, code)
		assert.Error(t, err)
	})
	t.Run("error wrong code", func(t *testing.T) {
		userRepository.On("GetByEmail", user.Email).Return(user, nil).Once()
		userRepository.On("Verifikasi", user.ID).Return(nil).Once()
		_, err := userService.VerifikasiRegister(user.Email, falseCode)
		assert.Error(t, err)
	})
}

func TestGetUserDataById(t *testing.T) {
	userService := serviceUser{
		repo:   &userRepository,
		config: config.Config{},
	}
	t.Run("success get data", func(t *testing.T) {
		userRepository.On("GetUserDataById", mock.Anything).Return(userVerfified, nil).Once()
		user, err := userService.GetUserDataById(uint(1))
		assert.NoError(t, err)
		assert.NotEmpty(t, user)
	})
	t.Run("error get data", func(t *testing.T) {
		emptyUser := models.User{}
		userRepository.On("GetUserDataById", mock.Anything).Return(emptyUser, errors.New("error")).Once()
		user, err := userService.GetUserDataById(uint(1))
		assert.Error(t, err)
		assert.Empty(t, user)
	})
}

func TestLogin(t *testing.T) {
	userService := serviceUser{
		repo:   &userRepository,
		config: config.Config{},
	}
	t.Run("succes login", func(t *testing.T) {
		userRepository.On("GetByEmail", userVerfified.Email).Return(userVerfified, nil).Once()
		token, code := userService.Login(userVerfified.Email, "123")
		assert.Equal(t, 202, code)
		assert.NotEmpty(t, token)
	})
	t.Run("failed login not verified", func(t *testing.T) {
		userRepository.On("GetByEmail", user.Email).Return(user, nil).Once()
		token, code := userService.Login(user.Email, "123")
		assert.Equal(t, 406, code)
		assert.NotEmpty(t, token)
	})
	t.Run("failed login not registered", func(t *testing.T) {
		userRepository.On("GetByEmail", user.Email).Return(user, errors.New("not registered")).Once()
		token, code := userService.Login(user.Email, "123")
		assert.Equal(t, 404, code)
		assert.NotEmpty(t, token)
	})
	t.Run("failed login wrong password", func(t *testing.T) {
		userRepository.On("GetByEmail", user.Email).Return(user, nil).Once()
		token, code := userService.Login(user.Email, "1234")
		assert.Equal(t, 401, code)
		assert.NotEmpty(t, token)
	})
}

func TestUpdatePassword(t *testing.T) {
	userService := serviceUser{
		repo:   &userRepository,
		config: config.Config{},
	}
	t.Run("failed update password", func(t *testing.T) {
		userRepository.On("GetResetPassword", userReset.Email).Return(userReset, nil).Once()
		userRepository.On("UpdatePassword", mock.Anything, "MzIx").Return(errors.New("error update")).Once()
		err := userService.UpdatePassword(userReset.Email, "321", code)
		assert.Error(t, err)
	})
	t.Run("failed update reset table", func(t *testing.T) {
		userRepository.On("GetResetPassword", mock.Anything).Return(userReset, nil).Once()
		userRepository.On("UpdatePassword", mock.Anything, mock.Anything).Return(nil).Once()
		userRepository.On("UpdateResetTable", mock.Anything).Return(errors.New("error table")).Once()
		err := userService.UpdatePassword(userReset.Email, "123", code)
		assert.Error(t, err)
	})
	t.Run("success update", func(t *testing.T) {
		userRepository.On("GetResetPassword", mock.Anything).Return(userReset, nil).Once()
		userRepository.On("UpdatePassword", mock.Anything, mock.Anything).Return(nil).Once()
		userRepository.On("UpdateResetTable", mock.Anything).Return(nil).Once()
		err := userService.UpdatePassword(userReset.Email, "123", code)
		assert.NoError(t, err)
	})
	t.Run("failed update no email", func(t *testing.T) {
		userRepository.On("GetResetPassword", mock.Anything).Return(userReset, errors.New("email not found")).Once()
		userRepository.On("UpdatePassword", mock.Anything, mock.Anything).Return(nil).Once()
		userRepository.On("UpdateResetTable", mock.Anything).Return(nil).Once()
		err := userService.UpdatePassword(userReset.Email, "123", code)
		assert.Error(t, err)
	})
	t.Run("failed update wrong code", func(t *testing.T) {
		userRepository.On("GetResetPassword", mock.Anything).Return(userReset, nil).Once()
		userRepository.On("UpdatePassword", mock.Anything, mock.Anything).Return(nil).Once()
		userRepository.On("UpdateResetTable", mock.Anything).Return(nil).Once()
		err := userService.UpdatePassword(userReset.Email, "123", falseCode)
		assert.Error(t, err)
	})
}

func TestUpdateUser(t *testing.T) {
	userService := serviceUser{
		repo:   &userRepository,
		config: config.Config{},
	}
	t.Run("update user", func(t *testing.T) {
		userRepository.On("UpdateUserData", user.ID, user).Return(nil).Once()
		err := userService.UpdateUserData(user.ID, user)
		assert.NoError(t, err)
	})
}

func TestCreateResetPassword(t *testing.T) {
	userService := serviceUser{
		repo:   &userRepository,
		config: config.Config{},
	}
	t.Run("error get user", func(t *testing.T) {
		userRepository.On("GetByEmail", mock.Anything).Return(user, errors.New("error get user")).Once()
		err := userService.CreateResetPassword(user.Email)
		assert.Error(t, err)
	})
	t.Run("error create reset", func(t *testing.T) {
		userRepository.On("GetByEmail", mock.Anything).Return(user, nil).Once()
		userRepository.On("CreateResetPassword", mock.Anything).Return(errors.New("error create reset")).Once()
		err := userService.CreateResetPassword(user.Email)
		assert.Error(t, err)
	})
	t.Run("error send email reset", func(t *testing.T) {
		defer func() { SendMail = old }()
		SendMail = func(code, email, name, context string) error {
			return errors.New("service error")
		}
		userRepository.On("GetByEmail", mock.Anything).Return(user, nil).Once()
		userRepository.On("CreateResetPassword", mock.Anything).Return(nil).Once()
		err := userService.CreateResetPassword(user.Email)
		assert.Error(t, err)
	})
	t.Run("success", func(t *testing.T) {
		defer func() { SendMail = old }()
		SendMail = func(code, email, name, context string) error {
			return nil
		}
		userRepository.On("GetByEmail", mock.Anything).Return(user, nil).Once()
		userRepository.On("CreateResetPassword", mock.Anything).Return(nil).Once()
		err := userService.CreateResetPassword(user.Email)
		assert.NoError(t, err)
	})
}

func TestNewUserService(t *testing.T) {
	result := NewUserService(&userRepository, config.Config{})
	assert.NotEmpty(t, result)
}

func TestGenerateCode(t *testing.T) {
	code = GenerateCode()
	assert.NotEmpty(t, code)
}
