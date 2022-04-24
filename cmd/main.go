package main

import (
	"wow-api/internal"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	internal.Assign(e)
	e.Logger.Fatal(e.Start(":6000"))
}
