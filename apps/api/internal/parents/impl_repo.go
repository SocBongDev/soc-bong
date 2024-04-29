package parents

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type ParentRepo struct {
	db *dbx.DB
}

func (r *ParentRepo) Insert(req *Parent) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Insert()
}

func (r *ParentRepo) Update(req *Parent) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Update()
}

func NewRepo(db *dbx.DB) ParentRepository {
	return &ParentRepo{db}
}
