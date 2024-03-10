package helpers

import (
	"database/sql"
	"gokanban/db/dbcolumn"
	"gokanban/db/dbprojectitem"
	"gokanban/structs/board"
	"gokanban/structs/column"
	"gokanban/structs/projectitem"
	s "strconv"
	"strings"
	t "time"

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

func FilterColumns(cols []column.Column, searchString string) []column.Column {
	if len(searchString) > 0 {
		for i, col := range cols {
			filteredItems := Filter(col.Items, func(p projectitem.ProjectItem) bool {
				return strings.Contains(strings.ToLower(p.Title), strings.ToLower(searchString))
			})
			cols[i].Items = filteredItems
		}
	}
	return cols
}

func CreateBoardStub(title string, projectid int64) *board.Board {
	return &board.Board{
		Title:     title,
		ProjectId: int(projectid),
		Created:   t.Now(),
		Columns:   []column.Column{},
	}
}
