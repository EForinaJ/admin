// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysCommentDao is the data access object for the table sys_comment.
type SysCommentDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns SysCommentColumns // columns contains all the column names of Table for convenient usage.
}

// SysCommentColumns defines and stores column names for the table sys_comment.
type SysCommentColumns struct {
	Id         string //
	ManageId   string //
	UserId     string //
	ProductId  string //
	Content    string //
	Rate       string //
	Images     string //
	IsTop      string //
	Status     string //
	Ip         string //
	CreateTime string //
	UpdateTime string //
}

// sysCommentColumns holds the columns for the table sys_comment.
var sysCommentColumns = SysCommentColumns{
	Id:         "id",
	ManageId:   "manage_id",
	UserId:     "user_id",
	ProductId:  "product_id",
	Content:    "content",
	Rate:       "rate",
	Images:     "images",
	IsTop:      "is_top",
	Status:     "status",
	Ip:         "ip",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewSysCommentDao creates and returns a new DAO object for table data access.
func NewSysCommentDao() *SysCommentDao {
	return &SysCommentDao{
		group:   "default",
		table:   "sys_comment",
		columns: sysCommentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysCommentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysCommentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysCommentDao) Columns() SysCommentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysCommentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysCommentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysCommentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
