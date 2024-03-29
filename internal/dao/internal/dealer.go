// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DealerDao is the data access object for table dealer.
type DealerDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns DealerColumns // columns contains all the column names of Table for convenient usage.
}

// DealerColumns defines and stores column names for table dealer.
type DealerColumns struct {
	Id           string // 主键ID
	Name         string // 名称
	Username     string // 用户名
	Password     string // 密码
	Avatar       string // 头像
	Phone        string // 手机号
	DescribeInfo string // 描述信息
	Status       string // 状态 0:禁用, 1:正常
	Token        string // token
	CreateTime   string // 创建时间
	UpdateTime   string // 更新时间
	DeleteTime   string // 删除时间
}

// dealerColumns holds the columns for table dealer.
var dealerColumns = DealerColumns{
	Id:           "id",
	Name:         "name",
	Username:     "username",
	Password:     "password",
	Avatar:       "avatar",
	Phone:        "phone",
	DescribeInfo: "describe_info",
	Status:       "status",
	Token:        "token",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
	DeleteTime:   "delete_time",
}

// NewDealerDao creates and returns a new DAO object for table data access.
func NewDealerDao() *DealerDao {
	return &DealerDao{
		group:   "default",
		table:   "dealer",
		columns: dealerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DealerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DealerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DealerDao) Columns() DealerColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DealerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DealerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DealerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
