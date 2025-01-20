package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shichen437/live-dog/api/v1/common"
	"github.com/shichen437/live-dog/internal/app/live/model"
)

type GetLiveHistoryListReq struct {
	g.Meta `path:"/live/history/list" method:"get" tags:"直播历史" summary:"历史列表"`
	common.PageReq
	LiveId int    `p:"liveId"`
	Anchor string `p:"anchor"`
}
type GetLiveHistoryListRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*model.LiveHistory `json:"rows"`
	Total  int                  `json:"total"`
}

type PostLiveHistoryReq struct {
	g.Meta    `path:"/live/history" method:"post" tags:"直播历史" summary:"添加直播历史"`
	LiveId    int         `p:"liveId"  v:"required"`
	StartTime *gtime.Time `p:"startTime"  v:"required"`
	EndTime   *gtime.Time `p:"endTime"  v:"required"`
}
type PostLiveHistoryRes struct {
	g.Meta `mime:"application/json"`
}

type PutLiveHistoryReq struct {
	g.Meta    `path:"/live/history" method:"put" tags:"直播历史" summary:"修改直播历史"`
	Id        int         `p:"id"  v:"required"`
	StartTime *gtime.Time `p:"startTime"  v:"required"`
	EndTime   *gtime.Time `p:"endTime"  v:"required"`
}
type PutLiveHistoryRes struct {
	g.Meta `mime:"application/json"`
}

type GetLiveHistoryReq struct {
	g.Meta `path:"/live/history/{id}" method:"get" tags:"直播历史" summary:"获取直播历史"`
	Id     int `p:"id"  v:"required"`
}
type GetLiveHistoryRes struct {
	g.Meta `mime:"application/json"`
	*model.LiveHistory
}

type DeleteLiveHistoryReq struct {
	g.Meta `path:"/live/history/{id}" method:"delete" tags:"直播历史" summary:"删除直播历史"`
	Id     string `p:"id"  v:"required"`
}
type DeleteLiveHistoryRes struct {
	g.Meta `mime:"application/json"`
}
