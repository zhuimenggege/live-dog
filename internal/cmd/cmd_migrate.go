package cmd

import (
	"context"
	"strings"

	"github.com/shichen437/live-dog/internal/app/admin/dao"
	"github.com/shichen437/live-dog/internal/app/admin/model/do"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate() (err error) {
	ctx := gctx.GetInitCtx()
	g.Log().Info(ctx, "database migrate start!")
	link := utils.DbLink
	if link == "" {
		g.Log().Error(ctx, "获取数据库链接配置失败", err)
		err = gerror.New("获取数据库链接配置失败")
		return
	}
	mysqlStr := strings.Replace(link, "mysql:", "mysql://", 1)
	m, err := migrate.New("file://manifest/migrate", mysqlStr)
	if err != nil {
		g.Log().Error(ctx, "创建数据库失败", err)
		return
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		g.Log().Error(ctx, "数据库迁移失败", err)
		return
	}
	err = initAdminInfo(ctx)
	if err != nil {
		g.Log().Error(ctx, "初始化管理员信息失败")
		return
	}
	g.Log().Info(ctx, "database migrate done!")
	return
}

func initAdminInfo(ctx context.Context) (err error) {
	var users *entity.SysUser
	err = dao.SysUser.Ctx(ctx).WherePri(1).Scan(&users)
	if err != nil {
		g.Log().Error(ctx, "获取管理员账户信息失败")
		return
	}
	enc, _ := utils.Encrypt(ctx, "admin123")
	if users == nil {
		_, err = dao.SysUser.Ctx(ctx).Insert(do.SysUser{
			UserId:     1,
			UserName:   "admin",
			NickName:   "live-dog",
			Password:   enc,
			UserType:   "00",
			Sex:        1,
			CreateBy:   "admin",
			CreateTime: gtime.Now(),
		})
	}
	return
}
