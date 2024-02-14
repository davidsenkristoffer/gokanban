package column

import (
	"gokanban/structs/projectitem"
	"time"
)

type Column struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	ColumnType  int       `json:"columntype"`
	ColumnOrder int       `json:"columnorder"`
	Created     time.Time `json:"created"`
	BoardId     int64     `json:"boardid"`
	Items       []projectitem.ProjectItem
}
