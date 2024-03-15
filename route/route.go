package route

import (
	"database/sql"
	"gokanban/services"

	"github.com/labstack/echo/v4"
)

func Init(database *sql.DB, ps services.ProjectService, ts services.TagService) *echo.Echo {
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &kanbanContext{c, database, ps, ts}
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
		project.DELETE("/:id", deleteProject)

		project.GET("/:id/board", getProjectBoards)
		project.POST("/:id/board", createProjectBoard)
	}

	board := e.Group("/board")
	{
		board.GET("/:boardid", getProjectBoard)
		board.GET("/:boardid/columns/:columnid/projectitem/:projectitemid", getProjectItem)
		board.GET("/:boardid/columns/:columnid/projectitem/new", createProjectItemForm)
		board.POST("/:boardid/columns/:columnid/projectitem/new", createProjectItem)
		board.GET("/:boardid/projectitem/:projectitemid/edit", updateProjectItemForm)
		board.PUT("/:boardid/projectitem/:projectitemid/edit", updateProjectItem)
		board.DELETE("/:boardid/columns/:columnid/projectitem/:projectitemid", deleteProjectItem)
		board.GET("/:boardid/columns/:columnid/items", getColumnItems)
	}

	validate := e.Group("/validate")
	{
		validate.GET("/title", validateTitle)
		validate.GET("/description", validateDescription)
		validate.GET("/estimatedtime", validateEstimatedtime)
		validate.GET("/spenttime", validateSpenttime)
	}

	column := e.Group("/column")
	{
		column.GET("/:columnid/count", getColumnCount)
	}

	admin := e.Group("/admin")
	{
		admin.GET("", getAdminIndex)

		admin.GET("/tags", getTags)
		admin.POST("/tags/new", createTag)
	}

	return e
}

type kanbanContext struct {
	echo.Context
	db *sql.DB
	ps services.ProjectService
	ts services.TagService
}
