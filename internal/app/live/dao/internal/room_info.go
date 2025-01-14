// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomInfoDao is the data access object for table room_info.
type RoomInfoDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns RoomInfoColumns // columns contains all the column names of Table for convenient usage.
}

// RoomInfoColumns defines and stores column names for table room_info.
type RoomInfoColumns struct {
	Id         string // 房间信息 ID
	LiveId     string // 房间 ID
	RoomName   string // 房间名称
	Anchor     string // 主播
	Platform   string // 直播平台
	Status     string // 状态
	CreateBy   string // 创建者
	CreateTime string // 创建时间
	ActionTime string // 修改时间
}

// roomInfoColumns holds the columns for table room_info.
var roomInfoColumns = RoomInfoColumns{
	Id:         "id",
	LiveId:     "live_id",
	RoomName:   "room_name",
	Anchor:     "anchor",
	Platform:   "platform",
	Status:     "status",
	CreateBy:   "create_by",
	CreateTime: "create_time",
	ActionTime: "action_time",
}

// NewRoomInfoDao creates and returns a new DAO object for table data access.
func NewRoomInfoDao() *RoomInfoDao {
	return &RoomInfoDao{
		group:   "default",
		table:   "room_info",
		columns: roomInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomInfoDao) Columns() RoomInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
