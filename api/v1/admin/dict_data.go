package v1

import (
	"github.com/shichen437/live-dog/api/v1/common"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 通过字典类型获取字典数据
type GetDictDataReq struct {
	g.Meta   `path:"/system/dict/type" method:"get" tags:"字典数据" summary:"获取"`
	DictType string `p:"dictType" v:"required"`
}
type GetDictDataRes struct {
	g.Meta   `mime:"application/json"`
	DictData []*entity.SysDictData `json:"dictData"`
}

type GetDictDataListReq struct {
	g.Meta `path:"/system/dict/data/list" method:"get" tags:"字典数据" summary:"列表"`
	common.PageReq
	DictType  string `p:"dictType"`
	DictLabel string `p:"dictLabel"`
	DictValue string `p:"dictValue"`
	Status    string `p:"status"`
}
type GetDictDataListRes struct {
	g.Meta `mime:"application/json"`
	Data   []*entity.SysDictData `json:"data"`
	Total  int                   `json:"total"`
}

type GetDictDataDetailReq struct {
	g.Meta   `path:"/system/dict/data/{dictCode}" method:"get" tags:"字典数据" summary:"详情"`
	DictCode int64 `p:"dictCode"  v:"required"`
}
type GetDictDataDetailRes struct {
	g.Meta `mime:"application/json"`
	*entity.SysDictData
}

type PostDictDataReq struct {
	g.Meta    `path:"/system/dict/data" method:"post" tags:"字典数据" summary:"新增"`
	DictSort  int    `p:"dictSort" v:"required"`
	DictLabel string `p:"dictLabel" v:"required"`
	DictValue string `p:"dictValue" v:"required"`
	DictType  string `p:"dictType" v:"required"`
	CssClass  string `p:"cssClass"`
	ListClass string `p:"listClass"`
	IsDefault string `p:"isDefault"`
	Status    string `p:"status"`
	Remark    string `p:"remark"`
}
type PostDictDataRes struct {
	g.Meta `mime:"application/json"`
}

type PutDictDataReq struct {
	g.Meta    `path:"/system/dict/data" method:"put" tags:"字典数据" summary:"修改"`
	DictCode  int64  `p:"dictCode" v:"required"`
	DictSort  int    `p:"dictSort" v:"required"`
	DictLabel string `p:"dictLabel" v:"required"`
	DictValue string `p:"dictValue" v:"required"`
	DictType  string `p:"dictType" v:"required"`
	CssClass  string `p:"cssClass"`
	ListClass string `p:"listClass"`
	IsDefault string `p:"isDefault"`
	Status    string `p:"status"`
	Remark    string `p:"remark"`
}
type PutDictDataRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteDictDataReq struct {
	g.Meta   `path:"/system/dict/data/{dictCode}" method:"delete" tags:"字典数据" summary:"删除"`
	DictCode string `p:"dictCode"  v:"required"`
}
type DeleteDictDataRes struct {
	g.Meta `mime:"application/json"`
}
