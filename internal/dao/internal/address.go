// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AddressDao is the data access object for table address.
type AddressDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns AddressColumns // columns contains all the column names of Table for convenient usage.
}

// AddressColumns defines and stores column names for table address.
type AddressColumns struct {
	Id             string // 主键ID
	BelongId       string // 所属ID
	BelongCategory string // 所属分类 1:经销商,2:用户
	LngLat         string // 经纬度
	Province       string // 省
	City           string // 市
	District       string // 区
	Detail         string // 详细地址
	CreateTime     string // 创建时间
	UpdateTime     string // 更新时间
	DeleteTime     string // 删除时间
}

// addressColumns holds the columns for table address.
var addressColumns = AddressColumns{
	Id:             "id",
	BelongId:       "belong_id",
	BelongCategory: "belong_category",
	LngLat:         "lng_lat",
	Province:       "province",
	City:           "city",
	District:       "district",
	Detail:         "detail",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
	DeleteTime:     "delete_time",
}

// NewAddressDao creates and returns a new DAO object for table data access.
func NewAddressDao() *AddressDao {
	return &AddressDao{
		group:   "default",
		table:   "address",
		columns: addressColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AddressDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AddressDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AddressDao) Columns() AddressColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AddressDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AddressDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AddressDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
