// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PushChannelWebDao is the data access object for table push_channel_web.
type PushChannelWebDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns PushChannelWebColumns // columns contains all the column names of Table for convenient usage.
}

// PushChannelWebColumns defines and stores column names for table push_channel_web.
type PushChannelWebColumns struct {
	Id           string // 记录 ID
	ChannelId    string // 渠道 ID
	Url          string // 推送 URL
	HttpMethod   string // 请求方式
	Secret       string // 密钥/token/key
	AppId        string // 应用 ID
	CorpId       string // 企业 ID
	ReceiverId   string // 接收人 ID
	ReceiverType string // 接收人类型
	ExtraParams  string // 额外参数
	CreateTime   string // 创建时间
	ActionTime   string // 修改时间
}

// pushChannelWebColumns holds the columns for table push_channel_web.
var pushChannelWebColumns = PushChannelWebColumns{
	Id:           "id",
	ChannelId:    "channel_id",
	Url:          "url",
	HttpMethod:   "http_method",
	Secret:       "secret",
	AppId:        "app_id",
	CorpId:       "corp_id",
	ReceiverId:   "receiver_id",
	ReceiverType: "receiver_type",
	ExtraParams:  "extra_params",
	CreateTime:   "create_time",
	ActionTime:   "action_time",
}

// NewPushChannelWebDao creates and returns a new DAO object for table data access.
func NewPushChannelWebDao() *PushChannelWebDao {
	return &PushChannelWebDao{
		group:   "default",
		table:   "push_channel_web",
		columns: pushChannelWebColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PushChannelWebDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PushChannelWebDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PushChannelWebDao) Columns() PushChannelWebColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PushChannelWebDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PushChannelWebDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PushChannelWebDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
