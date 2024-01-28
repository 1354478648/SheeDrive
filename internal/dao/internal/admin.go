// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminDao is the data access object for table admin.
type AdminDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns AdminColumns // columns contains all the column names of Table for convenient usage.
}

// AdminColumns defines and stores column names for table admin.
type AdminColumns struct {
	Id         string // 主键ID
	Name       string // 姓名
	Username   string // 用户名
	Password   string // 密码
	Avatar     string // 头像
	Phone      string // 手机号
	Status     string // 状态 0:禁用, 1:正常
	IsRoot     string // 是否是超级管理员 0:否, 1:是
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
}

// adminColumns holds the columns for table admin.
var adminColumns = AdminColumns{
	Id:         "id",
	Name:       "name",
	Username:   "username",
	Password:   "password",
	Avatar:     "avatar",
	Phone:      "phone",
	Status:     "status",
	IsRoot:     "isRoot",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
}

// NewAdminDao creates and returns a new DAO object for table data access.
func NewAdminDao() *AdminDao {
	return &AdminDao{
		group:   "default",
		table:   "admin",
		columns: adminColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminDao) Columns() AdminColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
