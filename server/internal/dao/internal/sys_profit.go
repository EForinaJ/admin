// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProfitDao is the data access object for the table sys_profit.
type SysProfitDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns SysProfitColumns // columns contains all the column names of Table for convenient usage.
}

// SysProfitColumns defines and stores column names for the table sys_profit.
type SysProfitColumns struct {
	Id         string //
	Related    string //
	Type       string //
	Amount     string //
	CreateTime string //
}

// sysProfitColumns holds the columns for the table sys_profit.
var sysProfitColumns = SysProfitColumns{
	Id:         "id",
	Related:    "related",
	Type:       "type",
	Amount:     "amount",
	CreateTime: "create_time",
}

// NewSysProfitDao creates and returns a new DAO object for table data access.
func NewSysProfitDao() *SysProfitDao {
	return &SysProfitDao{
		group:   "default",
		table:   "sys_profit",
		columns: sysProfitColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysProfitDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysProfitDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysProfitDao) Columns() SysProfitColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysProfitDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysProfitDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysProfitDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
