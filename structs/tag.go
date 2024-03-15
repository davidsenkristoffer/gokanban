package structs

import s "strconv"

type Tag struct {
	ID    int64  `json:"id"`
	Label string `json:"label"`
	Color int    `json:"color"`
}

type TagViewModel struct {
	Id    string
	Label string
	Color string
}

func (t Tag) ToViewModel() *TagViewModel {
	return &TagViewModel{
		Id:    s.Itoa(int(t.ID)),
		Label: t.Label,
		Color: s.Itoa(t.Color),
	}
}

func (t TagViewModel) ToModel() (*Tag, error) {
	id, err := s.ParseInt(t.Id, 10, 64)
	if len(t.Id) > 0 && err != nil {
		return nil, err
	}
	color, err := s.Atoi(t.Color)
	if err != nil {
		return nil, err
	}

	return &Tag{
		ID:    id,
		Label: t.Label,
		Color: color,
	}, nil
}
