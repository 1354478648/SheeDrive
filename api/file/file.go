package file

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/upload" method:"post" mime:"multipart/form-data"`
	File   *ghttp.UploadFile `p:"file" type:"file" dc:"上传文件"`
}

type FileUploadRes struct {
	Url string `json:"url" dc:"文件路径"`
}
