package dbcolumn

import (
	"database/sql"
	"gokanban/structs/column"
)

func GetColumns(db *sql.DB, id string) ([]column.Column, error) {
	rows, err := db.Query("select * from column where boardid = ?", id)
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
