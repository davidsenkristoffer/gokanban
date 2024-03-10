package helpers

import (
	"gokanban/structs/project"
	s "strconv"
	t "time"
)

func VerifyProjectId(id string) (bool, error) {
	if _, err := s.ParseInt(id, 10, 64); err != nil {
		return false, err
	}

	return true, nil
}

func CreateProjectStub(title string, description string) *project.Project {
	return &project.Project{
		Title:       title,
		Description: description,
		Created:     t.Now(),
	}
}
