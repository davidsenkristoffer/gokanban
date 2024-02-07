package board

import "time"

type Board struct {
	ID      int       `json:"id"`
	Created time.Time `json:"created"`
}
