package rest

import (
	"WallE/domains"
	"WallE/helper"
	"WallE/models"
	"fmt"
	"net/http"
	"strconv"

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
		if err.Error() == "resend" {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"kode":  http.StatusOK,
				"pesan": "sukses",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"kode":  http.StatusCreated,
		"pesan": "sukses",
	})
}

func (s *userController) GetUserData(c echo.Context) error {
	userid, _ := strconv.Atoi(c.Param("id"))
	user, err := s.services.GetUserDataById(uint(userid))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"kode":  http.StatusNotFound,
			"pesan": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode": http.StatusOK,
		"user": user,
	})
}

func (s *userController) Verification(c echo.Context) error {
	type body struct {
		Email string `form:"email" json:"email"`
		Code  string `form:"kode" json:"kode"`
	}
	var repBody body
	c.Bind(&repBody)
	fmt.Println(c.FormValue("code"))
	token, err := s.services.VerifikasiRegister(repBody.Email, repBody.Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":  http.StatusOK,
		"pesan": "sukses",
		"token": token,
	})
}

func (s *userController) Login(c echo.Context) error {
	login := make(map[string]interface{})
	c.Bind(&login)
	token, code := s.services.Login(login["email"].(string), login["password"].(string))
	switch code {
	case http.StatusNotFound:
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"pesan": "email tidak ditemukan",
		})
	case http.StatusUnauthorized:
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"pesan": "gagal login",
		})
	case http.StatusNotAcceptable:
		return c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"pesan": "belum verifikasi akun",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":  http.StatusOK,
		"pesan": "success",
		"token": token,
	})
}

func (s *userController) CreateResetPassword(c echo.Context) error {
	type body struct {
		Email string `form:"email" json:"email"`
	}
	var reqBody body
	c.Bind(&reqBody)
	fmt.Println(reqBody)
	err := s.services.CreateResetPassword(reqBody.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"kode":  http.StatusCreated,
		"pesan": "sukses",
	})
}

func (s *userController) UpdatePassword(c echo.Context) error {
	type body struct {
		Email    string `form:"email" json:"email"`
		Password string `form:"password" json:"password"`
		Code     string `form:"kode" json:"kode"`
	}
	var reqBody body
	c.Bind(&reqBody)
	err := s.services.UpdatePassword(reqBody.Email, reqBody.Password, reqBody.Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"kode":  http.StatusCreated,
		"pesan": "sukses",
	})
}

func (s *userController) Testing(c echo.Context) error {
	reqToken := c.Request().Header.Get("Authorization")
	role := helper.GetClaim(reqToken)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"kode":  http.StatusCreated,
		"pesan": role,
	})
}

func (s *userController) UpdateUserData(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	userId := helper.GetUserId(c.Request().Header.Get("Authorization"))
	err := s.services.UpdateUserData(uint(userId), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":  http.StatusOK,
		"pesan": "sukses",
	})
}
