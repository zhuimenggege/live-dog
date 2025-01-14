package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shichen437/live-dog/api/v1/common"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"
)

type GetSysConfigListReq struct {
	g.Meta     `path:"/system/sys/config/list" method:"Get" tags:"系统配置" summary:"获取配置列表"`
	ConfigName string      `p:"configName"`
	ConfigKey  string      `p:"configKey"`
	ConfigType string      `p:"configType"`
	CreateTime *gtime.Time `p:"createTime"`
	common.PageReq
}

type GetSysConfigListRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*entity.SysConfig `json:"rows"`
	Total  int                 `json:"total"`
}

type PostSysConfigReq struct {
	g.Meta      `path:"/system/sys/config" method:"Post" tags:"系统配置" summary:"新增"`
	ConfigName  string `p:"configName" `
	ConfigKey   string `p:"configKey" `
	ConfigValue string `p:"configValue" `
	ConfigType  string `p:"configType" `
	Remark      string `p:"remark" `
}

type PostSysConfigRes struct {
	g.Meta `mime:"application/json"`
}

type PutSysConfigReq struct {
	g.Meta      `path:"/system/sys/config" method:"Put" tags:"系统配置" summary:"修改"`
	ConfigId    int    `p:"configId"  v:"required" `
	ConfigName  string `p:"configName" `
	ConfigKey   string `p:"configKey" `
	ConfigValue string `p:"configValue" `
	ConfigType  string `p:"configType" `
	Remark      string `p:"remark" `
}

type PutSysConfigRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteSysConfigReq struct {
	g.Meta   `path:"/system/sys/config/{configId}" method:"Delete" tags:"系统配置" summary:"删除"`
	ConfigId string `p:"configId"  v:"required" `
}

type DeleteSysConfigRes struct {
	g.Meta `mime:"application/json"`
}

type GetSysConfigReq struct {
	g.Meta   `path:"/system/sys/config/{configId}" method:"Get" tags:"系统配置" summary:"获取"`
	ConfigId int `p:"configId"  v:"required" `
}

type GetSysConfigRes struct {
	g.Meta `mime:"application/json"`
	*entity.SysConfig
}
