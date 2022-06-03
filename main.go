package main

import (
	"WallE/config"
	"WallE/handlers/rest"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.InitConfig()
	e := echo.New()
	rest.RegisterMainAPI(e, config)
	e.Start(":8080")
}
