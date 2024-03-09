package route

import (
	"gokanban/components"
	"gokanban/db/dbcolumn"
	"gokanban/db/dbprojectitem"
	"gokanban/structs/projectitem"
	s "strconv"

	"github.com/labstack/echo/v4"
)

func getColumnCount(c echo.Context) error {
	columnid := c.Param("columnid")
	if _, err := s.ParseInt(columnid, 10, 64); err != nil {
		return c.NoContent(400)
	}

	db := c.(*kanbanContext).db
	count, err := dbcolumn.GetItemCount(db, columnid)
	if err != nil {
		return c.NoContent(500)
	}

	cmp := components.Counter(s.Itoa(count))

	return View(c, cmp)
}

func getColumnItems(c echo.Context) error {
	boardid := c.Param("boardid")
	if _, err := s.ParseInt(boardid, 10, 64); err != nil {
		return c.NoContent(400)
	}

	columnid, err := s.ParseInt(c.Param("columnid"), 10, 64)
	if err != nil {
		return c.NoContent(400)
	}

	db := c.(*kanbanContext).db
	items, err := dbprojectitem.GetProjectItems(db, columnid)
	if err != nil {
		return c.NoContent(500)
	}

	pvm := []projectitem.ProjectItemViewModel{}
	for _, item := range items {
		pvm = append(pvm, *item.ToViewModel())
	}

	cmp := components.UpdatedProjectItems(pvm, boardid)
	return View(c, cmp)
}
