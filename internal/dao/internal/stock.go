// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StockDao is the data access object for table stock.
type StockDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns StockColumns // columns contains all the column names of Table for convenient usage.
}

// StockColumns defines and stores column names for table stock.
type StockColumns struct {
	Id         string // 主键ID
	DealerId   string // 经销商ID
	CarId      string // 车辆ID
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
}

// stockColumns holds the columns for table stock.
var stockColumns = StockColumns{
	Id:         "id",
	DealerId:   "dealer_id",
	CarId:      "car_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
}

// NewStockDao creates and returns a new DAO object for table data access.
func NewStockDao() *StockDao {
	return &StockDao{
		group:   "default",
		table:   "stock",
		columns: stockColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *StockDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *StockDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *StockDao) Columns() StockColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *StockDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *StockDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *StockDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
