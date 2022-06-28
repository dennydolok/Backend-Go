package main

import (
	"WallE/config"
	"WallE/handlers/rest"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	config := config.InitConfig()
	e := echo.New()
	e.AutoTLSManager.Cache = autocert.DirCache("./")
	rest.RegisterMainAPI(e, config)
	// e.Logger.Fatal(e.Start(":8080"))
	e.Logger.Fatal(e.StartAutoTLS(":443"))
}
