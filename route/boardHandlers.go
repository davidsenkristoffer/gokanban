package route

import (
	"cmp"
	"gokanban/components"
	"gokanban/db/dbboard"
	"gokanban/db/dbcolumn"
	"gokanban/db/dbproject"
	"gokanban/helpers"
	"gokanban/structs/board"
	"gokanban/structs/column"
	"slices"
	s "strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func getProjectBoards(c echo.Context) error {
	db := c.(*kanbanContext).db
	projectid := c.Param("id")
	if _, err := s.ParseInt(projectid, 10, 64); err != nil {
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
	}

	for i, b := range boards {
		columns, err := helpers.GetColumns(db, c, b)
		if err != nil {
			c.Logger().Errorf("Error while selecting columns for board %d: %s", b.ID, err)
			continue
		} else {
			slices.SortFunc(columns, compareColumnOrder)
			boards[i].Columns = append(boards[i].Columns, columns...)
		}
	}
	project.Boards = boards

	pvm := project.ToViewModel()
	var cmp templ.Component

	hxReq := c.Request().Header.Get("HX-Request")
	if len(hxReq) > 0 {
		cmp = components.Boards(*pvm)
	} else {
		cmp = components.Kanban(*pvm)
	}
	return View(c, cmp)
}

func getProjectBoard(c echo.Context) error {
	boardid, err := s.ParseInt(c.Param("boardid"), 10, 64)
	if err != nil {
		return c.NoContent(400)
	}
	db := c.(*kanbanContext).db
	board, err := dbboard.GetBoard(db, boardid)
	if err != nil {
		return c.NoContent(404)
	}

	columns, err := helpers.GetColumns(db, c, *board)
	if err != nil {
		return c.NoContent(404)
	}

	slices.SortFunc(columns, compareColumnOrder)
	board.Columns = append(board.Columns, columns...)

	cmp := components.Board(*board.ToViewModel())

	return View(c, cmp)
}

func createProjectBoard(c echo.Context) error {
	id := c.Param("id")
	if len(id) == 0 {
		return c.JSON(400, "Bad request")
	}
	projectid, err := s.Atoi(id)
	if err != nil {
		return c.JSON(400, "Bad request")
	}
	db := c.(*kanbanContext).db

	board := &board.Board{
		Title:     c.FormValue("title"),
		ProjectId: projectid,
		Created:   time.Time{},
	}
	boardid, err := dbboard.CreateBoard(db, *board)
	if err != nil {
		return c.JSON(500, "Internal server error")
	}

	standardColumns := getStandardColumns(boardid)
	for _, col := range standardColumns {
		_, err := dbcolumn.CreateColumn(db, col)
		if err != nil {
			return c.JSON(500, "Internal server error")
		}
	}

	return c.Redirect(303, "/project/"+id+"/board")
}

func getStandardColumns(boardid int64) []column.Column {
	return []column.Column{
		{
			Title:       "New",
			ColumnType:  0,
			ColumnOrder: 0,
			Created:     time.Time{},
			BoardId:     boardid,
		},
		{
			Title:       "In progress",
			ColumnType:  1,
			ColumnOrder: 1,
			Created:     time.Time{},
			BoardId:     boardid,
		},
		{
			Title:       "Done",
			ColumnType:  2,
			ColumnOrder: 2,
			Created:     time.Time{},
			BoardId:     boardid,
		},
	}
}

func compareColumnOrder(a, b column.Column) int {
	return cmp.Compare(a.ColumnOrder, b.ColumnOrder)
}
