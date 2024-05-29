package parents

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type parentRepo struct {
	db *dbx.DB
}

var _ ParentRepository = (*parentRepo)(nil)

func (r *parentRepo) Insert(req *Parent) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Insert()
}

func (r *parentRepo) Update(req *Parent) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Update()
}

func NewRepo(db *dbx.DB) *parentRepo {
	return &parentRepo{db}
}
