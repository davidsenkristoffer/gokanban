package route

import (
	"fmt"
	"gokanban/components"
	"gokanban/db/dbprojectitem"
	"gokanban/structs/projectitem"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func getProjectItem(c echo.Context) error {
	boardid := c.Param("boardid")
	if _, err := strconv.ParseInt(boardid, 10, 64); err != nil {
		return c.NoContent(400)
	}
	projectitemid, err := strconv.ParseInt(c.Param("projectitemid"), 10, 64)
	if err != nil {
		return err
	}

	db := c.(*kanbanContext).db
	p, err := dbprojectitem.GetProjectItem(db, projectitemid)
	if err != nil {
		return err
	}

	hxReq := c.Request().Header.Get("HX-Request")
	cmp := components.ProjectItem(*p.ToViewModel(), boardid, len(hxReq) > 0)

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

func updateProjectItem(c echo.Context) error {
	boardid := c.Param("boardid")
	if _, err := strconv.ParseInt(boardid, 10, 64); err != nil {
		return c.NoContent(400)
	}
	projectitemid, err := strconv.ParseInt(c.Param("projectitemid"), 10, 64)
	if err != nil {
		return err
	}

	db := c.(*kanbanContext).db
	toUpdate, err := dbprojectitem.GetProjectItem(db, projectitemid)
	if err != nil {
		return c.NoContent(404)
	}

	estimatedTime, err := strconv.ParseFloat(c.FormValue("estimatedtime"), 64)
	if err != nil {
		estimatedTime = toUpdate.EstimatedTime
	}

	spentTime, err := strconv.ParseFloat(c.FormValue("spenttime"), 64)
	if err != nil {
		spentTime = toUpdate.SpentTime
	}

	columnId, err := strconv.ParseInt(c.FormValue("columnid"), 10, 64)
	if err != nil {
		columnId = toUpdate.ColumnId
	}

	clone := &projectitem.ProjectItem{
		ID:            projectitemid,
		Title:         c.FormValue("title"),
		Description:   c.FormValue("description"),
		EstimatedTime: estimatedTime,
		SpentTime:     spentTime,
		Updated:       time.Now(),
		ColumnId:      columnId,
	}

	_, err = dbprojectitem.UpdateProjectItem(db, *clone)
	if err != nil {
		return c.NoContent(500)
	}

	return c.Redirect(303, fmt.Sprintf("/board/%s", boardid))
}

func deleteProjectItem(c echo.Context) error {
	projectitemid, err := strconv.ParseInt(c.Param("projectitemid"), 10, 64)
	if err != nil {
		return err
	}

	db := c.(*kanbanContext).db
	_, err = dbprojectitem.DeleteProjectItem(db, int(projectitemid))
	if err != nil {
		return c.NoContent(500)
	}

	return c.NoContent(200)
}

func createProjectItemForm(c echo.Context) error {
	boardid := c.Param("boardid")
	if _, err := strconv.ParseInt(boardid, 10, 64); err != nil {
		return c.NoContent(400)
	}
	id := c.Param("columnid")
	if _, err := strconv.ParseInt(id, 10, 64); err != nil {
		return c.JSON(400, "Bad request")
	}
	cmp := components.CreateProjectItem(id, boardid)
	return View(c, cmp)
}

func updateProjectItemForm(c echo.Context) error {
	boardid := c.Param("boardid")
	if _, err := strconv.ParseInt(boardid, 10, 64); err != nil {
		return c.NoContent(400)
	}
	projectitemid, err := strconv.ParseInt(c.Param("projectitemid"), 10, 64)
	if err != nil {
		return c.NoContent(400)
	}
	p, err := dbprojectitem.GetProjectItem(c.(*kanbanContext).db, projectitemid)
	if err != nil {
		return c.NoContent(404)
	}

	cmp := components.EditProjectItem(*p.ToViewModel(), boardid)
	return View(c, cmp)
}
