package admin

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/admin"
	"github.com/shichen437/live-dog/internal/app/admin/consts"
	"github.com/shichen437/live-dog/internal/app/admin/model"
	"github.com/shichen437/live-dog/internal/app/admin/service"
	commonConst "github.com/shichen437/live-dog/internal/app/common/consts"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/util/gconv"
)

type sysUserController struct{}

var SysUser = sysUserController{}

// SignUp is the API for user sign up.
func (c *sysUserController) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	err = service.SysUser().Create(ctx, model.UserCreateInput{
		UserName: req.UserName,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	return
}

// IsSignedIn checks and returns whether the user is signed in.
func (c *sysUserController) IsSignedIn(ctx context.Context, req *v1.IsSignedInReq) (res *v1.IsSignedInRes, err error) {
	res = &v1.IsSignedInRes{
		OK: service.SysUser().IsSignedIn(ctx),
	}
	return
}

// CheckUserName checks and returns whether the user UserName is available.
func (c *sysUserController) CheckUserName(ctx context.Context, req *v1.CheckUserNameReq) (res *v1.CheckUserNameRes, err error) {
	available, err := service.SysUser().IsUserNameAvailable(ctx, req.UserName)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, utils.TError(ctx, consts.UserNameExists)
	}
	return
}

// CheckNickName checks and returns whether the user nickname is available.
func (c *sysUserController) CheckNickName(ctx context.Context, req *v1.CheckNickNameReq) (res *v1.CheckNickNameRes, err error) {
	available, err := service.SysUser().IsNicknameAvailable(ctx, req.Nickname, 0)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, utils.TError(ctx, consts.UserNicknameExists)
	}
	return
}

// Profile returns the user profile.
func (c *sysUserController) Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {
	res, err = service.UserProfile().GetProfile(ctx, req)
	return
}

func (c *sysUserController) UpdateProfile(ctx context.Context, req *v1.PutProfileReq) (res *v1.PutProfileRes, err error) {
	res, err = service.UserProfile().UpdateProfile(ctx, req)
	return
}

func (s *sysUserController) UpdateProfilePwd(ctx context.Context, req *v1.PutProfilePwdReq) (res *v1.PutProfilePwdRes, err error) {
	return service.UserProfile().UpdateProfilePwd(ctx, req)
}

func (c *sysUserController) UpdateAvatar(ctx context.Context, req *v1.PostProfileAvatarReq) (res *v1.PostProfileAvatarRes, err error) {
	res, err = service.UserProfile().UpdateProfileAvatar(ctx, req)
	return
}

// GetInfo returns the user info.
func (c *sysUserController) GetInfo(ctx context.Context, req *v1.GetInfoReq) (res *v1.GetInfoRes, err error) {
	var (
		user       = &model.SysUserInfoRes{}
		sysUserRes = &model.SysUserRes{}
	)
	userId := gconv.Int64(ctx.Value(commonConst.CtxAdminId))
	sysUserRes = service.SysUser().GetUserById(ctx, userId)
	user.SysUserRes = sysUserRes

	//获取用户角色
	roleIds, rolesList := getUserRoles(ctx, userId, user)
	//当前用户是否超级管理员
	if user.UserId == commonConst.ProAdminId {
		user.Admin = true
	}
	//获取权限
	permlist := &model.RolePerm{}
	permlist, err = service.SysMenu().GetPermByRoleids(ctx, roleIds)
	//设置每个角色权限
	for _, val := range user.Roles {
		val.Permissions = permlist.MapPerms[val.RoleId]
	}
	res = &v1.GetInfoRes{
		User:        user,
		Roles:       rolesList.Roles,
		Permissions: permlist.AllPerm,
	}
	return
}

// Profile returns the user profile.
func (c *sysUserController) GetRouters(ctx context.Context, req *v1.GetRoutersReq) (res *v1.GetRoutersRes, err error) {
	userId := gconv.Int64(ctx.Value(commonConst.CtxAdminId))
	//获取roleids
	roleIds := []int64{}
	rolesList, err := service.SysRole().GetRolesByUid(ctx, userId)
	if err != nil {
		return nil, err
	}
	for _, v := range rolesList.SysRole {
		roleIds = append(roleIds, v.RoleId)
	}
	//获取菜单
	menuList, err := service.SysMenu().GetRoutersByRoleids(ctx, roleIds)
	if err != nil {
		return nil, err
	}
	//menuTree
	res = &v1.GetRoutersRes{
		MenuList: menuList,
	}
	return
}

func (c *sysUserController) GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error) {
	res, err = service.SysUser().GetUserList(ctx, req)
	return
}

func (c *sysUserController) GetAddUser(ctx context.Context, req *v1.GetAddUserReq) (res *v1.GetAddUserRes, err error) {
	res = &v1.GetAddUserRes{}
	res.Roles, err = service.SysRole().GetNomalRole(ctx)
	if err != nil {
		return
	}
	if req.UserId > 0 {
		user := &model.UserList{}
		user.SysUser, err = service.SysUser().GetOneUserById(ctx, req.UserId)
		if err != nil {
			return nil, err
		}
		res.User = user
		res.RoleIds, _ = service.SysUserRole().GetRoleIdByUid(ctx, req.UserId)
	}
	return
}
func (c *sysUserController) Add(ctx context.Context, req *v1.PostAddUserReq) (res *v1.PostAddUserRes, err error) {
	res, err = service.SysUser().Add(ctx, req)
	return
}
func (c *sysUserController) Update(ctx context.Context, req *v1.PutUpdateUserReq) (res *v1.PutUpdateUserRes, err error) {
	res, err = service.SysUser().Update(ctx, req)
	return
}
func (c *sysUserController) Delete(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	res, err = service.SysUser().Delete(ctx, req)
	return
}
func (c *sysUserController) ChangeStatus(ctx context.Context, req *v1.ChangeUserStatusReq) (res *v1.ChangeUserStatusRes, err error) {
	res, err = service.SysUser().ChangeStatus(ctx, req)
	return
}
func (c *sysUserController) ResetPWD(ctx context.Context, req *v1.ResetPwdUserReq) (res *v1.ResetPwdUserRes, err error) {
	res, err = service.SysUser().ResetPWD(ctx, req)
	return
}
func (c *sysUserController) GetAuthRole(ctx context.Context, req *v1.GetAuthRoleUserReq) (res *v1.GetAuthRoleUserRes, err error) {
	res, err = service.SysUser().GetAuthRole(ctx, req)
	return
}
func (c *sysUserController) PutAuthRole(ctx context.Context, req *v1.PutAuthRoleUserReq) (res *v1.PutAuthRoleUserRes, err error) {
	res, err = service.SysUser().PutAuthRole(ctx, req)
	return
}

func getUserRoles(ctx context.Context, userId int64, user *model.SysUserInfoRes) ([]int64, *model.SysRolesRes) {
	roleIds := []int64{}
	rolesList, _ := service.SysRole().GetRolesByUid(ctx, userId)
	for _, v := range rolesList.SysRole {
		sysRoleRes := &model.SysRoleRes{}
		if v.RoleId == commonConst.ProAdminRoleId {
			sysRoleRes.Admin = true
		}
		sysRoleRes.SysRole = v
		user.Roles = append(user.Roles, sysRoleRes)
		roleIds = append(roleIds, v.RoleId)
	}
	return roleIds, rolesList
}
