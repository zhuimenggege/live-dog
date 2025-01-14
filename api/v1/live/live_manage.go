package v1

import (
	"github.com/shichen437/live-dog/api/v1/common"
	"github.com/shichen437/live-dog/internal/app/live/model"
	"github.com/shichen437/live-dog/internal/app/live/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetRoomInfoListReq struct {
	g.Meta `path:"/live/info/list" method:"get" tags:"直播管理" summary:"房间列表"`
	common.PageReq
	Anchor   string `p:"anchor"`
	RoomName string `p:"roomName"`
	Platform string `p:"platform"`
}
type GetRoomInfoListRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*model.RoomInfo `json:"rows"`
	Total  int               `json:"total"`
}

type PostLiveManageReq struct {
	g.Meta       `path:"/live/manage" method:"post" tags:"直播管理" summary:"添加房间"`
	RoomUrl      string `p:"roomUrl"  v:"required"`
	Interval     int    `p:"interval"`
	Format       string `p:"format"`
	EnableNotice int    `p:"enableNotice"`
	MonitorType  int    `p:"monitorType"`
	MonitorStart string `p:"monitorStart"`
	MonitorStop  string `p:"monitorStop"`
	Remark       string `p:"remark"`
}
type PostLiveManageRes struct {
	g.Meta `mime:"application/json"`
}

type PutLiveManageReq struct {
	g.Meta       `path:"/live/manage" method:"put" tags:"直播管理" summary:"修改房间"`
	Id           int    `p:"id"  v:"required"`
	Interval     int    `p:"interval"`
	Format       string `p:"format"`
	EnableNotice int    `p:"enableNotice"`
	MonitorType  int    `p:"monitorType"`
	MonitorStart string `p:"monitorStart"`
	MonitorStop  string `p:"monitorStop"`
	Remark       string `p:"remark"`
}
type PutLiveManageRes struct {
	g.Meta `mime:"application/json"`
}

type GetLiveManageReq struct {
	g.Meta `path:"/live/manage/{id}" method:"get" tags:"直播管理" summary:"房间管理"`
	Id     int `p:"id"  v:"required"`
}
type GetLiveManageRes struct {
	g.Meta `mime:"application/json"`
	*entity.LiveManage
}

type DeleteLiveManageReq struct {
	g.Meta `path:"/live/manage/{id}" method:"delete" tags:"直播管理" summary:"删除房间"`
	Id     int `p:"id"  v:"required"`
}
type DeleteLiveManageRes struct {
	g.Meta `mime:"application/json"`
}
