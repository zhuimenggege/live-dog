package sys_user

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/admin"
	"github.com/shichen437/live-dog/internal/app/admin/consts"
	"github.com/shichen437/live-dog/internal/app/admin/dao"
	"github.com/shichen437/live-dog/internal/app/admin/model"
	"github.com/shichen437/live-dog/internal/app/admin/model/do"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"
	"github.com/shichen437/live-dog/internal/app/admin/service"
	commonConst "github.com/shichen437/live-dog/internal/app/common/consts"
	commonService "github.com/shichen437/live-dog/internal/app/common/service"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type (
	sSysUser struct{}
)

func init() {
	service.RegisterSysUser(New())
}

func New() service.ISysUser {
	return &sSysUser{}
}

// Create creates user account.
func (s *sSysUser) Create(ctx context.Context, in model.UserCreateInput) (err error) {

	// If Nickname is not specified, it then uses UserName as its default Nickname.
	if in.Nickname == "" {
		in.Nickname = in.UserName
	}
	var (
		available bool
	)
	// UserName checks.
	available, err = s.IsUserNameAvailable(ctx, in.UserName)
	if err != nil {
		return err
	}
	if !available {
		return utils.TError(ctx, consts.UserNameExists)
	}
	// Nickname checks.
	available, err = s.IsNicknameAvailable(ctx, in.Nickname, 0)
	if err != nil {
		return err
	}
	if !available {
		return utils.TError(ctx, consts.UserNicknameExists)
	}
	return dao.SysUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.SysUser.Ctx(ctx).Data(do.SysUser{
			//UserName: in.UserName,
			Password: in.Password,
			//Nickname: in.Nickname,
		}).Insert()
		return err
	})
}

func (s *sSysUser) Add(ctx context.Context, req *v1.PostAddUserReq) (res *v1.PostAddUserRes, err error) {
	var (
		available bool
	)
	if req.NickName == "" {
		err = utils.TError(ctx, consts.UserNicknameEmpty)
		return
	}
	available, err = s.IsNicknameAvailable(ctx, req.NickName, 0)
	if err != nil {
		return
	}
	if !available {
		err = utils.TError(ctx, consts.UserNicknameExists)
		return
	}

	if req.UserName == "" {
		err = utils.TError(ctx, consts.UserNameEmpty)
		return
	}
	available, err = s.IsUserNameAvailable(ctx, req.UserName)
	if err != nil {
		return
	}
	if !available {
		err = utils.TError(ctx, consts.UserNameExists)
		return
	}
	if req.Password == "" || (len(req.Password) < 5 && len(req.Password) > 20) {
		err = utils.TError(ctx, consts.UserPasswordRules)
		return
	}
	newP, err := utils.Encrypt(ctx, req.Password)
	if err != nil {
		return
	}
	adminName := gconv.String(ctx.Value(commonConst.CtxAdminName))
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//添加用户信息
			userId, e := dao.SysUser.Ctx(ctx).TX(tx).InsertAndGetId(do.SysUser{
				UserName:    req.UserName,
				Password:    newP,
				NickName:    req.NickName,
				Phonenumber: req.Phonenumber,
				Email:       req.Email,
				Sex:         req.Sex,
				Status:      req.Status,
				Remark:      req.Remark,
				LoginIp:     utils.GetClientIp(ctx),
				CreateTime:  gtime.Now(),
				CreateBy:    adminName,
				LoginDate:   gtime.Now(),
				UpdateTime:  gtime.Now(),
				UpdateBy:    adminName,
			})
			utils.WriteErrLogT(ctx, e, commonConst.AddF)
			//设置用户角色
			err = service.SysUserRole().AddUserRoles(ctx, tx, userId, req.RoleIds)
			utils.WriteErrLogT(ctx, e, consts.UserRoleSetF)
		})
		return err
	})
	return
}

func (s *sSysUser) Update(ctx context.Context, req *v1.PutUpdateUserReq) (res *v1.PutUpdateUserRes, err error) {
	var (
		available bool
	)
	if req.NickName == "" {
		err = utils.TError(ctx, consts.UserNicknameEmpty)
		return
	}
	available, err = s.IsNicknameAvailable(ctx, req.NickName, req.UserId)
	if err != nil {
		return
	}
	if !available {
		err = utils.TError(ctx, consts.UserNicknameExists)
	}
	adminName := gconv.String(ctx.Value(commonConst.CtxAdminName))
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//添加用户信息
			_, e := dao.SysUser.Ctx(ctx).TX(tx).WherePri(req.UserId).Update(do.SysUser{
				NickName:    req.NickName,
				Phonenumber: req.Phonenumber,
				Email:       req.Email,
				Sex:         req.Sex,
				Status:      req.Status,
				Remark:      req.Remark,
				UpdateTime:  gtime.Now(),
				UpdateBy:    adminName,
			})
			utils.WriteErrLogT(ctx, e, commonConst.UpdateF)
			//设置用户角色
			err = service.SysUserRole().AddUserRoles(ctx, tx, req.UserId, req.RoleIds)
			utils.WriteErrLogT(ctx, e, consts.UserRoleSetF)
		})
		return err
	})
	return
}

