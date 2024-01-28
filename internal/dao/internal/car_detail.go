// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CarDetailDao is the data access object for table car_detail.
type CarDetailDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns CarDetailColumns // columns contains all the column names of Table for convenient usage.
}

// CarDetailColumns defines and stores column names for table car_detail.
type CarDetailColumns struct {
	Id           string // 主键ID
	Year         string // 年份
	Brand        string // 品牌
	Model        string // 型号
	Version      string // 版本
	Image        string // 图片
	Category     string // 类型 0:其他, 1:轿车, 2:SUV, 3:MPV, 4:卡车, 5:跑车
	Color        string // 颜色
	Price        string // 指导价
	Type         string // 类型 0:其他, 1:纯电动, 2:插电混动, 3:增程, 4:汽油, 5:汽油+48V轻混系统, 6:油电混动, 7:柴油
	Seats        string // 座位数 0:7座以上 1:1座, 2:2座, 4:4座, 5:5座, 6:6座, 7:7座
	DescribeInfo string // 描述信息
	CreateTime   string // 创建时间
	UpdateTime   string // 更新时间
	DeleteTime   string // 删除时间
}

// carDetailColumns holds the columns for table car_detail.
var carDetailColumns = CarDetailColumns{
	Id:           "id",
	Year:         "year",
	Brand:        "brand",
	Model:        "model",
	Version:      "version",
	Image:        "image",
	Category:     "category",
	Color:        "color",
	Price:        "price",
	Type:         "type",
	Seats:        "seats",
	DescribeInfo: "describe_info",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
	DeleteTime:   "delete_time",
}

// NewCarDetailDao creates and returns a new DAO object for table data access.
func NewCarDetailDao() *CarDetailDao {
	return &CarDetailDao{
		group:   "default",
		table:   "car_detail",
		columns: carDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CarDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CarDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CarDetailDao) Columns() CarDetailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CarDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CarDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CarDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
