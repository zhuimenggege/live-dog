package sys_menu

import (
	"context"
	"strconv"
	"strings"
	"time"

	v1 "github.com/shichen437/live-dog/api/v1/admin"
	"github.com/shichen437/live-dog/internal/app/admin/dao"
	"github.com/shichen437/live-dog/internal/app/admin/model"
	"github.com/shichen437/live-dog/internal/app/admin/model/do"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"
	"github.com/shichen437/live-dog/internal/app/admin/service"
	commonConst "github.com/shichen437/live-dog/internal/app/common/consts"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/exp/maps"
)

func init() {
	service.RegisterSysMenu(New())
}

func New() *sSysMenu {
	return &sSysMenu{}
}

type sSysMenu struct {
}

//特殊路由

var SpecialApiPath = map[string]bool{
	//user
	"get/getInfo":                       true,
	"get/getRouters":                    true,
	"get/system/user/authRole/{userId}": true,
	"get/system/user/profile":           true,
	"post/logout":                       true,
	"post/system/user/profile/avatar":   true,
	"post/user/sign-up":                 true,
	"put/system/user/profile":           true,
	"put/system/user/profile/updatePwd": true,
	//系统工具
	"get/system/menu/roleMenuTreeselect/{roleId}": true,
	//菜单
	"get/system/menu/treeselect": true,
	//字典
	"get/system/dict/type":              true,
	"get/system/dict/type/optionselect": true,
}

// 根据role_id获取perms
func (s *sSysMenu) GetPermByRoleids(ctx context.Context, roleIds []int64) (permList *model.RolePerm, err error) {
	var mapPerms = make(map[int64][]string, 0)
	permList = new(model.RolePerm)
	if utils.InSliceInt64(commonConst.ProAdminRoleId, &roleIds) { //如果包含超级管理员
		permList.AllPerm = []string{"*:*:*"}
		mapPerms[commonConst.ProAdminRoleId] = []string{}
		permList.MapPerms = mapPerms
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		perms := []model.PermsData{}
		//权限
		err = dao.SysRoleMenu.Ctx(ctx).As("rm").Fields("rm.role_id,m.perms").LeftJoin(dao.SysMenu.Table()+" m", "rm.menu_id=m.menu_id").WhereIn("rm.role_id", roleIds).Scan(&perms)
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
		for _, v := range perms {
			if v.Perms == "" {
				continue
			}
			mapPerms[v.RoleId] = append(mapPerms[v.RoleId], v.Perms)
			if !utils.InSliceString(v.Perms, &permList.AllPerm) {
				permList.AllPerm = append(permList.AllPerm, v.Perms)
			}
		}
		permList.MapPerms = mapPerms

	})
	return
}

// 根据role_id获取menu
func (s *sSysMenu) GetMenuByRoleids(ctx context.Context, roleIds []int64) (menu []*entity.SysMenu, err error) {
	menuType := []string{"M", "C"}
	if utils.InSliceInt64(commonConst.ProAdminRoleId, &roleIds) { //包含超级管理员角色，返回所有的菜单
		err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Status, commonConst.SysMenuStatusOk).WhereIn(dao.SysMenu.Columns().MenuType, menuType).OrderAsc("order_num").Scan(&menu)
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//权限
		err = dao.SysRoleMenu.Ctx(ctx).As("rm").Fields("m.*").LeftJoin(dao.SysMenu.Table()+" m", "rm.menu_id=m.menu_id").WhereIn("rm.role_id", roleIds).WhereIn("m.menu_type", menuType).OrderAsc("m.order_num").Scan(&menu)
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
	})
	return
}
func (s *sSysMenu) GetMenuTreeByRoleids(ctx context.Context, roleIds []int64) (menu []*entity.SysMenu, err error) {
	if utils.InSliceInt64(commonConst.ProAdminRoleId, &roleIds) { //包含超级管理员角色，返回所有的菜单
		err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Status, commonConst.SysMenuStatusOk).OrderAsc("order_num").Scan(&menu)
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//权限
		err = dao.SysRoleMenu.Ctx(ctx).As("rm").Fields("m.*").LeftJoin(dao.SysMenu.Table()+" m", "rm.menu_id=m.menu_id").WhereIn("rm.role_id", roleIds).Where("m.status", commonConst.SysMenuStatusOk).OrderAsc("m.order_num").Scan(&menu)
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
	})
	return
}

