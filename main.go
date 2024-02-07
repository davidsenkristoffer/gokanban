package main

import (
	"fmt"
	"gokanban/db"
	"gokanban/routes"
	"gokanban/templates"
	"html/template"

	"github.com/labstack/echo/v4"
)

func main() {
	var err error
	_, err = db.Connect()
	catch(err)

	e := echo.New()
	e.Renderer = &templates.Template{
		Templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e.GET("/", routes.Index)

	e.Logger.Fatal(e.Start(":1337"))
}

func catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
