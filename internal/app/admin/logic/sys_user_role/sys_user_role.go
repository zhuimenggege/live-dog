package sys_user_role

import (
	"context"

	"github.com/shichen437/live-dog/internal/app/admin/consts"
	"github.com/shichen437/live-dog/internal/app/admin/dao"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"
	"github.com/shichen437/live-dog/internal/app/admin/service"
	commonConst "github.com/shichen437/live-dog/internal/app/common/consts"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterSysUserRole(New())
}

func New() *sSysUserRole {
	return &sSysUserRole{}
}

type sSysUserRole struct {
}

func (s *sSysUserRole) AddUserRoles(ctx context.Context, tx gdb.TX, userId int64, roleIds []int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//删除旧用户角色
		_, err = dao.SysUserRole.Ctx(ctx).TX(tx).Where(dao.SysUserRole.Columns().UserId, userId).Delete()
		utils.WriteErrLogT(ctx, err, commonConst.DeleteF)
		if len(roleIds) == 0 {
			return
		}
		//添加用户角色信息
		data := g.List{}
		for _, v := range roleIds {
			data = append(data, g.Map{
				dao.SysUserRole.Columns().UserId: userId,
				dao.SysUserRole.Columns().RoleId: v,
			})
		}
		_, err = dao.SysUserRole.Ctx(ctx).TX(tx).Data(data).Insert()
		utils.WriteErrLogT(ctx, err, commonConst.AddF)
	})
	return
}

// 角色分配用户
func (s *sSysUserRole) AddRoleUsers(ctx context.Context, tx gdb.TX, roleId int64, userIds []int64) (err error) {

	err = g.Try(ctx, func(ctx context.Context) {
		//添加用户角色信息
		data := g.List{}
		for _, v := range userIds {
			count, err := dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns().UserId, v).Where(dao.SysUserRole.Columns().RoleId, roleId).Count()
			if err != nil {
				return
			}
			if count > 0 {
				continue
			}
			data = append(data, g.Map{
				dao.SysUserRole.Columns().RoleId: roleId,
				dao.SysUserRole.Columns().UserId: v,
			})
		}
		_, err = dao.SysUserRole.Ctx(ctx).TX(tx).Data(data).Insert()
		utils.WriteErrLogT(ctx, err, consts.UserRoleSetF)
	})
	return
}

// 取消角色分配用户
func (s *sSysUserRole) CancelRoleUsers(ctx context.Context, tx gdb.TX, roleId int64, userIds []int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//取消角色用户信息
		for _, v := range userIds {
			_, err = dao.SysUserRole.Ctx(ctx).TX(tx).Where(dao.SysUserRole.Columns().RoleId, roleId).Where(dao.SysUserRole.Columns().UserId, v).Delete()
		}
		utils.WriteErrLogT(ctx, err, consts.UserRoleCancelF)
	})
	return
}

func (s *sSysUserRole) GetRoleIdByUid(ctx context.Context, uid int64) (roleId []int64, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var user []*entity.SysUserRole
		//用户用户角色信息
		err = dao.SysUserRole.Ctx(ctx).Fields("role_id").Where(dao.SysUserRole.Columns().UserId, uid).Scan(&user)
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
		for _, v := range user {
			roleId = append(roleId, v.RoleId)
		}
	})
	return
}
func (s *sSysUserRole) GetUserIdByRoleId(ctx context.Context, roleId int64) (userId []int64, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var user []*entity.SysUserRole
		//用户用户角色信息
		err = dao.SysUserRole.Ctx(ctx).Fields(dao.SysUserRole.Columns().UserId).Where(dao.SysUserRole.Columns().RoleId, roleId).Scan(&user)
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
		for _, v := range user {
			userId = append(userId, v.UserId)
		}
	})
	return
}
