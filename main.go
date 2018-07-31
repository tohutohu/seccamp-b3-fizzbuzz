package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/ping", pingpong)
	e.GET("/fizzbuzz/:number", getFizzbuzzHandler)

	e.Start(":8080")
}

func pingpong(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func getFizzbuzzHandler(c echo.Context) error {
	number, err := strconv.ParseInt(c.Param("number"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid number")
	}

	if number%15 == 0 {
		return c.String(http.StatusOK, "fizzbuzz")
	} else if number%3 == 0 {
		return c.String(http.StatusOK, "fizz")
	} else if number%5 == 0 {
		return c.String(http.StatusOK, "buzz")
	}

	return c.String(http.StatusOK, c.Param("number"))
}
