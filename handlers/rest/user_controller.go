package rest

import (
	"WallE/domains"
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
			"code":     http.StatusInternalServerError,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":     http.StatusCreated,
		"messages": "success !",
	})
}

func (s *userController) Verification(c echo.Context) error {
	email := c.FormValue("email")
	code := c.FormValue("code")
	fmt.Println(c.FormValue("code"))
	err := s.services.VerifikasiRegister(email, code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":     http.StatusInternalServerError,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":     http.StatusCreated,
		"messages": "success !",
	})
}
