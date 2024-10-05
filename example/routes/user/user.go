package main

import (
	"github.com/labstack/echo/v4"
)

/*
* Change filename to "user:.go" to use c.Param
* https://github.com/golang/go/issues/41402
 */
func Get(c echo.Context) error {
	name := c.Param("user")
	return c.String(200, name)
}

func Post(c echo.Context) error {
	name := c.Param("user")
	return c.String(200, name)
}

func Put(c echo.Context) error {
	name := c.Param("user")
	return c.String(200, name)
}

func DELETE(c echo.Context) error {
	name := c.Param("user")
	return c.String(200, name)
}
