package route

import (
	"gokanban/components"
	"gokanban/db/dbcolumn"
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
