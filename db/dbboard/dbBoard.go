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

func GetBoards(db *sql.DB, projectid int) ([]board.Board, error) {
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
	stmt, err := db.Prepare("insert into board (title, created) values (?, ?)")
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(board.Title, board.Created)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, nil
	}

	return id, nil
}
