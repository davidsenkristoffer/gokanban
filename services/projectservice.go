package services

import (
	"database/sql"
	"gokanban/db/dbboard"
	"gokanban/db/dbcolumn"
	"gokanban/db/dbproject"
	"gokanban/helpers"
	"gokanban/structs/project"
)

type IProjectService interface {
	GetProjects(db *sql.DB) ([]project.ProjectViewModel, error)
	GetProject(db *sql.DB, projectid string) (project.ProjectViewModel, error)
	CreateProject(db *sql.DB, project *project.Project) (int64, error)
	DeleteProject(db *sql.DB, projectid string) error
}

type ProjectService struct{}

func (p ProjectService) GetProjects(db *sql.DB) ([]project.ProjectViewModel, error) {
	projects, err := dbproject.GetProjects(db)
	if err != nil {
		return nil, err
	}

	pvms := []project.ProjectViewModel{}
	for _, project := range projects {
		pvms = append(pvms, *project.ToViewModel())
	}

	return pvms, nil
}

func (p ProjectService) GetProject(db *sql.DB, projectid string) (*project.ProjectViewModel, error) {
	project, err := dbproject.GetProject(db, projectid)
	if err != nil {
		return nil, err
	}

	return project.ToViewModel(), nil
}

func (p ProjectService) CreateProject(db *sql.DB, project *project.Project) (int64, error) {
	projectid, err := dbproject.CreateProject(db, *project)
	if err != nil {
		return -1, err
	}

	board := helpers.CreateBoardStub("Standard", projectid)
	boardid, err := dbboard.CreateBoard(db, *board)
	if err != nil {
		return projectid, err
	}

	columns := helpers.GetStandardColumns(boardid)
	for _, col := range columns {
		_, err := dbcolumn.CreateColumn(db, col)
		if err != nil {
			continue
		}
	}

	return projectid, nil
}

func (p ProjectService) DeleteProject(db *sql.DB, projectid string) error {
	err := dbproject.DeleteProject(db, projectid)
	if err != nil {
		return err
	}

	return nil
}
