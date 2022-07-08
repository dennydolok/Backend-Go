package rest

import (
	"WallE/models"
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	service := MockService{}

	service.On("Register", mock.Anything).Return(errors.New("error")).Once()
	service.On("Register", mock.Anything).Return(errors.New("resend")).Once()
	service.On("Register", mock.Anything).Return(nil).Once()

	userController := UserController{
		services: &service,
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
	service := MockService{}
	user := models.User{}
	service.On("GetUserDataById").Return(user, errors.New("error")).Once()
	// service.On("GetUserDataById").Return(nil).Once()
	userController := UserController{
		services: &service,
	}

	e := echo.New()
	t.Run("error get user data", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc2ODY5NzUsImlkIjo0LCJyb2xlIjoxfQ.Le1-pu7wUwRAHC6qhx4ScwSmzIMd9AqzbTEHY4S8WhA"
		bearer := "Bearer " + token
		r := httptest.NewRequest("GET", "/", bytes.NewBuffer(nil))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		err := userController.GetUserData(eContext)
		assert.Error(t, err)
	})

	// t.Run("resend verification code", func(t *testing.T) {
	// 	r := httptest.NewRequest("GET", "/", nil)
	// 	w := httptest.NewRecorder()
	// 	eContext := e.NewContext(r, w)
	// 	err := userController.GetUserData(eContext)
	// 	assert.NoError(t, err)
	// })
}

// func TestVerification(t *testing.T) {
// 	service := MockService{}

// 	service.On("VerifikasiRegister", mock.Anything).Return("", errors.New("error")).Once()
// 	service.On("VerifikasiRegister", mock.Anything).Return("", nil).Once()

// 	userController := UserController{
// 		services: &service,
// 	}

// 	e := echo.New()
// 	t.Run("error verification", func(t *testing.T) {
// 		r := httptest.NewRequest("POST", "/", nil)
// 		w := httptest.NewRecorder()
// 		eContext := e.NewContext(r, w)
// 		userController.Verification(eContext)
// 		assert.Equal(t, 500, w.Result().StatusCode)
// 	})

// 	t.Run("ok verification", func(t *testing.T) {
// 		r := httptest.NewRequest("POST", "/", nil)
// 		w := httptest.NewRecorder()
// 		eContext := e.NewContext(r, w)
// 		userController.Verification(eContext)
// 		assert.Equal(t, 200, w.Result().StatusCode)
// 	})
// }
