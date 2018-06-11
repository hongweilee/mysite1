package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

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

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t
	e.GET("/hello", Hello)
	e.Static("/", "public")
	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}
