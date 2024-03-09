package route

import (
	"fmt"
	"gokanban/components"
	"gokanban/db/dbcolumn"
	"gokanban/db/dbprojectitem"
	"gokanban/helpers"
	"gokanban/structs/projectitem"
	"gokanban/structs/selectitem"
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

	headers := c.Request().Header
	trigger := headers.Get("HX-Trigger")
	if len(trigger) > 0 {
		c.Response().Header().Set("HX-Trigger", trigger)
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

	_, err = dbprojectitem.CreateProjectItem(db, *p)
	if err != nil {
		return err
	}

	c.Response().Header().Set("HX-Trigger", fmt.Sprintf("column-updated-%s", columnid))
	return c.NoContent(201)
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
	pvm := &projectitem.ProjectItemViewModel{
		Title:         c.FormValue("title"),
		Description:   c.FormValue("description"),
		EstimatedTime: c.FormValue("estimatedtime"),
		SpentTime:     c.FormValue("spenttime"),
		ColumnId:      c.FormValue("columnid"),
	}

	validation, containsErrors := helpers.ValidateProjectItem(*pvm)
	if containsErrors {
		columns, err := dbcolumn.GetColumns(db, boardid)
		if err != nil {
			c.Logger().Errorf("Error while retrieving columns")
		}

		csi := []selectitem.Selectitem{}
		for _, c := range columns {
			csi = append(csi, *c.ToSelectOption())
		}

		cmp := components.EditProjectItem(*pvm, boardid, validation, csi)
		return View(c, cmp)
	}

	columnid, err := strconv.ParseInt(pvm.ColumnId, 10, 64)
	if err != nil {
		return c.NoContent(400)
	}

	toUpdate, err := dbprojectitem.GetProjectItem(db, projectitemid)
	if err != nil {
		return c.NoContent(404)
	}

	estimatedTime, err := strconv.ParseFloat(pvm.EstimatedTime, 64)
	if err != nil {
		estimatedTime = toUpdate.EstimatedTime
	}

	spentTime, err := strconv.ParseFloat(pvm.SpentTime, 64)
	if err != nil {
		spentTime = toUpdate.SpentTime
	}

	clone := &projectitem.ProjectItem{
		ID:            projectitemid,
		Title:         pvm.Title,
		Description:   pvm.Description,
		EstimatedTime: estimatedTime,
		SpentTime:     spentTime,
		Created:       toUpdate.Created,
		Updated:       time.Now(),
		ColumnId:      columnid,
	}

	_, err = dbprojectitem.UpdateProjectItem(db, *clone)
	if err != nil {
		return c.NoContent(500)
	}

	return c.Redirect(303, fmt.Sprintf("/board/%s", boardid))
}

func deleteProjectItem(c echo.Context) error {
	columnid := c.Param("columnid")
	if _, err := strconv.ParseInt(columnid, 10, 64); err != nil {
		return err
	}

	projectitemid, err := strconv.ParseInt(c.Param("projectitemid"), 10, 64)
	if err != nil {
		return err
	}

	db := c.(*kanbanContext).db
	_, err = dbprojectitem.DeleteProjectItem(db, int(projectitemid))
	if err != nil {
		return c.NoContent(500)
	}

	c.Response().Header().Set("HX-Trigger", fmt.Sprintf("column-updated-%s", columnid))

	return c.NoContent(204)
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
	db := c.(*kanbanContext).db
	p, err := dbprojectitem.GetProjectItem(db, projectitemid)
	if err != nil {
		return c.NoContent(404)
	}

	columns, err := dbcolumn.GetColumns(db, boardid)
	if err != nil {
		c.Logger().Errorf("Error while retrieving columns")
	}

	csi := []selectitem.Selectitem{}
	for _, c := range columns {
		csi = append(csi, *c.ToSelectOption())
	}

	cmp := components.EditProjectItem(*p.ToViewModel(), boardid, make(map[string][]string), csi)
	return View(c, cmp)
}
