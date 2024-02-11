package board

import (
	"gokanban/structs/column"
	"time"
)

type Board struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	ProjectId int       `json:"projectid"`
	Created   time.Time `json:"created"`
	Columns   []column.Column
}
