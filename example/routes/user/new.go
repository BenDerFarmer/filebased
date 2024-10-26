package main

import (
	"github.com/labstack/echo/v4"
)

func post_new(c echo.Context) error {
	return c.String(200, "post new user")
}
