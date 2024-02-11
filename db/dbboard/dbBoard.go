package dbboard

import (
	"database/sql"
	"gokanban/structs/board"
)

func GetBoard(db *sql.DB, id int) (*board.Board, error) {
	stmt := db.QueryRow("select id, title, created from board where id = ?", id)
	board := &board.Board{}
	err := stmt.Scan(&board.ID, &board.Title, &board.Created)
	if err != nil {
		return nil, err
	}

	return board, nil
}

func GetBoards(db *sql.DB, projectid string) ([]board.Board, error) {
	stmt, err := db.Query("select id, title, created from board where projectid = ?", projectid)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	boards := []board.Board{}
	for stmt.Next() {
		board := new(board.Board)
		err = stmt.Scan(&board.ID, &board.Title, &board.Created)
		if err != nil {
			return nil, err
		}
		boards = append(boards, *board)
	}

	return boards, nil
}

func CreateBoard(db *sql.DB, board board.Board) (int64, error) {
	stmt, err := db.Prepare("insert into board (title, projectid, created) values (?, ?, ?)")
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(&board.Title, &board.ProjectId, &board.Created)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	id, err := res.LastInsertId()
	if err != nil {
		return -1, nil
	}

	return id, nil
}

func UpdateBoard(db *sql.DB, board board.Board) error {
	stmt, err := db.Prepare("update board set title = ? where id = ?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(&board.Title, &board.ID)
	if err != nil {
		return err
	}
	if row, err := res.RowsAffected(); err != nil {
		return err
	} else if row == 0 {
		return sql.ErrNoRows
	}
	return nil
}
