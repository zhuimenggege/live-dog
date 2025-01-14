// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PushMessageLogDao is the data access object for table push_message_log.
type PushMessageLogDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns PushMessageLogColumns // columns contains all the column names of Table for convenient usage.
}

// PushMessageLogColumns defines and stores column names for table push_message_log.
type PushMessageLogColumns struct {
	Id         string // 主键 ID
	ChannelId  string // 渠道 ID
	Status     string // 0：失败 1 成功
	Message    string // 消息内容
	PushType   string // 推送类型
	CreateTime string // 推送时间
}

// pushMessageLogColumns holds the columns for table push_message_log.
var pushMessageLogColumns = PushMessageLogColumns{
	Id:         "id",
	ChannelId:  "channel_id",
	Status:     "status",
	Message:    "message",
	PushType:   "push_type",
	CreateTime: "create_time",
}

// NewPushMessageLogDao creates and returns a new DAO object for table data access.
func NewPushMessageLogDao() *PushMessageLogDao {
	return &PushMessageLogDao{
		group:   "default",
		table:   "push_message_log",
		columns: pushMessageLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PushMessageLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PushMessageLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PushMessageLogDao) Columns() PushMessageLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PushMessageLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PushMessageLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PushMessageLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
