package file

import (
	apiFile "SheeDrive/api/file"
	"SheeDrive/internal/consts"
	"SheeDrive/utility"
	"context"
	"os"
	"path"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var FileController = &cFile{}

type cFile struct{}

// 单文件上传
func (c *cFile) FileUpload(ctx context.Context, req *apiFile.FileUploadReq) (res *apiFile.FileUploadRes, err error) {
	const (
		tempLocalPath    = consts.TempLocalPath
		imagesObjectName = consts.ImagesObjectName
	)

	// 获取文件
	file := req.File
	// 判断文件是否为空
	if file == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}

	// 生成UUID防止文件重名
	// path.Ext方法能够获取文件扩展名
	file.Filename = utility.GenStrUUID() + path.Ext(file.Filename)

	// 先将文件临时存到本地中
	fileName, err := file.Save(tempLocalPath)
	if err != nil {
		return nil, err
	}

	url := utility.UploadFile(imagesObjectName+fileName, tempLocalPath+fileName)

	// 删除本地临时文件
	err = os.RemoveAll(tempLocalPath)
	if err != nil {
		return nil, err
	}

	// 定义返回结构体
	res = &apiFile.FileUploadRes{
		Url: url,
	}

	return
}
