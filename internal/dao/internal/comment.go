// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CommentDao is the data access object for table comment.
type CommentDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns CommentColumns // columns contains all the column names of Table for convenient usage.
}

// CommentColumns defines and stores column names for table comment.
type CommentColumns struct {
	Id          string // 主键ID
	OrderId     string // 订单ID
	Content     string // 评价内容
	TotalScore  string // 总评分 1~5星
	DealerScore string // 经销商评分 1~5星
	CarScore    string // 汽车评分 1~5星
	CreateTime  string // 创建时间
	DeleteTime  string // 删除时间
}

// commentColumns holds the columns for table comment.
var commentColumns = CommentColumns{
	Id:          "id",
	OrderId:     "order_id",
	Content:     "content",
	TotalScore:  "total_score",
	DealerScore: "dealer_score",
	CarScore:    "car_score",
	CreateTime:  "create_time",
	DeleteTime:  "delete_time",
}

// NewCommentDao creates and returns a new DAO object for table data access.
func NewCommentDao() *CommentDao {
	return &CommentDao{
		group:   "default",
		table:   "comment",
		columns: commentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CommentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CommentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CommentDao) Columns() CommentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CommentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CommentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CommentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
