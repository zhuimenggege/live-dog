package v1

import (
	"github.com/shichen437/live-dog/api/v1/common"
	"github.com/shichen437/live-dog/internal/app/admin/model"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type SignUpReq struct {
	g.Meta    `path:"/user/sign-up" method:"post" tags:"用户管理" summary:"注册"`
	UserName  string `p:"userName" v:"required|length:6,16"`
	Password  string `p:"password" v:"required|length:6,16"`
	Password2 string `v:"required|length:6,16|same:Password"`
	Nickname  string
}

type SignUpRes struct{}

type SignInReq struct {
	g.Meta    `path:"/login" method:"post" tags:"用户管理" summary:"登录"`
	UserName  string `p:"userName" v:"required"`
	Password  string `p:"password" v:"required"`
	Code      string `p:"code" v:"required#验证码不能为空"`
	VerifyKey string `p:"verifyKey"`
}

type SignInRes struct {
	Token string `p:"token"`
}

type CheckUserNameReq struct {
	g.Meta   `path:"/user/check-UserName" method:"post" tags:"用户管理" summary:"用户名检查"`
	UserName string `v:"required"`
}
type CheckUserNameRes struct{}

type CheckNickNameReq struct {
	g.Meta   `path:"/user/check-nick-name" method:"post" tags:"用户管理" summary:"昵称检查"`
	Nickname string `v:"required"`
}
type CheckNickNameRes struct{}

type IsSignedInReq struct {
	g.Meta `path:"/user/is-signed-in" method:"post" tags:"用户管理" summary:"Check current user is already signed-in"`
}
type IsSignedInRes struct {
	OK bool `dc:"True if current user is signed in; or else false"`
}

type GetInfoReq struct {
	g.Meta `path:"/getInfo" method:"get" tags:"用户管理" summary:"详情"`
}

type GetInfoRes struct {
	g.Meta      `mime:"application/json"`
	Permissions []string              `json:"permissions"`
	Roles       []string              `json:"roles"`
	User        *model.SysUserInfoRes `json:"user"`
}
type GetRoutersReq struct {
	g.Meta `path:"/getRouters" method:"get" tags:"用户管理" summary:"路由"`
}
type GetRoutersRes struct {
	g.Meta   `mime:"application/json"`
	MenuList []*model.UserMenuRes `json:"menuList"`
}
type GetAddUserReq struct {
	g.Meta `path:"/system/user" method:"get" tags:"用户管理" summary:"详情"`
	UserId int64 `p:"userId"`
}
type GetAddUserRes struct {
	g.Meta  `mime:"application/json"`
	Roles   []*entity.SysRole `json:"roles"`
	RoleIds []int64           `json:"roleIds"`
	User    *model.UserList   `json:"user"`
}
type PostAddUserReq struct {
	g.Meta      `path:"/system/user" method:"post" tags:"用户管理" summary:"添加"`
	UserName    string  `p:"userName"`
	NickName    string  `p:"nickName"`
	Password    string  `p:"password"`
	Phonenumber string  `p:"phonenumber"`
	Email       string  `p:"email"`
	Sex         string  `p:"sex"`
	Status      string  `p:"status"`
	Remark      string  `p:"remark"`
	RoleIds     []int64 `p:"roleIds"`
}
type PostAddUserRes struct {
	g.Meta `mime:"application/json"`
}
type DeleteUserReq struct {
	g.Meta `path:"/system/user/{userId}" method:"DELETE" tags:"用户管理" summary:"删除"`
	UserId string `p:"userId" v:"required"`
}
type DeleteUserRes struct {
	g.Meta `mime:"application/json"`
}
type ChangeUserStatusReq struct {
	g.Meta `path:"/system/user/changeStatus" method:"PUT" tags:"用户管理" summary:"状态修改"`
	UserId int64 `p:"userId" v:"required"`
	Status int64 `p:"status" v:"required"`
}
type ChangeUserStatusRes struct {
	g.Meta `mime:"application/json"`
}
type ResetPwdUserReq struct {
	g.Meta   `path:"/system/user/resetPwd" method:"PUT" tags:"用户管理" summary:"重置密码"`
	UserId   int64  `p:"userId" v:"required"`
	Password string `p:"password" v:"required"`
}
type ResetPwdUserRes struct {
	g.Meta `mime:"application/json"`
}

type PutUpdateUserReq struct {
	g.Meta      `path:"/system/user" method:"put" tags:"用户管理" summary:"修改"`
	UserId      int64   `p:"userId" v:"required"`
	NickName    string  `p:"nickName"`
	Phonenumber string  `p:"phonenumber"`
	Email       string  `p:"email"`
	Sex         string  `p:"sex"`
	Status      string  `p:"status"`
	Remark      string  `p:"remark"`
	RoleIds     []int64 `p:"roleIds"`
}
type PutUpdateUserRes struct {
	g.Meta `mime:"application/json"`
}

type GetUserListReq struct {
	g.Meta `path:"/system/user/list" method:"get" tags:"用户管理" summary:"用户列表"`
	common.PageReq
	UserName    string `p:"userName"`
	Phonenumber string `p:"phonenumber"`
	Status      string `p:"status"`
}
type GetUserListRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*model.UserList `json:"rows"`
	Total  int               `json:"total"`
}
type GetAuthRoleUserReq struct {
	g.Meta `path:"/system/user/authRole/{userId}" method:"get" tags:"用户管理" summary:"获取角色"`
	UserId int64 `p:"userId" v:"required"`
	common.PageReq
}
type GetAuthRoleUserRes struct {
	g.Meta `mime:"application/json"`
	Roles  []*entity.SysRole   `json:"roles"`
	User   *model.AuthRoleUser `json:"user"`
}
type PutAuthRoleUserReq struct {
	g.Meta  `path:"/system/user/authRole" method:"put" tags:"用户管理" summary:"修改角色"`
	UserId  int64  `p:"userId" v:"required"`
	RoleIds string `p:"roleIds" v:"required"`
}
type PutAuthRoleUserRes struct {
	g.Meta `mime:"application/json"`
}
