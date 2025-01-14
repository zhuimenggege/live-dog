// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StatDailyDao is the data access object for table stat_daily.
type StatDailyDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns StatDailyColumns // columns contains all the column names of Table for convenient usage.
}

// StatDailyColumns defines and stores column names for table stat_daily.
type StatDailyColumns struct {
	Id          string // 记录ID
	Anchor      string // 主播
	DisplayName string // 展示名称
	DisplayType string // 展示类型（1 歌曲 2吉他）
	DisplayDate string // 展示时间
	Count       string // 次数
	Remark      string // 备注
	CreateBy    string // 创建者
	CreateTime  string // 创建时间
	UpdateBy    string // 更新者
	UpdateTime  string // 更新时间
	Action      string // 标识：0 新增 1 修改 2 删除
}

// statDailyColumns holds the columns for table stat_daily.
var statDailyColumns = StatDailyColumns{
	Id:          "id",
	Anchor:      "anchor",
	DisplayName: "display_name",
	DisplayType: "display_type",
	DisplayDate: "display_date",
	Count:       "count",
	Remark:      "remark",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
	Action:      "action",
}

// NewStatDailyDao creates and returns a new DAO object for table data access.
func NewStatDailyDao() *StatDailyDao {
	return &StatDailyDao{
		group:   "default",
		table:   "stat_daily",
		columns: statDailyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *StatDailyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *StatDailyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *StatDailyDao) Columns() StatDailyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *StatDailyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *StatDailyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *StatDailyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
