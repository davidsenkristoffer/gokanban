package column

import (
	"gokanban/structs/projectitem"
	"time"
)

type Column struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	ColumnType  int       `json:"columntype"`
	ColumnOrder int       `json:"columnorder"`
	Created     time.Time `json:"created"`
	Items       []projectitem.ProjectItem
}
