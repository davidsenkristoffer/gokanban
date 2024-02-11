package route

import (
	"gokanban/db/dbboard"
	"gokanban/db/dbproject"
	"gokanban/structs/board"
	"gokanban/structs/project"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

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
	db := c.(*kanbanContext).db

	project := &project.Project{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Created:     time.Now(),
	}

	projectid, err := dbproject.CreateProject(db, *project)
	if err != nil {
		c.Logger().Errorf("Error while creating project: %s", err)
		return c.JSON(500, "Internal server error")
	} else if projectid == -1 {
		return c.Redirect(303, "/")
	}

	board := &board.Board{
		Title:     "Standard",
		Created:   time.Time{},
		ProjectId: int(projectid),
	}

	_, err = dbboard.CreateBoard(db, *board)
	if err != nil {
		c.Logger().Errorf("Error while creating board for project with id %d: %s", projectid, err)
		return c.JSON(500, "Internal server error")
	}

	return c.Redirect(303, "/project/"+strconv.FormatInt(projectid, 10)+"/card")
}
