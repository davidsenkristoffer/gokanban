package route

import (
	"database/sql"
	"gokanban/templates"
	"html/template"

	"github.com/labstack/echo/v4"
)

func Init(database *sql.DB) *echo.Echo {
	e := echo.New()

	e.Renderer = &templates.Template{
		Templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &kanbanContext{c, database}
			return next(cc)
		}
	})
	e.Static("/static", "static")

	e.GET("/", index)

	project := e.Group("/project")
	{
		project.GET("/new", newProject)
		project.POST("/new", createProject)
		project.GET("/:id/card", getProject)

		project.GET("/:id/board", getProjectBoards)
		project.POST("/:id/board", createProjectBoard)
	}

	projectitem := e.Group("/projectitem")
	{
		projectitem.GET("/:columnid/new", CreateProjectItemForm)
		projectitem.POST("/:columnid/new", createProjectItem)
		projectitem.GET("/:projectitemid", getProjectItem)
	}

	return e
}

type kanbanContext struct {
	echo.Context
	db *sql.DB
}
