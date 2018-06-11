package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// User
	type User struct {
		Name  string `json:"name" xml:"name"`
		Email string `json:"email" xml:"email"`
	}
	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.GET("/a", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "joe@labstack.com",
		}
		return c.JSONPretty(http.StatusOK, u, "  ")
	})
	e.POST("/", func(c echo.Context) error {
		u := &User{
			Name:  c.FormValue("name"),
			Email: c.FormValue("email"),
		}
		return c.JSONPretty(http.StatusOK, u, "  ")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
