package roles

import (
	"github.com/pocketbase/dbx"
)

type roleRepo struct {
	db *dbx.DB
}

func (r *roleRepo) Find(query *RoleQuery) ([]Role, error) {
	return nil, nil
}

func (r *roleRepo) FindOne(query *Role) error {
	return nil
}

func (r *roleRepo) Insert(query *Role) error {
	return nil
}

func (r *roleRepo) Update(query *Role) error {
	return nil
}

func NewRepo(db *dbx.DB) *roleRepo {
	return &roleRepo{db}
}
