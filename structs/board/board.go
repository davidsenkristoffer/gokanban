package board

import "time"

type Board struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Created time.Time `json:"created"`
}
