package project

import (
	"gokanban/structs/board"
	"strconv"
	t "time"
)

type Project struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     t.Time `json:"created"`
	Boards      []board.Board
}

type ProjectViewModel struct {
	Id          string
	Title       string
	Description string
	Created     string
	Boards      []board.BoardViewModel
}

func (p Project) ToViewModel() *ProjectViewModel {
	boards := []board.BoardViewModel{}
	for _, b := range p.Boards {
		boards = append(boards, *b.ToViewModel())
	}
	return &ProjectViewModel{
		Id:          strconv.Itoa(p.ID),
		Title:       p.Title,
		Description: p.Description,
		Created:     p.Created.Format(t.DateOnly),
		Boards:      boards,
	}
}
