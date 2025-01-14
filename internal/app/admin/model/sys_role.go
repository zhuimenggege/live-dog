package model

import "github.com/shichen437/live-dog/internal/app/admin/model/entity"

type SysRoleRes struct {
	*entity.SysRole
	Flag        bool     `json:"flag" `
	MenuIds     string   `json:"menuIds" `
	Admin       bool     `json:"admin"            description:"是否是admin"`
	Permissions []string `json:"permissions"            description:"权限"`
}
type SysRolesRes struct {
	RoleIds   []int64           `json:"roleIds" `
	Roles     []string          `json:"roles" `
	RoleNames []string          `json:"roleNames" `
	SysRole   []*entity.SysRole `json:"SysRole" `
}

type RoleList struct {
	*entity.SysRole
	Flag    bool   `json:"flag" `
	MenuIds string `json:"menuIds" `
	Admin   bool   `json:"admin"            description:"是否是admin"`
}
