package sys_user

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/admin"
	"github.com/shichen437/live-dog/internal/app/admin/consts"
	"github.com/shichen437/live-dog/internal/app/admin/dao"
	"github.com/shichen437/live-dog/internal/app/admin/model/do"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"
	"github.com/shichen437/live-dog/internal/app/admin/service"
	commonConst "github.com/shichen437/live-dog/internal/app/common/consts"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type (
	sUserProfile struct{}
)

func init() {
	service.RegisterUserProfile(NewProfile())
}

func NewProfile() service.IUserProfile {
	return &sUserProfile{}
}

func (s *sUserProfile) GetProfile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {
	userId := gconv.Int64(ctx.Value(commonConst.CtxAdminId))
	res = &v1.ProfileRes{}
	if userId == 0 {
		err = utils.TError(ctx, commonConst.IDEmpty)
		return
	}
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserId, userId).Scan(&res.User)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	rolesList, err := service.SysRole().GetRolesByUid(ctx, userId)
	utils.WriteErrLogT(ctx, err, commonConst.GetF)
	if rolesList != nil && len(rolesList.RoleNames) > 0 {
		res.Roles = rolesList.RoleNames
	}
	return
}

func (s *sUserProfile) UpdateProfile(ctx context.Context, req *v1.PutProfileReq) (res *v1.PutProfileRes, err error) {
	userId := gconv.Int64(ctx.Value(commonConst.CtxAdminId))
	adminName := gconv.String(ctx.Value(commonConst.CtxAdminName))
	if userId == 0 {
		err = utils.TError(ctx, commonConst.IDEmpty)
		return
	}
	available, err := service.SysUser().IsNicknameAvailable(ctx, req.NickName, userId)
	if err != nil {
		return
	}
	if !available {
		err = utils.TError(ctx, consts.UserNicknameExists)
		return
	}
	//添加用户信息
	_, err = dao.SysUser.Ctx(ctx).WherePri(userId).Update(do.SysUser{
		NickName:    req.NickName,
		Email:       req.Email,
		Phonenumber: req.Phonenumber,
		Sex:         req.Sex,
		UpdateTime:  gtime.Now(),
		UpdateBy:    adminName,
	})
	return
}

func (s *sUserProfile) UpdateProfilePwd(ctx context.Context, req *v1.PutProfilePwdReq) (res *v1.PutProfilePwdRes, err error) {
	userId := gconv.Int64(ctx.Value(commonConst.CtxAdminId))
	if userId == 0 {
		err = utils.TError(ctx, commonConst.IDEmpty)
		return
	}
	var user *entity.SysUser
	err = g.Try(ctx, func(ctx context.Context) {
		//用户用户信息
		err = dao.SysUser.Ctx(ctx).WherePri(userId).Scan(&user)
	})
	if user == nil || err != nil {
		err = utils.TError(ctx, commonConst.DataNotFound)
		return
	}
	oldP, _ := utils.Encrypt(ctx, req.OldPassword)
	if oldP != user.Password {
		err = utils.TError(ctx, consts.UserOldPasswordError)
		return
	}
	newP, _ := utils.Encrypt(ctx, req.NewPassword)
	_, err = dao.SysUser.Ctx(ctx).WherePri(userId).Update(do.SysUser{
		Password:   newP,
		UpdateTime: gtime.Now(),
	})
	utils.WriteErrLogT(ctx, err, commonConst.UpdateF)
	return
}

func (s *sUserProfile) UpdateProfileAvatar(ctx context.Context, req *v1.PostProfileAvatarReq) (res *v1.PostProfileAvatarRes, err error) {
	userId := gconv.Int64(ctx.Value(commonConst.CtxAdminId))
	res = &v1.PostProfileAvatarRes{}
	file := req.Avatar
	file.Filename = "a.jpg"
	g.Log().Info(ctx, file)
	name, err := file.Save(utils.GetFileUploadPath()+commonConst.AvatarPrefix, true)
	if err != nil {
		err = gerror.New("上传失败")
		return
	}
	res.ImgUrl = commonConst.AvatarPrefix + name
	var user *entity.SysUser
	err = g.Try(ctx, func(ctx context.Context) {
		//用户用户信息
		err = dao.SysUser.Ctx(ctx).WherePri(userId).Scan(&user)
	})
	if user == nil || err != nil {
		err = utils.TError(ctx, commonConst.DataNotFound)
		return
	}
	_, err = dao.SysUser.Ctx(ctx).WherePri(userId).Update(do.SysUser{
		Avatar:     res.ImgUrl,
		UpdateTime: gtime.Now(),
	})
	utils.WriteErrLogT(ctx, err, commonConst.UpdateF)
	return
}
