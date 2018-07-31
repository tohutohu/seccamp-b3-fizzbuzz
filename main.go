package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/ping", pingpong)

	e.Start(":8080")
}

func pingpong(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func getFizzbuzzHandler(c echo.Context) error {
	return nil
}
