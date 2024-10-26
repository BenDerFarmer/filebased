package main

import (
	"github.com/labstack/echo/v4"
)

/*
* Change filename to "user:.go" to use c.Param
* https://github.com/golang/go/issues/41402
 */
func get(c echo.Context) error {
	name := c.Param("user")
	return c.String(200, name)
}

func post(c echo.Context) error {
	name := c.Param("user")
	return c.String(200, name)
}

func put(c echo.Context) error {
	name := c.Param("user")
	return c.String(200, name)
}

func delete(c echo.Context) error {
	name := c.Param("user")
	return c.String(200, name)
}
