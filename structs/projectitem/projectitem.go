package projectitem

import "time"

type ProjectItem struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	EstimatedTime float64   `json:"estimatedtime"`
	SpentTime     float64   `json:"spenttime"`
	Created       time.Time `json:"created"`
	Updated       time.Time `json:"updated"`
	ColumnId      int64     `json:"columnid"`
}
