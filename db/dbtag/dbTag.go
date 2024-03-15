package dbtag

import (
	"database/sql"
	"gokanban/structs"
)

func GetTags(db *sql.DB) ([]structs.Tag, error) {
	query, err := db.Query("select * from tag")
	if err != nil {
		return nil, err
	}

	defer query.Close()
	tags := []structs.Tag{}

	for query.Next() {
		t := &structs.Tag{}
		err = query.Scan(&t.ID, &t.Label, &t.Color)
		if err != nil {
			return nil, err
		}

		tags = append(tags, *t)
	}

	return tags, nil
}

func GetTag(db *sql.DB, tagid int64) (*structs.Tag, error) {
	query := db.QueryRow("select * from tag where id = ?", tagid)

	t := &structs.Tag{}
	err := query.Scan(&t.ID, &t.Label, &t.Color)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func CreateTag(db *sql.DB, t *structs.Tag) (int64, error) {
	stmt, err := db.Prepare("insert into tag (label, color) values (?, ?)")
	if err != nil {
		return -1, err
	}

	defer stmt.Close()
	res, err := stmt.Exec(t.Label, t.Color)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func UpdateTag(db *sql.DB, t *structs.Tag) error {
	stmt, err := db.Prepare("update tag set label = ?, color = ? where id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(t.Label, t.Color, t.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTag(db *sql.DB, tagid int64) error {
	stmt, err := db.Prepare("delete from tag where id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(tagid)
	if err != nil {
		return err
	}

	return nil
}
