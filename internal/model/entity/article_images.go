// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ArticleImages is the golang structure for table article_images.
type ArticleImages struct {
	Id         int64       `json:"id"         ` // 主键ID
	ArticleId  int64       `json:"articleId"  ` // 文章ID
	Url        string      `json:"url"        ` // 图片URL
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	DeleteTime *gtime.Time `json:"deleteTime" ` // 删除时间
}
