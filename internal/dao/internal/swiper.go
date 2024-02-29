// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SwiperDao is the data access object for table swiper.
type SwiperDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns SwiperColumns // columns contains all the column names of Table for convenient usage.
}

// SwiperColumns defines and stores column names for table swiper.
type SwiperColumns struct {
	Id           string // 主键ID
	CarId        string // 车辆ID
	ImageUrl     string // 图片地址
	DescribeInfo string // 描述信息
	CreateTime   string // 创建时间
	DeleteTime   string // 删除时间
}

// swiperColumns holds the columns for table swiper.
var swiperColumns = SwiperColumns{
	Id:           "id",
	CarId:        "car_id",
	ImageUrl:     "image_url",
	DescribeInfo: "describe_info",
	CreateTime:   "create_time",
	DeleteTime:   "delete_time",
}

// NewSwiperDao creates and returns a new DAO object for table data access.
func NewSwiperDao() *SwiperDao {
	return &SwiperDao{
		group:   "default",
		table:   "swiper",
		columns: swiperColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SwiperDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SwiperDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SwiperDao) Columns() SwiperColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SwiperDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SwiperDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SwiperDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