func (s *sSysMenu) GetRoutersByRoleids(ctx context.Context, roleIds []int64) (userMenuRes []*model.UserMenuRes, err error) {
	menuList, err := s.GetMenuByRoleids(ctx, roleIds)
	if err != nil {
		return nil, err
	}
	userMenuRes = s.GetMenuTree(menuList, 0)
	return
}
func (s *sSysMenu) GetMenuTree(menuList []*entity.SysMenu, pid int64) (userMenuRes []*model.UserMenuRes) {
	tree := make([]*model.UserMenuRes, 0)
	for _, v := range menuList {
		if v.ParentId == pid {
			m := &model.UserMenuRes{}
			meta := &model.MenuMeta{}
			meta.Icon = v.Icon
			if v.IsFrame == 0 {
				meta.Link = v.Perms
			}
			if v.IsCache == 1 {
				meta.NoCache = true
			}
			meta.Title = v.MenuName
			m.Meta = meta
			m.Component = v.Component

			//m.AlwaysShow
			if v.Visible == "0" {
				m.Hidden = false
			}
			m.Path = v.Path
			if v.MenuType == "M" {
				m.Redirect = "noRedirect"
				m.Component = "Layout"
				m.Path = "/" + m.Path
				//只有是目录才设置为true
				m.AlwaysShow = true
			}
			m.Name = strings.ReplaceAll(v.Path, "/", "")
			m.Name = utils.StrFirstToUpper(m.Name)
			child := s.GetMenuTree(menuList, v.MenuId)
			if child != nil {
				m.Children = child
			}
			tree = append(tree, m)
		}
	}
	return tree
}

// 检测路由权限
func (s *sSysMenu) CheckUrlPerms(r *ghttp.Request) bool {
	url := strings.ToLower(r.Router.Method) + r.Router.Uri
	//获取userid
	userId := gconv.Int64(r.Context().Value(commonConst.CtxAdminId))
	//如果超级管理员，全部放行
	if userId == commonConst.ProAdminId {
		return true
	}
	perms := s.GetPermsUrlByUserId(r.Context(), userId)
	if _, ok := perms[url]; !ok {
		return false
	}
	return true
}

func (s *sSysMenu) GetPermsUrlByUserId(ctx context.Context, userId int64) (perms map[string]bool) {
	gcache := gcache.New()
	iPerms, err := gcache.Get(ctx, commonConst.CacheKeyPermsUrl+strconv.FormatInt(userId, 10))
	if err != nil {
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
	}
	if iPerms != nil {
		iPerms.Struct(&perms)
		return
	}
	perms = make(map[string]bool, 10)
	roleIds := []int64{}
	rolesList, err := service.SysRole().GetRolesByUid(ctx, userId)
	if err != nil {
		return
	}
	for _, v := range rolesList.SysRole {
		roleIds = append(roleIds, v.RoleId)
	}
	//获取菜单
	menuList, _ := s.GetMenuTreeByRoleids(ctx, roleIds)
	for _, menu := range menuList {
		perms[menu.ApiPath] = true
	}
	maps.Copy(perms, SpecialApiPath)
	//放入缓存
	gcache.Set(ctx, commonConst.CacheKeyPermsUrl+strconv.FormatInt(userId, 10), perms, time.Minute*5)
	return
}

func (s *sSysMenu) GetMenuTreeSelect(ctx context.Context) (menuTree []*model.SysMenuTreeRes, err error) {
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
	menuList, err := s.GetMenuTreeByRoleids(ctx, roleIds)
	if err != nil {
		return nil, err
	}
	menuTree = s.MenuTreeSelect(menuList, 0)
	return
}

func (s *sSysMenu) MenuTreeSelect(menuList []*entity.SysMenu, pid int64) (tree []*model.SysMenuTreeRes) {
	tree = make([]*model.SysMenuTreeRes, 0)
	for _, v := range menuList {
		if v.ParentId == pid {
			dd := &model.SysMenuTreeRes{Id: v.MenuId, Label: v.MenuName}
			child := s.MenuTreeSelect(menuList, v.MenuId)
			if child != nil {
				dd.Children = child
			}
			tree = append(tree, dd)
		}
	}
	return
}

