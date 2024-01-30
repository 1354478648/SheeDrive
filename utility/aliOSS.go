package utility

import (
	"SheeDrive/internal/consts"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

// 创建OSSClient实例
func initOSSClient() (bucket *oss.Bucket, err error) {
	// 存储空间名称
	bucketName := consts.BucketName

	// 从环境变量中获取访问凭证
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		handleError(err)
	}

	// 创建OSSClient实例。
	// Endpoint
	client, err := oss.New("https://"+consts.EndpointName, "", "", oss.SetCredentialsProvider(&provider))
	if err != nil {
		handleError(err)
	}

	// 获取存储空间
	bucket, err = client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}

	return
}

// 上传文件
func UploadFile(objectName, localFileName string) (url string) {
	bucket, err := initOSSClient()
	if err != nil {
		handleError(err)
	}
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		handleError(err)
	}
	url = fmt.Sprintf("https://%s.%s/%s", consts.BucketName, consts.EndpointName, objectName)
	return
}
