// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "github.com/shichen437/live-dog/api/v1/admin"
	"github.com/shichen437/live-dog/internal/app/admin/model"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"
)

type (
	ISysUser interface {
		// Create creates user account.
		Create(ctx context.Context, in model.UserCreateInput) (err error)
		Add(ctx context.Context, req *v1.PostAddUserReq) (res *v1.PostAddUserRes, err error)
		Update(ctx context.Context, req *v1.PutUpdateUserReq) (res *v1.PutUpdateUserRes, err error)
		// 假删除 支持批量删除
		Delete(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error)
		// 更改状态
		ChangeStatus(ctx context.Context, req *v1.ChangeUserStatusReq) (res *v1.ChangeUserStatusRes, err error)
		// 更改密码
		ResetPWD(ctx context.Context, req *v1.ResetPwdUserReq) (res *v1.ResetPwdUserRes, err error)
		// IsSignedIn checks and returns whether current user is already signed-in.
		IsSignedIn(ctx context.Context) bool
		// IsUserNameAvailable checks and returns given UserName is available for signing up.
		IsUserNameAvailable(ctx context.Context, UserName string) (bool, error)
		// IsNicknameAvailable checks and returns given nickname is available for signing up.
		IsNicknameAvailable(ctx context.Context, nickname string, userId int64) (bool, error)
		GetUserById(ctx context.Context, userId int64) (user *model.SysUserRes)
		GetOneUserById(ctx context.Context, id int64) (user *entity.SysUser, err error)
		GetUserList(ctx context.Context, req *v1.GetUserListReq) (userList *v1.GetUserListRes, err error)
		GetAuthRole(ctx context.Context, req *v1.GetAuthRoleUserReq) (res *v1.GetAuthRoleUserRes, err error)
		PutAuthRole(ctx context.Context, req *v1.PutAuthRoleUserReq) (res *v1.PutAuthRoleUserRes, err error)
	}
	IUserProfile interface {
		GetProfile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error)
		UpdateProfile(ctx context.Context, req *v1.PutProfileReq) (res *v1.PutProfileRes, err error)
		UpdateProfilePwd(ctx context.Context, req *v1.PutProfilePwdReq) (res *v1.PutProfilePwdRes, err error)
		UpdateProfileAvatar(ctx context.Context, req *v1.PostProfileAvatarReq) (res *v1.PostProfileAvatarRes, err error)
	}
)

var (
	localSysUser     ISysUser
	localUserProfile IUserProfile
)

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for interface ISysUser, forgot register?")
	}
	return localSysUser
}

func RegisterSysUser(i ISysUser) {
	localSysUser = i
}

func UserProfile() IUserProfile {
	if localUserProfile == nil {
		panic("implement not found for interface IUserProfile, forgot register?")
	}
	return localUserProfile
}

func RegisterUserProfile(i IUserProfile) {
	localUserProfile = i
}
