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
	produkAPI.POST("/tambah", controllerProduk.AddProduct, middleware.RemoveTrailingSlash(), middleware.Logger())
	produkAPI.GET("", controllerProduk.GetProdukByKategoriProvider, middleware.RemoveTrailingSlash(), middleware.Logger())
	produkAPI.GET("/pilih", controllerProduk.GetPurchaseableProduct, middleware.RemoveTrailingSlash(), middleware.Logger())
	produkAPI.PUT("/update/:id", controllerProduk.UpdateProductById, middleware.RemoveTrailingSlash(), middleware.Logger())

	providerAPI := e.Group("/provider")
	providerAPI.Use(middleware.CORS())
	providerAPI.GET("/:kategori_id", controllerProduk.GetProviderByKategori, middleware.RemoveTrailingSlash(), middleware.Logger())

	kategoriAPI := e.Group("/kategori")
	kategoriAPI.Use(middleware.CORS())
	kategoriAPI.GET("", controllerProduk.GetKategori, middleware.RemoveTrailingSlash(), middleware.Logger())
	kategoriAPI.GET("/produk/:id", controllerProduk.GetProdukByKategori, middleware.RemoveTrailingSlash(), middleware.Logger())
	kategoriAPI.POST("/saldo", controllerProduk.AddSaldo, middleware.RemoveTrailingSlash(), middleware.Logger())
	kategoriAPI.GET("/saldo", controllerProduk.GetSaldo, middleware.RemoveTrailingSlash(), middleware.Logger())

	userRepository := repositories.NewUserRepository(database)
	userService := services.NewUserService(userRepository, conf)
	controllerUser := userController{
		services: userService,
	}
	userAPI := e.Group("/user")
	userAPI.Use(middleware.CORS())
	userAPI.GET("/testing", controllerUser.Testing, middleware.Logger(), middleware.JWT([]byte(conf.SECRET_KEY)))
	userAPI.POST("", controllerUser.Register, middleware.RemoveTrailingSlash(), middleware.Logger())
	userAPI.POST("/verifikasi", controllerUser.Verification, middleware.RemoveTrailingSlash(), middleware.Logger())
	userAPI.POST("/login", controllerUser.Login, middleware.RemoveTrailingSlash(), middleware.Logger())
	userAPI.POST("/reset", controllerUser.CreateResetPassword, middleware.RemoveTrailingSlash(), middleware.Logger())
	userAPI.POST("/reset/update", controllerUser.UpdatePassword, middleware.RemoveTrailingSlash(), middleware.Logger())
}
