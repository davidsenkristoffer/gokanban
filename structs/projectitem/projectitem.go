package projectitem

import "time"

type ProjectItem struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	EstimatedTime float32   `json:"estimatedtime"`
	SpentTime     float32   `json:"spenttime"`
	ColumnId      int64     `json:"columnid"`
	Created       time.Time `json:"created"`
	Updated       time.Time `json:"updated"`
}
