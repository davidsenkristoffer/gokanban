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

	e.Group("/project")
	{
		e.GET("/new", newProject)
		e.POST("/new", createProject)
		e.GET("/:id/card", getProject)

		e.GET("/:id/board", getProjectBoards)
		e.POST("/:id/board", createProjectBoard)
	}

	return e
}

type kanbanContext struct {
	echo.Context
	db *sql.DB
}
