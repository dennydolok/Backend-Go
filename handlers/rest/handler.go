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
	produkRepository := repositories.NewProductRepository(database)
	produkServices := services.NewProdukService(produkRepository)
	controllerProduk := productController{
		services: produkServices,
	}
	produkAPI := e.Group("/produk")
	produkAPI.Use(middleware.CORS())
	produkAPI.POST("/tambah", controllerProduk.TambahProduk, middleware.RemoveTrailingSlash(), middleware.Logger())
	produkAPI.GET("/kategori", controllerProduk.AmbilProdukBerdasarkanKategori, middleware.RemoveTrailingSlash(), middleware.Logger())
	produkAPI.POST("/saldo", controllerProduk.TambahSaldo, middleware.RemoveTrailingSlash(), middleware.Logger())
	userRepository := repositories.NewUserRepository(database)
	userService := services.NewUserService(userRepository, conf)
	controllerUser := userController{
		services: userService,
	}
	userAPI := e.Group("/user")
	userAPI.Use(middleware.CORS())
	userAPI.POST("", controllerUser.Register, middleware.RemoveTrailingSlash(), middleware.Logger())
	userAPI.POST("/verifikasi", controllerUser.Verification, middleware.RemoveTrailingSlash(), middleware.Logger())
	userAPI.POST("/login", controllerUser.Login, middleware.RemoveTrailingSlash(), middleware.Logger())
	userAPI.POST("/reset", controllerUser.CreateResetPassword, middleware.RemoveTrailingSlash(), middleware.Logger())
	userAPI.POST("/reset/update", controllerUser.UpdatePassword, middleware.RemoveTrailingSlash(), middleware.Logger())
}
