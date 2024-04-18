package agencies

import "github.com/SocBongDev/soc-bong/internal/common"

type WriteAgencyRequest struct {
	Address string `json:"address"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
}

type FindAgencyResp common.FindResponse[Agency]

type AgencyQuery struct {
	common.Pagination
	common.Sorter
	Search string `json:"search"`
}

type Agency struct {
	common.BaseEntity
	WriteAgencyRequest
}

func (e *Agency) TableName() string {
	return "agencies"
}
