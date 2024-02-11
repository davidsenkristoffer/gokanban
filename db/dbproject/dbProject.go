package dbproject

import (
	"database/sql"
	"gokanban/structs/project"
)

func GetProject(db *sql.DB, id string) (*project.Project, error) {
	query := db.QueryRow("select * from project where id = ?", id)
	project := &project.Project{}
	if err := query.Scan(&project.ID, &project.Title, &project.Description, &project.Created); err == sql.ErrNoRows {
		return nil, err
	}
	return project, nil
}

func GetProjects(db *sql.DB) ([]project.Project, error) {
	rows, err := db.Query("select * from project")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := []project.Project{}

	for rows.Next() {
		project := new(project.Project)
		err = rows.Scan(&project.ID, &project.Title, &project.Description, &project.Created)
		if err != nil {
			return nil, err
		}
		projects = append(projects, *project)
	}

	return projects, nil
}

func CreateProject(db *sql.DB, project project.Project) (int64, error) {
	query, err := db.Prepare("insert into project (title, description, created) values (?, ?, ?)")
	if err != nil {
		return -1, err
	}
	defer query.Close()

	res, err := query.Exec(project.Title, project.Description, project.Created)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, nil
	}

	return id, nil
}
