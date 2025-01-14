package sys_role_menu

import (
	"context"

	"github.com/shichen437/live-dog/internal/app/admin/dao"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"
	"github.com/shichen437/live-dog/internal/app/admin/service"
	commonConst "github.com/shichen437/live-dog/internal/app/common/consts"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterSysRoleMenu(New())
}

func New() *sSysRoleMenu {
	return &sSysRoleMenu{}
}

type sSysRoleMenu struct {
}

func (s *sSysRoleMenu) AddRoleMenus(ctx context.Context, tx gdb.TX, roleId int64, MenuIds []int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//删除旧角色菜单
		_, err = dao.SysRoleMenu.Ctx(ctx).TX(tx).Where(dao.SysRoleMenu.Columns().RoleId, roleId).Delete()
		utils.WriteErrLogT(ctx, err, commonConst.AddF)
		if len(MenuIds) == 0 {
			return
		}
		//添加角色菜单信息
		data := g.List{}
		for _, v := range MenuIds {
			data = append(data, g.Map{
				dao.SysRoleMenu.Columns().RoleId: roleId,
				dao.SysRoleMenu.Columns().MenuId: v,
			})
		}
		_, err = dao.SysRoleMenu.Ctx(ctx).TX(tx).Data(data).Insert()
		utils.WriteErrLogT(ctx, err, commonConst.AddF)
	})
	return
}

func (s *sSysRoleMenu) GetMenuIdsByRoleId(ctx context.Context, roleId int64) (menuIds []int64, err error) {
	var roleMenu []*entity.SysRoleMenu
	err = dao.SysRoleMenu.Ctx(ctx).Fields(dao.SysRoleMenu.Columns().MenuId).Where(dao.SysRoleMenu.Columns().RoleId, roleId).Scan(&roleMenu)
	utils.WriteErrLogT(ctx, err, commonConst.GetF)
	for _, v := range roleMenu {
		menuIds = append(menuIds, v.MenuId)
	}
	return
}
