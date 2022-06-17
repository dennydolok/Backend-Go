package rest

import (
	"WallE/domains"
	"WallE/helper"
	"WallE/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userController struct {
	services domains.UserService
}

func (s *userController) Register(c echo.Context) error {
	newUser := models.User{}
	c.Bind(&newUser)
	err := s.services.Register(newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"kode":  http.StatusCreated,
		"pesan": "Sukses",
	})
}

func (s *userController) Verification(c echo.Context) error {
	email := c.FormValue("email")
	code := c.FormValue("code")
	fmt.Println(c.FormValue("code"))
	err := s.services.VerifikasiRegister(email, code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"kode":  http.StatusCreated,
		"pesan": "Sukses",
	})
}

func (s *userController) Login(c echo.Context) error {
	login := make(map[string]interface{})
	c.Bind(&login)
	token, code := s.services.Login(login["email"].(string), login["password"].(string))
	switch code {
	case http.StatusNotFound:
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"pesan": "Email Tidak Ditemukan",
		})
	case http.StatusUnauthorized:
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"pesan": "Gagal Login",
		})
	}
	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"pesan": "success",
		"token": token,
	})
}

func (s *userController) CreateResetPassword(c echo.Context) error {
	email := c.FormValue("email")
	fmt.Println(email)
	err := s.services.CreateResetPassword(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"kode":  http.StatusCreated,
		"pesan": "Sukses",
	})
}

func (s *userController) UpdatePassword(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	kode := c.FormValue("kode")
	err := s.services.UpdatePassword(email, password, kode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"kode":  http.StatusCreated,
		"pesan": "Sukses",
	})
}

func (s *userController) Testing(c echo.Context) error {
	reqToken := c.Request().Header.Get("Authorization")
	id, role := helper.GetClaim(reqToken)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"kode":  http.StatusCreated,
		"pesan": role,
		"id":    id,
	})
}
