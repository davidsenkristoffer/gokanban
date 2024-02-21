package dbcolumn

import (
	"database/sql"
	"gokanban/structs/board"
	"gokanban/structs/column"
	"gokanban/structs/projectitem"
)

func GetColumns(db *sql.DB, board board.Board) ([]column.Column, error) {
	rows, err := db.Query("select * from column where boardid = ?", board.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns := []column.Column{}

	for rows.Next() {
		column := &column.Column{}
		err = rows.Scan(&column.ID, &column.Title, &column.ColumnType, &column.ColumnOrder, &column.BoardId, &column.Created)
		if err != nil {
			return nil, err
		}
		columns = append(columns, *column)
	}

	for _, c := range columns {
		itemRows, err := db.Query("select * from projectitem inner join column where projectitem.columnid = ?", c.ID)
		if err != nil {
			return columns, nil
		}
		defer itemRows.Close()

		for rows.Next() {
			projectitem := &projectitem.ProjectItem{}
			err = rows.Scan(
				&projectitem.ID,
				&projectitem.Title,
				&projectitem.Description,
				&projectitem.EstimatedTime,
				&projectitem.SpentTime,
				&projectitem.Updated,
				&projectitem.Created)
			if err != nil {
				continue
			}
			c.Items = append(c.Items, *projectitem)
		}
	}
	return columns, nil
}

func CreateColumn(db *sql.DB, column column.Column) (int64, error) {
	stmt, err := db.Prepare("insert into column (title, columntype, columnorder, boardid, created) values (?, ?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(&column.Title, &column.ColumnType, &column.ColumnOrder, &column.BoardId, &column.Created)
	if err != nil {
		return -1, err
	}

	columnid, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return columnid, nil
}
