package board

import (
	"gokanban/structs/column"
	"strconv"
	t "time"
)

type Board struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	ProjectId int    `json:"projectid"`
	Created   t.Time `json:"created"`
	Columns   []column.Column
}

type BoardViewModel struct {
	Id        string
	Title     string
	ProjectId string
	Created   string
	Columns   []column.ColumnViewModel
}

func (b Board) ToViewModel() *BoardViewModel {
	columns := []column.ColumnViewModel{}
	for _, col := range b.Columns {
		columns = append(columns, *col.ToViewModel())
	}
	return &BoardViewModel{
		Id:        strconv.Itoa(b.ID),
		Title:     b.Title,
		ProjectId: strconv.Itoa(b.ProjectId),
		Created:   b.Created.In(t.Local).Format("dd.MM.yyyy"),
		Columns:   columns,
	}
}
