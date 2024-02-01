package utility

import (
	"SheeDrive/internal/consts"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gogf/gf/errors/gerror"
)

// 封装返回结果
type Response struct {
	Geocodes []Geocode `json:"geocodes"`
}

// 地理编码（保留冗余字段，后续可能有用处）
type Geocode struct {
	FormattedAddress string      `json:"formatted_address" dc:"结构化地址"`
	Province         string      `json:"province" dc:"省"`
	City             string      `json:"city" dc:"市"`
	District         string      `json:"district" dc:"区"`
	Street           string      `json:"street" dc:"街道"`
	Number           interface{} `json:"number" dc:"门牌号"`
	Location         string      `json:"location" dc:"经纬度"`
}

// 地理编码
// 结构化地址 -> 经纬度
func Geocoding(address, city string) (geocode *Geocode, err error) {
	// 设置高德地图地理编码API的URL和参数
	url := "https://restapi.amap.com/v3/geocode/geo"
	parameters := fmt.Sprintf("?key=%s&address=%s&city=%s", consts.GaodeKey, address, city)

	// 发送HTTP请求
	result, err := http.Get(url + parameters)
	if err != nil {
		return nil, gerror.New("发送高德API请求时发生错误")
	}
	defer result.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, gerror.New("读取高德API响应时发生错误")
	}

	// 解析JSON数据
	var response Response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, gerror.New("解析JSON数据时发生错误")
	}

	// 返回的可能有多个数据，选择第一个
	if len(response.Geocodes) > 0 {
		return &response.Geocodes[0], nil
	}
	return nil, gerror.New("未找到匹配的地址")
}

// 逆向地理编码
// 经纬度 -> 结构化地址
// func ReGeocoding()
