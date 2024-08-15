package signup

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type userSignUpRepo struct {
	db *dbx.DB
}

func (r *userSignUpRepo) Insert(req *User) error {
	return r.db.Model(req).Exclude(common.BaseExcludeFields...).Insert()
}

func NewRepo(db *dbx.DB) *userSignUpRepo {
	return &userSignUpRepo{db}
}
