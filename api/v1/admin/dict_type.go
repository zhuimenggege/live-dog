package v1

import (
	"github.com/shichen437/live-dog/api/v1/common"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetSysDictTypeListReq struct {
	g.Meta     `path:"/system/dict/type/list" method:"Get" tags:"字典类型" summary:"列表"`
	CreateTime *gtime.Time `p:"createTime"`
	DictName   string      `p:"dictName"`
	DictType   string      `p:"dictType"`
	Status     string      `p:"status"`
	common.PageReq
}

type GetSysDictTypeListRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.SysDictType `json:"list"`
	Total  int                   `json:"total"`
}

type PostSysDictTypeReq struct {
	g.Meta   `path:"/system/dict/type" method:"Post" tags:"字典类型" summary:"新增"`
	Remark   string `p:"remark" `
	DictName string `p:"dictName" `
	DictType string `p:"dictType" `
	Status   string `p:"status" `
}

type PostSysDictTypeRes struct {
	g.Meta `mime:"application/json"`
}

type PutSysDictTypeReq struct {
	g.Meta   `path:"/system/dict/type" method:"Put" tags:"字典类型" summary:"修改"`
	Remark   string `p:"remark" `
	DictName string `p:"dictName" `
	DictType string `p:"dictType" `
	DictId   int64  `p:"dictId"  v:"required" `
	Status   string `p:"status" `
}

type PutSysDictTypeRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteSysDictTypeReq struct {
	g.Meta `path:"/system/dict/type/{dictId}" method:"Delete" tags:"字典类型" summary:"删除"`
	DictId string `p:"dictId"  v:"required" `
}

type DeleteSysDictTypeRes struct {
	g.Meta `mime:"application/json"`
}

type GetSysDictTypeReq struct {
	g.Meta `path:"/system/dict/type/{dictId}" method:"Get" tags:"字典类型" summary:"详情"`
	DictId int64 `p:"dictId"  v:"required" `
}

type GetSysDictTypeRes struct {
	g.Meta `mime:"application/json"`
	*entity.SysDictType
}

// 下拉框获取字典类型
type GetDictTypeOptionSelectReq struct {
	g.Meta `path:"/system/dict/type/optionselect" method:"get" tags:"字典类型" summary:"下拉框获取"`
}
type GetDictTypeOptionSelectRes struct {
	g.Meta   `mime:"application/json"`
	DictType []*entity.SysDictType `json:"dictType"`
}
