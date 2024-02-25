package route

import (
	"gokanban/components"
	"gokanban/db/dbprojectitem"
	"gokanban/structs/projectitem"
	"strconv"

	"github.com/labstack/echo/v4"
)

func getProjectItem(c echo.Context) error {
	db := c.(*kanbanContext).db
	p, err := dbprojectitem.GetProjectItem(db, c.Param("projectitemid"))
	if err != nil {
		return err
	}

	cmp := components.ProjectItem(*p.ToViewModel())

	return View(c, cmp)
}

func createProjectItem(c echo.Context) error {
	columnid, err := strconv.ParseInt(c.Param("columnid"), 10, 64)
	if err != nil {
		return err
	}
	estimatedTime, err := strconv.ParseFloat(c.FormValue("estimatedtime"), 64)
	if err != nil {
		estimatedTime = 0
	}
	db := c.(*kanbanContext).db

	id, err := dbprojectitem.CreateProjectItem(db, projectitem.ProjectItem{
		Title:         c.FormValue("title"),
		Description:   c.FormValue("description"),
		EstimatedTime: estimatedTime,
		SpentTime:     0,
		ColumnId:      columnid,
	})
	if err != nil {
		return err
	}

	return c.Redirect(303, "/projectitem/"+strconv.FormatInt(id, 10))
}

func CreateProjectItemForm(c echo.Context) error {
	id := c.Param("columnid")
	if _, err := strconv.ParseInt(id, 10, 64); err != nil {
		return c.JSON(400, "Bad request")
	}
	cmp := components.CreateProjectItem(id)
	return View(c, cmp)
}
