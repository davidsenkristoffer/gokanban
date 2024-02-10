package main

import (
	"database/sql"
	"fmt"
	"gokanban/db"
	"gokanban/db/dbproject"
	"gokanban/structs/project"
	"gokanban/templates"
	"html/template"
	"strconv"
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

	e.GET("/", index)

	e.GET("/project/new", newProject)
	e.POST("/project/new", createProject)
	e.GET("/project/:id/card", getProject)

	e.GET("/project/:id/board", getProjectBoards)
	e.POST("/project/:id/board", createProjectBoard)

	e.Logger.Fatal(e.Start(":1337"))
}

func catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func index(c echo.Context) error {
	projects, err := dbproject.GetProjects(c.(*kanbanContext).db)
	if err != nil {
		return c.Render(500, "fatalerror", err)
	}
	return c.Render(200, "index", projects)
}

func getProject(c echo.Context) error {
	id := c.Param("id")
	if len(id) == 0 {
		return c.JSON(400, "Bad request")
	}
	project, err := dbproject.GetProject(c.(*kanbanContext).db, id)
	if err != nil {
		return c.JSON(404, "Resource not found")
	}

	return c.Render(200, "projectcard", project)
}

func newProject(c echo.Context) error {
	return c.Render(200, "createnewproject", nil)
}

func createProject(c echo.Context) error {
	project := &project.Project{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Created:     time.Now(),
	}
	id, err := dbproject.CreateProject(c.(*kanbanContext).db, *project)
	if err != nil {
		return c.JSON(500, "Internal server error")
	} else if id == -1 {
		return c.Redirect(303, "/")
	}

	return c.Redirect(303, "/project/"+strconv.FormatInt(id, 10)+"/card")
}

func getProjectBoards(c echo.Context) error {
	return c.Render(200, "boards", nil)
}

func createProjectBoard(c echo.Context) error {
	id := c.Param("id")
	return c.Redirect(303, "/project/"+id+"/board")
}
