package v1

import (
	"github.com/shichen437/live-dog/api/v1/common"
	"github.com/shichen437/live-dog/internal/app/live/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetCookieListReq struct {
	g.Meta `path:"/live/cookie/list" method:"get" tags:"Cookie管理" summary:"Cookie列表"`
	common.PageReq
	Platform string `p:"platform"`
}
type GetCookieListRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*entity.LiveCookie `json:"rows"`
	Total  int                  `json:"total"`
}

type PostLiveCookieReq struct {
	g.Meta   `path:"/live/cookie" method:"post" tags:"Cookie管理" summary:"添加Cookie"`
	Platform string `p:"platform"  v:"required"`
	Cookie   string `p:"cookie"  v:"required"`
	Remark   string `p:"remark"`
}
type PostLiveCookieRes struct {
	g.Meta `mime:"application/json"`
}

type PutLiveCookieReq struct {
	g.Meta   `path:"/live/cookie" method:"put" tags:"Cookie管理" summary:"修改Cookie"`
	Id       int    `p:"id"  v:"required"`
	Platform string `p:"platform"`
	Cookie   string `p:"cookie"  v:"required"`
	Remark   string `p:"remark"`
}
type PutLiveCookieRes struct {
	g.Meta `mime:"application/json"`
}

type GetLiveCookieReq struct {
	g.Meta `path:"/live/cookie/{id}" method:"get" tags:"Cookie管理" summary:"获取Cookie"`
	Id     int `p:"id"  v:"required"`
}
type GetLiveCookieRes struct {
	g.Meta `mime:"application/json"`
	*entity.LiveCookie
}

type DeleteLiveCookieReq struct {
	g.Meta `path:"/live/cookie/{id}" method:"delete" tags:"Cookie管理" summary:"删除Cookie"`
	Id     string `p:"id"  v:"required"`
}
type DeleteLiveCookieRes struct {
	g.Meta `mime:"application/json"`
}
