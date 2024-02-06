// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ArticleImages is the golang structure of table article_images for DAO operations like Where/Data.
type ArticleImages struct {
	g.Meta     `orm:"table:article_images, do:true"`
	Id         interface{} // 主键ID
	ArticleId  interface{} // 文章ID
	Url        interface{} // 图片URL
	CreateTime *gtime.Time // 创建时间
	DeleteTime *gtime.Time // 删除时间
}
