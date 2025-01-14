// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LiveManageDao is the data access object for table live_manage.
type LiveManageDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns LiveManageColumns // columns contains all the column names of Table for convenient usage.
}

// LiveManageColumns defines and stores column names for table live_manage.
type LiveManageColumns struct {
	Id           string // 房间 id
	RoomUrl      string // 房间 url
	Interval     string // 轮询间隔
	Format       string // 导出视频格式
	EnableNotice string // 启用通知
	MonitorType  string // 监控类型
	MonitorStart string // 监控开始时间
	MonitorStop  string // 监控结束时间
	Remark       string // 房间备注
	CreateBy     string // 创建人
	CreateTime   string // 创建时间
	ActionBy     string // 修改人
	ActionTime   string // 修改时间
}

// liveManageColumns holds the columns for table live_manage.
var liveManageColumns = LiveManageColumns{
	Id:           "id",
	RoomUrl:      "room_url",
	Interval:     "interval",
	Format:       "format",
	EnableNotice: "enable_notice",
	MonitorType:  "monitor_type",
	MonitorStart: "monitor_start",
	MonitorStop:  "monitor_stop",
	Remark:       "remark",
	CreateBy:     "create_by",
	CreateTime:   "create_time",
	ActionBy:     "action_by",
	ActionTime:   "action_time",
}

// NewLiveManageDao creates and returns a new DAO object for table data access.
func NewLiveManageDao() *LiveManageDao {
	return &LiveManageDao{
		group:   "default",
		table:   "live_manage",
		columns: liveManageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LiveManageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LiveManageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LiveManageDao) Columns() LiveManageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LiveManageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LiveManageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LiveManageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
