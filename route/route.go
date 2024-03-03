package route

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func Init(database *sql.DB) *echo.Echo {
	e := echo.New()

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

	board := e.Group("/board")
	{
		board.GET("/:boardid", getProjectBoard)
		board.GET("/:boardid/projectitem/:projectitemid", getProjectItem)
		board.GET("/:boardid/projectitem/:columnid/new", createProjectItemForm)
		board.POST("/:boardid/projectitem/:columnid/new", createProjectItem)
		board.GET("/:boardid/projectitem/:projectitemid/edit", updateProjectItemForm)
		board.PUT("/:boardid/projectitem/:projectitemid/edit", updateProjectItem)
		board.DELETE("/:boardid/projectitem/:projectitemid", deleteProjectItem)
	}

	validate := e.Group("/validate")
	{
		validate.GET("/title", validateTitle)
		validate.GET("/description", validateDescription)
		validate.GET("/estimatedtime", validateEstimatedtime)
		validate.GET("/spenttime", validateSpenttime)
	}

	return e
}

type kanbanContext struct {
	echo.Context
	db *sql.DB
}
