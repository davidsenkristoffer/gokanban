package route

import (
	"gokanban/components"
	"gokanban/helpers"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
	ps := c.(*kanbanContext).ps
	db := c.(*kanbanContext).db

	pvms, err := ps.GetProjects(db)
	if err != nil {
		return c.NoContent(404)
	}

	cmp := components.Index(pvms)
	return View(c, cmp)
}

func getProject(c echo.Context) error {
	projectid := c.Param("id")
	if _, err := helpers.VerifyProjectId(projectid); err != nil {
		c.Logger().Errorf("Param \"id\" was not formatted correctly. Received value: %s", projectid)
		return c.JSON(400, "Bad request")
	}

	db := c.(*kanbanContext).db
	ps := c.(*kanbanContext).ps

	project, err := ps.GetProject(db, projectid)
	if err != nil {
		return c.JSON(404, "Resource not found")
	}

	cmp := components.Projectcard(*project)
	return View(c, cmp)
}

func newProject(c echo.Context) error {
	cmp := components.CreateNewProject()
	return View(c, cmp)
}

func createProject(c echo.Context) error {
	db := c.(*kanbanContext).db
	ps := c.(*kanbanContext).ps

	project := helpers.CreateProjectStub(c.FormValue("title"), c.FormValue("description"))
	_, err := ps.CreateProject(db, project)
	if err != nil {
		c.Logger().Errorf("Error while creating project: %s", err)
		return c.NoContent(500)
	}

	c.Response().Header().Set("HX-Trigger", "project-updated")

	return c.NoContent(201)
}

func View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
