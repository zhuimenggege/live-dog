package v1

import (
	"github.com/shichen437/live-dog/api/v1/common"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetSysNoticeListReq struct {
	g.Meta `path:"/system/notice/list" method:"get" tags:"通知" summary:"列表"`
	common.PageReq
	NoticeTitle string `p:"noticeTitle"`
	NoticeType  string `p:"noticeType"`
	Status      string `p:"status"`
}
type GetSysNoticeListRes struct {
	g.Meta `mime:"application/json"`
	Data   []*entity.SysNotice `json:"data"`
	Total  int                 `json:"total"`
}

type GetSysNoticeDetailReq struct {
	g.Meta   `path:"/system/notice/{noticeId}" method:"get" tags:"通知" summary:"详情"`
	NoticeId int `p:"noticeId"  v:"required"`
}
type GetSysNoticeDetailRes struct {
	g.Meta `mime:"application/json"`
	*entity.SysNotice
}

type PostSysNoticeReq struct {
	g.Meta        `path:"/system/notice" method:"post" tags:"通知" summary:"新增"`
	NoticeTitle   string `p:"noticeTitle" v:"required"`
	NoticeType    string `p:"noticeType" v:"required"`
	NoticeContent []byte `p:"noticeContent" v:"required"`
	Status        string `p:"status" v:"required"`
	Remark        string `p:"remark"`
}
type PostSysNoticeRes struct {
	g.Meta `mime:"application/json"`
}

type PutSysNoticeReq struct {
	g.Meta        `path:"/system/notice" method:"put" tags:"通知" summary:"修改"`
	NoticeId      int    `p:"NoticeId" v:"required"`
	NoticeTitle   int    `p:"NoticeTitle" v:"required"`
	NoticeType    string `p:"NoticeType" v:"required"`
	NoticeContent []byte `p:"NoticeContent" v:"required"`
	Status        string `p:"status" v:"required"`
	Remark        string `p:"remark"`
}
type PutSysNoticeRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteSysNoticeReq struct {
	g.Meta   `path:"/system/notice/{noticeId}" method:"delete" tags:"通知" summary:"删除"`
	NoticeId string `p:"noticeId"  v:"required"`
}
type DeleteSysNoticeRes struct {
	g.Meta `mime:"application/json"`
}
