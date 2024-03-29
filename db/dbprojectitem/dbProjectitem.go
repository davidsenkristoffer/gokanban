package dbprojectitem

import (
	"database/sql"
	"gokanban/db/dbtag"
	"gokanban/structs"
	"gokanban/structs/projectitem"
	"strconv"
	"time"
)

func GetProjectItem(db *sql.DB, id int64) (*projectitem.ProjectItem, error) {
	row := db.QueryRow("select * from projectitem where id = ?", id)
	p := &projectitem.ProjectItem{}
	err := row.Scan(&p.ID, &p.Title, &p.Description, &p.Tags, &p.EstimatedTime, &p.SpentTime, &p.Created, &p.Updated, &p.ColumnId)
	if err != nil {
		return nil, err
	}

	if len(p.Tags) > 0 {
		t, err := getTag(db, p.Tags)
		if err == nil {
			p.Taglist = append(p.Taglist, *t)
		}
	}

	return p, nil
}

func GetProjectItems(db *sql.DB, columnid int64) ([]projectitem.ProjectItem, error) {
	rows, err := db.Query("select * from projectitem where columnid = ?", columnid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projectitems := []projectitem.ProjectItem{}

	for rows.Next() {
		p := &projectitem.ProjectItem{}
		err = rows.Scan(&p.ID, &p.Title, &p.Description, &p.Tags, &p.EstimatedTime, &p.SpentTime, &p.Created, &p.Updated, &p.ColumnId)
		if err != nil {
			return nil, err
		}

		if len(p.Tags) > 0 {
			t, err := getTag(db, p.Tags)
			if err == nil {
				p.Taglist = append(p.Taglist, *t)
			}
		}

		projectitems = append(projectitems, *p)
	}

	return projectitems, nil
}

func CreateProjectItem(db *sql.DB, p projectitem.ProjectItem) (int64, error) {
	stmt, err := db.Prepare("insert into projectitem (title, description, tags, estimatedtime, spenttime, created, updated, columnid) values (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec(&p.Title, &p.Description, &p.Tags, &p.EstimatedTime, &p.SpentTime, time.Now(), time.Now(), &p.ColumnId)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func UpdateProjectItem(db *sql.DB, p projectitem.ProjectItem) (int64, error) {
	stmt, err := db.Prepare("update projectitem set title = ?, description = ?, tags = ?, estimatedtime = ?, spenttime = ?, columnid = ?, updated = ? where id = ?")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(&p.Title, &p.Description, &p.Tags, &p.EstimatedTime, &p.SpentTime, &p.ColumnId, time.Now(), &p.ID)
	if err != nil {
		return 0, err
	}
	i, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return i, nil
}

func DeleteProjectItem(db *sql.DB, id int) (int64, error) {
	stmt, err := db.Prepare("delete from projectitem where id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}
	i, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return i, nil
}

func getTag(db *sql.DB, tag string) (*structs.Tag, error) {
	id, err := strconv.ParseInt(tag, 10, 64)
	if err != nil {
		return nil, err
	}
	t, err := dbtag.GetTag(db, id)
	if err != nil {
		return nil, err
	}

	return t, nil
}
