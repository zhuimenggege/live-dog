// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LiveHistoryDao is the data access object for table live_history.
type LiveHistoryDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns LiveHistoryColumns // columns contains all the column names of Table for convenient usage.
}

// LiveHistoryColumns defines and stores column names for table live_history.
type LiveHistoryColumns struct {
	Id        string //
	LiveId    string // 直播ID
	StartTime string // 直播开始时间
	EndTime   string // 直播结束时间
	Duration  string // 直播时长
}

// liveHistoryColumns holds the columns for table live_history.
var liveHistoryColumns = LiveHistoryColumns{
	Id:        "id",
	LiveId:    "live_id",
	StartTime: "start_time",
	EndTime:   "end_time",
	Duration:  "duration",
}

// NewLiveHistoryDao creates and returns a new DAO object for table data access.
func NewLiveHistoryDao() *LiveHistoryDao {
	return &LiveHistoryDao{
		group:   "default",
		table:   "live_history",
		columns: liveHistoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LiveHistoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LiveHistoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LiveHistoryDao) Columns() LiveHistoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LiveHistoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LiveHistoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LiveHistoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
