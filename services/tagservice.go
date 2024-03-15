package services

import (
	"database/sql"
	"gokanban/db/dbtag"
	"gokanban/structs"
	s "strconv"
)

type TagService struct{}

func (ts TagService) GetTags(db *sql.DB) ([]structs.TagViewModel, error) {
	tags, err := dbtag.GetTags(db)
	if err != nil {
		return nil, err
	}

	tvm := []structs.TagViewModel{}
	for _, tag := range tags {
		tvm = append(tvm, *tag.ToViewModel())
	}

	return tvm, nil
}

func (ts TagService) GetTag(db *sql.DB, tagid string) (*structs.TagViewModel, error) {
	id, err := s.ParseInt(tagid, 10, 64)
	if err != nil {
		return nil, err
	}
	tag, err := dbtag.GetTag(db, id)
	if err != nil {
		return nil, err
	}

	return tag.ToViewModel(), nil
}

func (ts TagService) CreateTag(db *sql.DB, t *structs.Tag) (int64, error) {
	tagid, err := dbtag.CreateTag(db, t)
	if err != nil {
		return -1, err
	}

	return tagid, nil
}

func (ts TagService) UpdateTag(db *sql.DB, t *structs.Tag) error {
	return dbtag.UpdateTag(db, t)
}

func (ts TagService) DeleteTag(db *sql.DB, tagid string) error {
	id, err := s.ParseInt(tagid, 10, 64)
	if err != nil {
		return err
	}

	return dbtag.DeleteTag(db, id)
}
