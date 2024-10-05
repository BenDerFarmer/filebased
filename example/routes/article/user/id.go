package main

import (
	"github.com/labstack/echo/v4"
)

/*
* Change filename to "id:.go" and the dir name to ":user" to use c.Param
* https://github.com/golang/go/issues/41402
 */

func Get(c echo.Context) error {
	name := c.Param("user")
	id := c.Param("id")
	return c.String(200, name+" "+id)
}