// 菜单列表
func (s *sSysMenu) GetMenuList(ctx context.Context, req *v1.GetMenuListReq) (menuList *v1.GetMenuListRes, err error) {
	var list []*entity.SysMenu
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysMenu.Ctx(ctx)
		if req.Status != "" {
			m = m.Where(dao.SysMenu.Columns().Status, req.Status)
		}
		if req.MenuName != "" {
			m = m.WhereLike(dao.SysMenu.Columns().MenuName, "%"+req.MenuName+"%")
		}
		err = m.OrderAsc(dao.SysMenu.Columns().OrderNum).Scan(&list)
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
	})
	menuArr := make([]*model.SysMenuList, len(list))
	for k, menu := range list {
		m := &model.SysMenuList{}
		m.SysMenu = menu
		menuArr[k] = m
	}
	menuList = &v1.GetMenuListRes{}
	menuList.List = menuArr
	return
}

// 获取单条数据
func (s sSysMenu) GetOneMenuById(ctx context.Context, menuId int64) (menu *entity.SysMenu, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//菜单信息
		err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().MenuId, menuId).Scan(&menu)
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
	})
	return
}

// 修改数据
func (s sSysMenu) UpdateMenu(ctx context.Context, req *v1.PutMenuUpdateReq) (res *v1.PutMenuUpdateRes, err error) {
	adminName := gconv.String(ctx.Value(commonConst.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysMenu.Ctx(ctx).WherePri(req.MenuId).Update(do.SysMenu{
			ParentId:   req.ParentId,
			MenuType:   req.MenuType,
			Icon:       req.Icon,
			MenuName:   req.MenuName,
			OrderNum:   req.OrderNum,
			Status:     req.Status,
			IsFrame:    req.IsFrame,
			Path:       req.Path,
			ApiPath:    req.ApiPath,
			Component:  req.Component,
			Query:      req.Query,
			IsCache:    req.IsCache,
			Perms:      req.Perms,
			Remark:     req.Remark,
			Visible:    req.Visible,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utils.WriteErrLogT(ctx, e, commonConst.UpdateF)
	})
	return
}

// 初始化apipath
func (s *sSysMenu) InitApiPath(ctx context.Context) (err error) {

	err = g.Try(ctx, func(ctx context.Context) {
		var menu []*entity.SysMenu
		err = dao.SysMenu.Ctx(ctx).Where("api_path=''").Scan(&menu)
		for _, m := range menu {
			if m.Perms == "" {
				continue
			}
			perms := strings.Split(m.Perms, ":")
			apiPath := "/" + perms[0] + "/" + perms[1]
			switch perms[2] {
			case "add":
				apiPath = "post" + apiPath
			case "list":
				apiPath = "get" + apiPath + "/list"
			case "query":
				apiPath = "get" + apiPath + "/{" + perms[1] + "Id}"
			case "edit":
				apiPath = "put" + apiPath
			case "remove":
				apiPath = "delete" + apiPath + "/{" + perms[1] + "Id}"
			default:
				continue
			}
			_, err = dao.SysMenu.Ctx(ctx).WherePri(&m.MenuId).Update(do.SysMenu{
				ApiPath: apiPath,
			})
		}
	})
	return
}

// 添加数据
func (s sSysMenu) Add(ctx context.Context, req *v1.PostMenuAddReq) (err error) {
	adminName := gconv.String(ctx.Value(commonConst.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysMenu.Ctx(ctx).Data(do.SysMenu{
			ParentId:   req.ParentId,
			MenuType:   req.MenuType,
			Icon:       req.Icon,
			MenuName:   req.MenuName,
			OrderNum:   req.OrderNum,
			Status:     req.Status,
			IsFrame:    req.IsFrame,
			Path:       req.Path,
			ApiPath:    req.ApiPath,
			Visible:    req.Visible,
			Component:  req.Component,
			Query:      req.Query,
			IsCache:    req.IsCache,
			Perms:      req.Perms,
			Remark:     req.Remark,
			UpdateTime: gtime.Now(),
			CreateTime: gtime.Now(),
			CreateBy:   adminName,
			UpdateBy:   adminName,
		}).Insert()
		utils.WriteErrLogT(ctx, e, commonConst.AddF)
	})
	return
}

// 删除数据
func (s sSysMenu) Delete(ctx context.Context, req *v1.MenuDeleteReq) (err error) {
	adminName := gconv.String(ctx.Value(commonConst.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysMenu.Ctx(ctx).WherePri(req.MenuId).Update(do.SysMenu{
			Status:     commonConst.SysMenuStatusNo,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utils.WriteErrLogT(ctx, e, commonConst.DeleteF)
	})
	return
}
