package main

import (
	"gokanban/routes"
	"gokanban/templates"
	"html/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Renderer = &templates.Template{
		Templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e.GET("/", routes.Index)

	e.Logger.Fatal(e.Start(":1337"))
}
