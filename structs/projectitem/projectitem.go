package projectitem

import (
	s "strconv"
	t "time"
)

type ProjectItem struct {
	ID            int64   `json:"id"`
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
		Id:            s.FormatInt(p.ID, 10),
		Title:         p.Title,
		Description:   p.Description,
		EstimatedTime: s.FormatFloat(p.EstimatedTime, 'f', 0, 64),
		SpentTime:     s.FormatFloat(p.SpentTime, 'f', 0, 64),
		Created:       p.Created.Format(t.DateOnly),
		Updated:       p.Updated.Format(t.DateOnly),
		ColumnId:      s.FormatInt(p.ColumnId, 10),
	}
}

func (pvm ProjectItemViewModel) ToModel() (*ProjectItem, error) {
	id, err := s.ParseInt(pvm.Id, 10, 64)
	if len(pvm.Id) > 0 && err != nil {
		return nil, err
	}
	estimatedtime, err := s.ParseFloat(pvm.EstimatedTime, 64)
	if len(pvm.EstimatedTime) > 0 && err != nil {
		return nil, err
	}
	spenttime, err := s.ParseFloat(pvm.SpentTime, 64)
	if len(pvm.SpentTime) > 0 && err != nil {
		return nil, err
	}
	columnid, err := s.ParseInt(pvm.ColumnId, 10, 64)
	if len(pvm.ColumnId) > 0 && err != nil {
		return nil, err
	}
	created, err := t.Parse(t.DateOnly, pvm.Created)
	if len(pvm.Created) > 0 && err != nil {
		return nil, err
	}
	updated, err := t.Parse(t.DateOnly, pvm.Updated)
	if len(pvm.Updated) > 0 && err != nil {
		return nil, err
	}

	return &ProjectItem{
		ID:            id,
		Title:         pvm.Title,
		Description:   pvm.Description,
		EstimatedTime: estimatedtime,
		SpentTime:     spenttime,
		Created:       created,
		Updated:       updated,
		ColumnId:      columnid,
	}, nil
}
