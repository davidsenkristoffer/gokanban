package helpers

import (
	"database/sql"
	"gokanban/db/dbcolumn"
	"gokanban/db/dbprojectitem"
	"gokanban/structs/board"
	"gokanban/structs/column"
	s "strconv"

	"github.com/labstack/echo/v4"
)

func GetColumns(db *sql.DB, c echo.Context, b board.Board) ([]column.Column, error) {
	columns, err := dbcolumn.GetColumns(db, s.Itoa(b.ID))
	if err != nil {
		return nil, err
	}

	for j, col := range columns {
		items, err := dbprojectitem.GetProjectItems(db, col.ID)
		if err != nil {
			return nil, err
		}
		columns[j].Items = append(columns[j].Items, items...)
	}

	return columns, nil
}
