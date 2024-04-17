package agencies

import (
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/pocketbase/dbx"
)

type WriteAgencyRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type AgencyQuery struct {
	common.Pagination
	common.Sorter
	Search string `json:"search"`
}

type FindAgencyResp common.FindResponse[Agency]

type DeleteAgencyQuery struct {
	Ids []int `query:"ids"`
}

type Agency struct {
	common.BaseEntity
	WriteAgencyRequest
}

var _ = (*dbx.TableModel)(nil)

func (e *Agency) TableName() string {
	return "agencies"
}
