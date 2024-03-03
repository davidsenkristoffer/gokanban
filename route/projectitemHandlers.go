package route

import (
	"fmt"
	"gokanban/components"
	"gokanban/db/dbprojectitem"
	"gokanban/helpers"
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

	cmp := components.ProjectItem(*p.ToViewModel(), boardid)
	return View(c, cmp)
}

func createProjectItem(c echo.Context) error {
	boardid := c.Param("boardid")
	if _, err := strconv.ParseInt(boardid, 10, 64); err != nil {
		return c.NoContent(400)
	}
	columnid := c.Param("columnid")
	if _, err := strconv.ParseInt(columnid, 10, 64); err != nil {
		return c.NoContent(400)
	}

	pvm := &projectitem.ProjectItemViewModel{
		Title:         c.FormValue("title"),
		Description:   c.FormValue("description"),
		EstimatedTime: c.FormValue("estimatedtime"),
		ColumnId:      columnid,
	}

	validation, containsErrors := helpers.ValidateProjectItem(*pvm)
	if containsErrors {
		cmp := components.CreateProjectItem(*pvm, boardid, validation)
		return View(c, cmp)
	}

	p, err := pvm.ToModel()
	if err != nil {
		c.Logger().Errorf("Error while creating projectitem from view model: %s", err)
		return c.NoContent(500)
	}

	db := c.(*kanbanContext).db

	id, err := dbprojectitem.CreateProjectItem(db, *p)
	if err != nil {
		return err
	}

	return c.Redirect(303, fmt.Sprintf("/board/%s/projectitem/%s", boardid, strconv.FormatInt(id, 10)))
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
	columnid := c.Param("columnid")
	if _, err := strconv.ParseInt(columnid, 10, 64); err != nil {
		return c.JSON(400, "Bad request")
	}
	p := &projectitem.ProjectItemViewModel{
		Title:         "",
		Description:   "",
		EstimatedTime: "",
		SpentTime:     "",
		ColumnId:      columnid,
	}

	cmp := components.CreateProjectItem(*p, boardid, make(map[string][]string))
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
