package main

import (
	"database/sql"
	"fmt"
	"gokanban/db"
	"gokanban/structs/project"
	"gokanban/templates"
	"html/template"
	"time"

	"github.com/labstack/echo/v4"
)

type kanbanContext struct {
	echo.Context
	db *sql.DB
}

func main() {
	var err error
	database, err := db.Connect()
	catch(err)

	e := echo.New()

	t := &templates.Template{
		Templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &kanbanContext{c, database}
			return next(cc)
		}
	})
	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", nil)
	})

	e.GET("/project/:id", getProject)
	e.GET("/project/new", func(c echo.Context) error {
		return c.Render(200, "createnewproject", nil)
	})
	e.POST("/project/new", createProject)

	e.Logger.Fatal(e.Start(":1337"))
}

func catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func getProject(c echo.Context) error {
	cc := c.(*kanbanContext)
	id := c.QueryParam("id")
	query := cc.db.QueryRow("select * from project where id = ?", id)
	project := &project.Project{}
	var err error
	if err = query.Scan(&project.ID, &project.Title, &project.Description, &project.Created); err == sql.ErrNoRows {
		return c.Render(404, "notfound", "Project with specified id not found")
	}
	return c.Render(200, "project", project)
}

func createProject(c echo.Context) error {
	cc := c.(*kanbanContext)
	project := &project.Project{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Created:     time.Now(),
	}
	query, err := cc.db.Prepare("insert into project (title, description, created) values (?, ?, ?)")
	if err != nil {
		return c.Render(500, "fatalerror", err)
	}
	defer query.Close()

	_, err = query.Exec(project.Title, project.Description, project.Created)
	if err != nil {
		return c.Render(500, "fatalerror", err)
	}

	return c.Render(200, "newproject", project)
}
