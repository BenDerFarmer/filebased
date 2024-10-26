package main

import "github.com/labstack/echo/v4"

func get_world(c echo.Context) error {
	return c.String(200, "World")
}
