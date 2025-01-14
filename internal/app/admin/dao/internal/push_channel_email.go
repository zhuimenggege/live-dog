// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PushChannelEmailDao is the data access object for table push_channel_email.
type PushChannelEmailDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns PushChannelEmailColumns // columns contains all the column names of Table for convenient usage.
}

// PushChannelEmailColumns defines and stores column names for table push_channel_email.
type PushChannelEmailColumns struct {
	Id         string // 主键 ID
	ChannelId  string // 渠道 ID
	From       string // 发送人
	To         string // 接收人
	Server     string // 发送服务器地址
	Port       string // 发送端口
	AuthCode   string // 授权码
	CreateTime string // 创建时间
	ActionTime string // 修改时间
}

// pushChannelEmailColumns holds the columns for table push_channel_email.
var pushChannelEmailColumns = PushChannelEmailColumns{
	Id:         "id",
	ChannelId:  "channel_id",
	From:       "from",
	To:         "to",
	Server:     "server",
	Port:       "port",
	AuthCode:   "auth_code",
	CreateTime: "create_time",
	ActionTime: "action_time",
}

// NewPushChannelEmailDao creates and returns a new DAO object for table data access.
func NewPushChannelEmailDao() *PushChannelEmailDao {
	return &PushChannelEmailDao{
		group:   "default",
		table:   "push_channel_email",
		columns: pushChannelEmailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PushChannelEmailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PushChannelEmailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PushChannelEmailDao) Columns() PushChannelEmailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PushChannelEmailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PushChannelEmailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PushChannelEmailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
