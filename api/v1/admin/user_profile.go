package v1

import (
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type ProfileReq struct {
	g.Meta `path:"/system/user/profile" method:"get" tags:"个人信息" summary:"个人信息"`
}

type ProfileRes struct {
	User  *entity.SysUser `json:"user"`
	Roles []string        `json:"roles"`
}

type PutProfileReq struct {
	g.Meta      `path:"/system/user/profile" method:"put" tags:"个人信息" summary:"用户信息修改"`
	NickName    string `p:"nickName" v:"required|length:2,20"`
	Email       string `p:"email" v:"required|email"`
	Phonenumber string `p:"phonenumber" v:"required|phone"`
	Sex         int    `p:"sex" v:"required|in:0,1"`
}

type PutProfileRes struct {
	g.Meta `mime:"application/json"`
}

type PutProfilePwdReq struct {
	g.Meta      `path:"/system/user/profile/updatePwd" method:"put" tags:"个人信息" summary:"修改密码"`
	OldPassword string `p:"oldPassword" v:"required"`
	NewPassword string `p:"newPassword" v:"required"`
}

type PutProfilePwdRes struct {
	g.Meta `mime:"application/json"`
}

type PostProfileAvatarReq struct {
	g.Meta `path:"/system/user/profile/avatar" method:"post" tags:"个人信息" summary:"用户头像修改"`
	Avatar *ghttp.UploadFile `p:"avatar" type:"file" v:"required"`
}

type PostProfileAvatarRes struct {
	g.Meta `mime:"application/json"`
	ImgUrl string `json:"imgUrl"`
}
