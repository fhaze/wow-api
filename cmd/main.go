package main

import (
	"wow-api/internal"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "X-22-KEY"},
	}))
	e.Use(middleware.Logger())
	internal.Assign(e)
	e.Logger.Fatal(e.Start(":6000"))
}
