package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	var err error
	db, err := sql.Open("sqlite3", "./data.sqlite")
	if err != nil {
		return nil, err
	}

	projectStmt := `create table if not exists project(
		id integer not null primary key autoincrement, 
		title text not null, 
		description text, 
		created datetime)`
	boardStmt := `create table if not exists board (
		id integer not null primary key autoincrement, 
		title text not null,
		projectid integer not null, 
		created datetime, 
		foreign key (projectid) references project (id) on delete cascade)`
	columnStmt := `create table if not exists column (
		id integer not null primary key autoincrement, 
		title text not null, 
		columntype integer not null, 
		columnorder integer not null, 
		boardid integer not null, 
		created datetime, 
		foreign key (boardid) references board (id) on delete cascade)`
	itemStmt := `create table if not exists projectitem (
		id integer not null primary key autoincrement, 
		title text not null, 
		description text,
		tags text, 
		estimatedtime real, 
		spenttime real, 
		created datetime, 
		updated datetime, 
		columnid integer not null, 
		foreign key (columnid) references column (id))`
	tagStmt := `create table if not exists tag (
		id integer not null primary key autoincrement,
		label text not null,
		color integer not null)`

	if _, err = db.Exec(projectStmt); err != nil {
		return nil, err
	}

	if _, err = db.Exec(boardStmt); err != nil {
		return nil, err
	}

	if _, err = db.Exec(columnStmt); err != nil {
		return nil, err
	}

	if _, err = db.Exec(itemStmt); err != nil {
		return nil, err
	}

	if _, err = db.Exec(tagStmt); err != nil {
		return nil, err
	}

	return db, nil
}
