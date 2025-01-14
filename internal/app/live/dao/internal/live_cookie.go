// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LiveCookieDao is the data access object for table live_cookie.
type LiveCookieDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns LiveCookieColumns // columns contains all the column names of Table for convenient usage.
}

// LiveCookieColumns defines and stores column names for table live_cookie.
type LiveCookieColumns struct {
	Id         string // ID
	Platform   string // 平台
	Cookie     string // cookie
	Remark     string // 备注
	CreateTime string // 创建时间
	ActionTime string // 更新时间
}

// liveCookieColumns holds the columns for table live_cookie.
var liveCookieColumns = LiveCookieColumns{
	Id:         "id",
	Platform:   "platform",
	Cookie:     "cookie",
	Remark:     "remark",
	CreateTime: "create_time",
	ActionTime: "action_time",
}

// NewLiveCookieDao creates and returns a new DAO object for table data access.
func NewLiveCookieDao() *LiveCookieDao {
	return &LiveCookieDao{
		group:   "default",
		table:   "live_cookie",
		columns: liveCookieColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LiveCookieDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LiveCookieDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LiveCookieDao) Columns() LiveCookieColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LiveCookieDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LiveCookieDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LiveCookieDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
