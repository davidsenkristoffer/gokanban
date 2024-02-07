package project

import (
	"gokanban/board"
	"time"
)

type Project struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Board       *board.Board
}
