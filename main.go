package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	//e.File("/", "./public/index.html")
	e.Static("/", "./public/")
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	e.Logger.Fatal(e.Start(":80"))
}