// 假删除 支持批量删除
func (s *sSysUser) Delete(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	adminName := gconv.String(ctx.Value(commonConst.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		userIds := utils.ParamStrToSlice(req.UserId, ",")
		//删除用户信息
		_, e := dao.SysUser.Ctx(ctx).WhereIn(dao.SysUser.Columns().UserId, userIds).Update(do.SysUser{
			Status:     commonConst.SysUserStatusNo,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utils.WriteErrLogT(ctx, e, commonConst.DeleteF)
	})
	return
}

// 更改状态
func (s *sSysUser) ChangeStatus(ctx context.Context, req *v1.ChangeUserStatusReq) (res *v1.ChangeUserStatusRes, err error) {
	adminName := gconv.String(ctx.Value(commonConst.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		//删除用户信息
		_, e := dao.SysUser.Ctx(ctx).WherePri(req.UserId).Update(do.SysUser{
			Status:     req.Status,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utils.WriteErrLogT(ctx, e, commonConst.UpdateF)
	})
	return
}

// 更改密码
func (s *sSysUser) ResetPWD(ctx context.Context, req *v1.ResetPwdUserReq) (res *v1.ResetPwdUserRes, err error) {
	adminName := gconv.String(ctx.Value(commonConst.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		//删除用户信息
		password, _ := utils.Encrypt(ctx, req.Password)
		_, e := dao.SysUser.Ctx(ctx).WherePri(req.UserId).Update(do.SysUser{
			Password:   password,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utils.WriteErrLogT(ctx, e, consts.UserPasswordResetF)
	})
	return
}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *sSysUser) IsSignedIn(ctx context.Context) bool {
	if v := commonService.BizCtx().Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// IsUserNameAvailable checks and returns given UserName is available for signing up.
func (s *sSysUser) IsUserNameAvailable(ctx context.Context, UserName string) (bool, error) {
	count, err := dao.SysUser.Ctx(ctx).Where(do.SysUser{
		UserName: UserName,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsNicknameAvailable checks and returns given nickname is available for signing up.
func (s *sSysUser) IsNicknameAvailable(ctx context.Context, nickname string, userId int64) (bool, error) {
	sysUser := entity.SysUser{}
	if userId != 0 {
		m := dao.SysUser.Ctx(ctx).Where(do.SysUser{
			NickName: nickname,
		})
		total, err := m.Count()
		if err != nil {
			return false, err
		}
		if total > 0 {
			m.Scan(&sysUser)
			if userId > 0 && userId != sysUser.UserId {
				return false, nil
			}
		}
	} else {
		count, err := dao.SysUser.Ctx(ctx).Where(do.SysUser{
			NickName: nickname,
		}).Count()
		if err != nil {
			return false, err
		}
		return count == 0, nil
	}

	return true, nil
}

func (s *sSysUser) GetUserById(ctx context.Context, userId int64) (user *model.SysUserRes) {
	var err error
	err = g.Try(ctx, func(ctx context.Context) {
		//用户用户信息
		err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserId, userId).Scan(&user)
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
	})
	return
}
func (s *sSysUser) GetOneUserById(ctx context.Context, id int64) (user *entity.SysUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//用户用户信息
		err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserId, id).Scan(&user)
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
	})
	return
}

func (s *sSysUser) GetUserList(ctx context.Context, req *v1.GetUserListReq) (userList *v1.GetUserListRes, err error) {
	userList = &v1.GetUserListRes{}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = commonConst.PageSize
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//user := []*model.UserList{}
		user := []*entity.SysUser{}
		//用户用户信息
		m := dao.SysUser.Ctx(ctx)
		//手机号
		if req.Phonenumber != "" {
			m = m.Where(dao.SysUser.Columns().Phonenumber, req.Phonenumber)
		}
		//状态
		if req.Status != "" {
			m = m.Where(dao.SysUser.Columns().Status, req.Status)
		}
		//时间
		if len(req.Params) > 0 {
			m = m.WhereBetween(dao.SysUser.Columns().CreateTime, req.Params["beginTime"], req.Params["endTime"])
		}
		//用户名不为空
		if req.UserName != "" {
			m = m.WhereLike(dao.SysUser.Columns().UserName, "%"+req.UserName+"%")
		}
		userList.Total, err = m.Count()
		err = m.Page(req.PageNum, req.PageSize).Scan(&user)
		if err != nil {
			return
		}
		userRows := make([]*model.UserList, len(user))
		for k, value := range user {
			ul := &model.UserList{}
			ul.SysUser = value
			userRows[k] = ul
		}
		userList.Rows = userRows
		utils.WriteErrLogT(ctx, err, commonConst.ListF)
	})

	return
}
func (s *sSysUser) GetAuthRole(ctx context.Context, req *v1.GetAuthRoleUserReq) (res *v1.GetAuthRoleUserRes, err error) {
	res = &v1.GetAuthRoleUserRes{}
	var user *entity.SysUser
	var roleList []*entity.SysRole
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = commonConst.PageSize
	}
	user, err = s.GetOneUserById(ctx, req.UserId)
	if err != nil {
		return
	}
	authRoleUser := &model.AuthRoleUser{}
	authRoleUser.SysUser = user
	if req.UserId == commonConst.ProAdminId {
		authRoleUser.Admin = true
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//用户用户信息
		err = dao.SysRole.Ctx(ctx).As("r").Fields("r.*").LeftJoin(dao.SysUserRole.Table()+" ur", "r.role_id=ur.role_id").Where("ur.user_id", req.UserId).Page(req.PageNum, req.PageSize).Scan(&roleList)
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
	})
	res.Roles = roleList
	res.User = authRoleUser
	return
}
func (s *sSysUser) PutAuthRole(ctx context.Context, req *v1.PutAuthRoleUserReq) (res *v1.PutAuthRoleUserRes, err error) {
	roleIds := utils.ParamStrToSlice(req.RoleIds, ",")
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//设置角色
			err = service.SysUserRole().AddUserRoles(ctx, tx, req.UserId, roleIds)
			utils.WriteErrLogT(ctx, err, consts.UserRoleSetF)
		})
		return err
	})
	return
}
