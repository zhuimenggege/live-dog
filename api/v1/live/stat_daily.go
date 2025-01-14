package v1

import (
	"github.com/shichen437/live-dog/api/v1/common"
	"github.com/shichen437/live-dog/internal/app/live/model"
	"github.com/shichen437/live-dog/internal/app/live/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetStatDailyListReq struct {
	g.Meta `path:"/live/daily/list" method:"get" tags:"每日统计" summary:"列表"`
	common.PageReq
	Anchor      string `p:"anchor"`
	DisplayName string `p:"displayName"`
	DisplayType int    `p:"displayType"`
	DisplayDate string `p:"displayDate"`
}
type GetStatDailyListRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*model.StatDailyList `json:"rows"`
	Total  int                    `json:"total"`
}
type PostStatDailyReq struct {
	g.Meta      `path:"/live/daily" method:"post" tags:"每日统计" summary:"新增"`
	Anchor      string `p:"anchor"  v:"required"`
	DisplayName string `p:"displayName"  v:"required"`
	DisplayType int    `p:"displayType"  v:"required"`
	DisplayDate string `p:"displayDate"  v:"required"`
	Count       int    `p:"count"`
	Remark      string `p:"remark"`
}
type PostStatDailyRes struct {
	g.Meta `mime:"application/json"`
}
type GetStatDailyReq struct {
	g.Meta `path:"/live/daily/{id}" method:"get" tags:"每日统计" summary:"详情"`
	Id     int64 `p:"id"  v:"required"`
}
type GetStatDailyRes struct {
	g.Meta `mime:"application/json"`
	*entity.StatDaily
}
type PutStatDailyReq struct {
	g.Meta      `path:"/live/daily" method:"put" tags:"每日统计" summary:"修改"`
	Id          int64  `p:"id"  v:"required"`
	Anchor      string `p:"anchor"  v:"required"`
	DisplayName string `p:"displayName"  v:"required"`
	DisplayType int    `p:"displayType"  v:"required"`
	DisplayDate string `p:"displayDate"  v:"required"`
	Count       int    `p:"count" v:"required"`
	Remark      string `p:"remark"`
}
type PutStatDailyRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteStatDailyReq struct {
	g.Meta `path:"/live/daily/{id}" method:"delete" tags:"每日统计" summary:"删除"`
	Id     string `p:"id"  v:"required"`
}
type DeleteStatDailyRes struct {
	g.Meta `mime:"application/json"`
}
