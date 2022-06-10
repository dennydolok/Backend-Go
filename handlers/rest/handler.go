package rest

import (
	"WallE/config"
	"WallE/database"
	"WallE/repositories"
	"WallE/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterMainAPI(e *echo.Echo, conf config.Config) {
	var database = database.InitMysql(conf)

	e.GET("/check", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "It's Working   it's working",
		})
	})

	categoryRepository := repositories.NewCategoryRepository(database)
	categoryService := services.NewCategoryService(categoryRepository)
	controllerCategory := categoryController{
		services: categoryService,
	}

	categoryAPI := e.Group("/category")
	categoryAPI.GET("", controllerCategory.GetCategoriesController, middleware.RemoveTrailingSlash(), middleware.Logger())
	categoryAPI.POST("", controllerCategory.CreateCategoryController, middleware.RemoveTrailingSlash(), middleware.Logger())

	userRepository := repositories.NewUserRepository(database)
	userService := services.NewUserService(userRepository)
	controllerUser := userController{
		services: userService,
	}
	userAPI := e.Group("/user")
	userAPI.POST("", controllerUser.Register, middleware.RemoveTrailingSlash(), middleware.Logger())
	userAPI.POST("/verifikasi", controllerUser.Verification, middleware.RemoveTrailingSlash(), middleware.Logger())
}
