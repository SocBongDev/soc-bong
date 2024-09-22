package roles

import "github.com/SocBongDev/soc-bong/internal/common"

const TABLE = "roles"

type WriteRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BaseEntity struct {
	Id string `json:"id"`
}

type RoleQuery struct {
	common.Pagination
	common.Sorter

	Ids    []string `json:"ids"`
	Search string   `json:"search"`
}

type FindRoleResp common.FindResponse[Role]

type Role struct {
	BaseEntity
	WriteRoleRequest
}

func (e *Role) TableName() string {
	return TABLE
}
