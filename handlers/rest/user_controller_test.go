package rest

import (
	m "WallE/domains/mocks"
	"WallE/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userService m.UserService

func TestRegister(t *testing.T) {

	userService.On("Register", mock.Anything).Return(errors.New("error")).Once()
	userService.On("Register", mock.Anything).Return(errors.New("resend")).Once()
	userService.On("Register", mock.Anything).Return(nil).Once()

	userController := UserController{
		services: &userService,
	}

	e := echo.New()
	t.Run("error register", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		userController.Register(eContext)
		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("resend verification code", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		userController.Register(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("resend verification code", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		userController.Register(eContext)
		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetUserData(t *testing.T) {
	user := models.User{
		ID:           uint(1),
		Nama:         "Denny",
		Email:        "test@gmail.com",
		Password:     "MTIz",
		NomorHP:      "12313141231",
		Kode:         "123456",
		Verifikasi:   true,
		DiBuatPada:   time.Now(),
		DiUpdatePada: time.Now(),
		RoleID:       uint(1),
	}
	userService.On("GetUserDataById", mock.AnythingOfType("uint")).Return(user, nil).Once()
	userService.On("GetUserDataById", mock.AnythingOfType("uint")).Return(user, errors.New("error")).Once()
	userController := UserController{
		services: &userService,
	}

	e := echo.New()
	t.Run("get user data", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc2ODY5NzUsImlkIjo0LCJyb2xlIjoxfQ.Le1-pu7wUwRAHC6qhx4ScwSmzIMd9AqzbTEHY4S8WhA"
		bearer := "Bearer " + token
		r := httptest.NewRequest("GET", "/", bytes.NewBuffer(nil))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		err := userController.GetUserData(eContext)
		assert.NoError(t, err)
	})
	t.Run("error user data", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc2ODY5NzUsImlkIjo0LCJyb2xlIjoxfQ.Le1-pu7wUwRAHC6qhx4ScwSmzIMd9AqzbTEHY4S8WhA"
		bearer := "Bearer " + token
		r := httptest.NewRequest("GET", "/", bytes.NewBuffer(nil))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		userController.GetUserData(eContext)
		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestVerification(t *testing.T) {

	userService.On("VerifikasiRegister", mock.Anything, mock.Anything).Return("test", errors.New("error")).Once()
	userService.On("VerifikasiRegister", mock.Anything, mock.Anything).Return("sukses", nil).Once()

	userController := UserController{
		services: &userService,
	}

	e := echo.New()
	t.Run("error verification", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		userController.Verification(eContext)
		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("ok verification", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		userController.Verification(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestLogin(t *testing.T) {
	userService.On("Login", mock.Anything, mock.Anything).Return("email not found", 404).Once()
	userService.On("Login", mock.Anything, mock.Anything).Return("failed login", 401).Once()
	userService.On("Login", mock.Anything, mock.Anything).Return("success login", 200).Once()
	userService.On("Login", mock.Anything, mock.Anything).Return("failed login not verified", 406).Once()
	userController := UserController{
		services: &userService,
	}
	postBody := map[string]interface{}{
		"email":    "test@gmail.com",
		"password": "123",
	}
	body, _ := json.Marshal(postBody)
	e := echo.New()
	t.Run("email not found", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc2ODY5NzUsImlkIjo0LCJyb2xlIjoxfQ.Le1-pu7wUwRAHC6qhx4ScwSmzIMd9AqzbTEHY4S8WhA"
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		userController.Login(eContext)
		assert.Equal(t, 404, w.Result().StatusCode)
	})
	t.Run("failed login", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc2ODY5NzUsImlkIjo0LCJyb2xlIjoxfQ.Le1-pu7wUwRAHC6qhx4ScwSmzIMd9AqzbTEHY4S8WhA"
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		userController.Login(eContext)
		assert.Equal(t, 401, w.Result().StatusCode)
	})
	t.Run("success login", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc2ODY5NzUsImlkIjo0LCJyb2xlIjoxfQ.Le1-pu7wUwRAHC6qhx4ScwSmzIMd9AqzbTEHY4S8WhA"
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		userController.Login(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
	t.Run("failed login", func(t *testing.T) {
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		userController.Login(eContext)
		assert.Equal(t, 406, w.Result().StatusCode)
	})
}

func TestCreateResetPassword(t *testing.T) {
	userService.On("CreateResetPassword", mock.Anything).Return(nil).Once()
	userService.On("CreateResetPassword", mock.Anything).Return(errors.New("error")).Once()
	userController := UserController{
		services: &userService,
	}
	postBody := map[string]interface{}{
		"email": "test@gmail.com",
	}
	body, _ := json.Marshal(postBody)
	e := echo.New()
	t.Run("create reset password success", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc2ODY5NzUsImlkIjo0LCJyb2xlIjoxfQ.Le1-pu7wUwRAHC6qhx4ScwSmzIMd9AqzbTEHY4S8WhA"
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		userController.CreateResetPassword(eContext)
		assert.Equal(t, 201, w.Result().StatusCode)
	})
	t.Run("create reset password error", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc2ODY5NzUsImlkIjo0LCJyb2xlIjoxfQ.Le1-pu7wUwRAHC6qhx4ScwSmzIMd9AqzbTEHY4S8WhA"
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		userController.CreateResetPassword(eContext)
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestUpdatePassword(t *testing.T) {
	userService.On("UpdatePassword", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	userService.On("UpdatePassword", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()
	userController := UserController{
		services: &userService,
	}
	postBody := map[string]interface{}{
		"email":    "test@gmail.com",
		"password": "12345",
		"kode":     "123456",
	}
	body, _ := json.Marshal(postBody)
	e := echo.New()
	t.Run("update password success", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc2ODY5NzUsImlkIjo0LCJyb2xlIjoxfQ.Le1-pu7wUwRAHC6qhx4ScwSmzIMd9AqzbTEHY4S8WhA"
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		userController.UpdatePassword(eContext)
		assert.Equal(t, 201, w.Result().StatusCode)
	})
	t.Run("update password error", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc2ODY5NzUsImlkIjo0LCJyb2xlIjoxfQ.Le1-pu7wUwRAHC6qhx4ScwSmzIMd9AqzbTEHY4S8WhA"
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		userController.UpdatePassword(eContext)
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestUpdateUserData(t *testing.T) {
	userService.On("UpdateUserData", mock.Anything, mock.Anything).Return(nil).Once()
	userService.On("UpdateUserData", mock.Anything, mock.Anything).Return(errors.New("error")).Once()
	userController := UserController{
		services: &userService,
	}
	user := models.User{
		ID:           uint(1),
		Nama:         "Denny",
		Email:        "test@gmail.com",
		Password:     "MTIz",
		NomorHP:      "12313141231",
		Kode:         "123456",
		Verifikasi:   true,
		DiBuatPada:   time.Now(),
		DiUpdatePada: time.Now(),
		RoleID:       uint(1),
	}
	body, _ := json.Marshal(user)
	e := echo.New()
	t.Run("update user data success", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc2ODY5NzUsImlkIjo0LCJyb2xlIjoxfQ.Le1-pu7wUwRAHC6qhx4ScwSmzIMd9AqzbTEHY4S8WhA"
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		userController.UpdateUserData(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
	t.Run("update user data error", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc2ODY5NzUsImlkIjo0LCJyb2xlIjoxfQ.Le1-pu7wUwRAHC6qhx4ScwSmzIMd9AqzbTEHY4S8WhA"
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		userController.UpdateUserData(eContext)
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}
