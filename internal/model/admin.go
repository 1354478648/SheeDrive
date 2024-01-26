package model

import (
	"SheeDrive/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

// AdminGetListInput 分页与关键字查询管理员列表
type AdminGetListInput struct {
	Page       int
	PageSize   int
	Username   string
	Name       string
	BeforeDate *gtime.Time
	AfterDate  *gtime.Time
}

// AdminGetListOutput 分页与关键字查询管理员列表结果
type AdminGetListOutput struct {
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
	Total    int             `json:"total"`
	Items    []*entity.Admin `json:"items"`
}
