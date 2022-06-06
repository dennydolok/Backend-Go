package rest

import (
	"WallE/domains"
	"WallE/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type categoryController struct {
	services domains.CategoryService
}

func (cont *categoryController) CreateCategoryController(c echo.Context) error {
	newCategory := models.Category{}
	c.Bind(newCategory)
	err := cont.services.CreateCategory(newCategory)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":     http.StatusInternalServerError,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":     http.StatusCreated,
		"messages": "success!",
	})
}

func (cont *categoryController) GetCategoriesController(c echo.Context) error {
	categories := cont.services.GetCategories()
	return c.JSON(http.StatusFound, map[string]interface{}{
		"code":       http.StatusFound,
		"categories": categories,
	})
}
