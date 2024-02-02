package stock

import (
	"SheeDrive/api/pagination"

	"github.com/gogf/gf/v2/frame/g"
)

// 库存列表分页与关键字查询（管理员）
type StockGetListReq struct {
	g.Meta `path:"/list" method:"get"`
	pagination.CommonPaginationReq
	//关键字查询可选字段
	DealerName string `json:"dealer_name" dc:"经销商名称"`
}

type StockGetListRes struct {
	pagination.CommonPaginationRes
}
