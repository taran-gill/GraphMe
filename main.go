package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Static("/", "GraphMe-Views/dist")
	e.File("/second", "GraphMe-Views/dist/second.html")

	e.Logger.Fatal(e.Start(":8081"))
}
