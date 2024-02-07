package main

import (
	"database/sql"
)

func connect() (*sql.DB, error) {
	var err error
	db, err := sql.Open("sqlite3", "./data.sqlite")
	if err != nil {
		return nil, err
	}

	projectStmt := `create table if not exists project(id integer not null primary key autoincrement, title text, description text, created datetime)`
	boardStmt := `create table if not exists board (id integer not null primary key autoincrement, projectid integer not null, created datetime, foreign key (projectid) references project (id))`
	columnStmt := `create table if not exists column (id integer not null primary key autoincrement, title text, columntype integer, boardId integer not null, created datetime, foreign key (boardId) references board (id))`
	itemStmt := `create table if not exists projectitem (id integer not null primary key autoincrement, title text not null, description text, estimatedtime real, spenttime real, created datetime, updated datetime, columnid integer not null, projectid integer not null, foreign key (project) references project (id))`

	_, err = db.Exec(projectStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(boardStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(columnStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(itemStmt)
	if err != nil {
		return nil, err
	}

	return db, nil
}
