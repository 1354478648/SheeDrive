// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderDao is the data access object for table order.
type OrderDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns OrderColumns // columns contains all the column names of Table for convenient usage.
}

// OrderColumns defines and stores column names for table order.
type OrderColumns struct {
	Id          string // 主键ID
	UserId      string // 用户ID
	DealerId    string // 经销商ID
	CarId       string // 车辆ID
	AddrId      string // 用户地址ID
	Status      string // 订单状态 -1:异常,0:取消,1:未确认,2:已确认,3:签署协议,4:试驾中,5:试驾结束,6:待评价,7:已评价
	ConfirmTime string // 确认时间
	SignTime    string // 签署协议时间
	StartTime   string // 试驾开始时间
	EndTime     string // 试驾结束时间
	CommentTime string // 评价时间
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
	DeleteTime  string // 删除时间
}

// orderColumns holds the columns for table order.
var orderColumns = OrderColumns{
	Id:          "id",
	UserId:      "user_id",
	DealerId:    "dealer_id",
	CarId:       "car_id",
	AddrId:      "addr_id",
	Status:      "status",
	ConfirmTime: "confirm_time",
	SignTime:    "sign_time",
	StartTime:   "start_time",
	EndTime:     "end_time",
	CommentTime: "comment_time",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
	DeleteTime:  "delete_time",
}

// NewOrderDao creates and returns a new DAO object for table data access.
func NewOrderDao() *OrderDao {
	return &OrderDao{
		group:   "default",
		table:   "order",
		columns: orderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrderDao) Columns() OrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
