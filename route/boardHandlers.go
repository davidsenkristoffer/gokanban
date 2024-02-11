package route

import (
	"gokanban/db/dbboard"
	"gokanban/db/dbproject"

	"github.com/labstack/echo/v4"
)

func getProjectBoards(c echo.Context) error {
	db := c.(*kanbanContext).db
	projectid := c.Param("id")
	if len(projectid) == 0 {
		return c.JSON(400, "Bad request")
	}

	project, err := dbproject.GetProject(db, projectid)
	if err != nil {
		c.Logger().Errorf("Error while selecting project: %s", err)
		return c.JSON(404, "Project not found")
	}
	boards, err := dbboard.GetBoards(db, projectid)
	if err != nil {
		c.Logger().Errorf("Error while selecting boards for project %s: %s", projectid, err)
	} else {
		project.Boards = boards
	}

	return c.Render(200, "boards", project)
}

func createProjectBoard(c echo.Context) error {
	id := c.Param("id")
	return c.Redirect(303, "/project/"+id+"/board")
}
