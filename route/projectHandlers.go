package route

import (
	"gokanban/components"
	"gokanban/db/dbboard"
	"gokanban/db/dbcolumn"
	"gokanban/db/dbproject"
	"gokanban/structs/board"
	"gokanban/structs/project"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
	projects, err := dbproject.GetProjects(c.(*kanbanContext).db)
	if err != nil {
		return c.JSON(404, "Resource not found")
	}
	pvms := []project.ProjectViewModel{}
	for _, p := range projects {
		pvms = append(pvms, *p.ToViewModel())
	}

	cmp := components.Index(pvms)
	return View(c, cmp)
}

func getProject(c echo.Context) error {
	id := c.Param("id")
	if _, err := strconv.ParseInt(id, 10, 64); err != nil {
		c.Logger().Errorf("Param \"id\" was not formatted correctly. Received value: %s", id)
		return c.JSON(400, "Bad request")
	}
	project, err := dbproject.GetProject(c.(*kanbanContext).db, id)
	if err != nil {
		return c.JSON(404, "Resource not found")
	}

	pvm := project.ToViewModel()
	cmp := components.Projectcard(*pvm)
	return View(c, cmp)
}

func newProject(c echo.Context) error {
	cmp := components.CreateNewProject()
	return View(c, cmp)
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
		Created:   time.Now(),
		ProjectId: int(projectid),
	}

	boardid, err := dbboard.CreateBoard(db, *board)
	if err != nil {
		c.Logger().Errorf("Error while creating board for project with id %d: %s", projectid, err)
		return c.JSON(500, "Internal server error")
	}

	columns := getStandardColumns(boardid)
	for _, column := range columns {
		_, err := dbcolumn.CreateColumn(db, column)
		if err != nil {
			c.Logger().Errorf("Error while creating column for board with id %d: %s", boardid, err)
			return c.JSON(500, "Internal server error")
		}
	}

	return c.Redirect(303, "/project/"+strconv.FormatInt(projectid, 10)+"/card")
}

func View(c echo.Context, cmp templ.Component) error {
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
