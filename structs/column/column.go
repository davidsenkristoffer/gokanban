package column

import (
	"gokanban/structs/projectitem"
	s "strconv"
	t "time"
)

type Column struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	ColumnType  int    `json:"columntype"`
	ColumnOrder int    `json:"columnorder"`
	Created     t.Time `json:"created"`
	BoardId     int64  `json:"boardid"`
	Items       []projectitem.ProjectItem
}

type ColumnViewModel struct {
	Id           string
	Title        string
	ColumnType   int
	ColumnOrder  int
	Created      string
	BoardId      string
	ProjectItems []projectitem.ProjectItemViewModel
}

func (c Column) ToViewModel() *ColumnViewModel {
	projectitems := []projectitem.ProjectItemViewModel{}
	for _, item := range c.Items {
		projectitems = append(projectitems, *item.ToViewModel())
	}
	return &ColumnViewModel{
		Id:           s.Itoa(int(c.ID)),
		Title:        c.Title,
		ColumnType:   c.ColumnType,
		ColumnOrder:  c.ColumnOrder,
		Created:      c.Created.In(t.Local).Format("yyyy-MM-dd"),
		BoardId:      s.Itoa(int(c.BoardId)),
		ProjectItems: projectitems,
	}
}
