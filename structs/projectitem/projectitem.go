package projectitem

import (
	s "strconv"
	t "time"
)

type ProjectItem struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	EstimatedTime float64 `json:"estimatedtime"`
	SpentTime     float64 `json:"spenttime"`
	Created       t.Time  `json:"created"`
	Updated       t.Time  `json:"updated"`
	ColumnId      int64   `json:"columnid"`
}

type ProjectItemViewModel struct {
	Id            string
	Title         string
	Description   string
	EstimatedTime string
	SpentTime     string
	Created       string
	Updated       string
	ColumnId      string
}

func (p ProjectItem) ToViewModel() *ProjectItemViewModel {
	return &ProjectItemViewModel{
		Id:            s.Itoa(p.ID),
		Title:         p.Title,
		Description:   p.Description,
		EstimatedTime: s.FormatFloat(p.EstimatedTime, 'f', 0, 64),
		SpentTime:     s.FormatFloat(p.SpentTime, 'f', 0, 64),
		Created:       p.Created.In(t.Local).Format("dd.MM.yyyy"),
		Updated:       p.Updated.In(t.Local).Format("dd.MM.yyyy"),
		ColumnId:      s.FormatInt(p.ColumnId, 10),
	}
}
