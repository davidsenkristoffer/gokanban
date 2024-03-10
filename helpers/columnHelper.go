package helpers

import (
	"gokanban/structs/column"
	t "time"
)

func GetStandardColumns(boardid int64) []column.Column {
	return []column.Column{
		{
			Title:       "New",
			ColumnType:  0,
			ColumnOrder: 0,
			Created:     t.Now(),
			BoardId:     boardid,
		},
		{
			Title:       "In progress",
			ColumnType:  1,
			ColumnOrder: 1,
			Created:     t.Now(),
			BoardId:     boardid,
		},
		{
			Title:       "Done",
			ColumnType:  2,
			ColumnOrder: 2,
			Created:     t.Now(),
			BoardId:     boardid,
		},
	}
}
