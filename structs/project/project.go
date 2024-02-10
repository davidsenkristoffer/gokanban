package project

import (
	"gokanban/structs/board"
	"time"
)

type Project struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Boards      []board.Board
}
