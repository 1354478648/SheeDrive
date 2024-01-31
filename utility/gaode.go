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
// Latitude         float64 `json:"latitude"`
// Longitude        float64 `json:"longitude"`
// FormattedAddress string  `json:"formatted_address"`
// Province         string  `json:"province"`
// City             string  `json:"city"`
// District         string  `json:"district"`
// Street           string  `json:"street"`
// StreetNumber     string  `json:"number"`
// PoiName          string  `json:"poiName"`
type Response struct {
	Geocodes []Geocode `json:"geocodes"`
}

type Geocode struct {
	FormattedAddress string      `json:"formatted_address"`
	Province         string      `json:"province"`
	City             string      `json:"city"`
	District         string      `json:"district"`
	Street           string      `json:"street"`
	Number           interface{} `json:"number"`
	Location         string      `json:"location"`
}

// 地理编码
// 结构化地址 -> 经纬度
func Geocoding(address, city string) (response *Response, err error) {
	// 设置高德地图地理编码API的URL和参数
	url := "https://restapi.amap.com/v3/geocode/geo"
	parameters := fmt.Sprintf("?key=%s&address=%s&city=%s", consts.GaodeKey, address, city)

	fmt.Println("-------------")
	fmt.Println(url + parameters)

	result, err := http.Get(url + parameters)
	if err != nil {
		return nil, gerror.New("高德API发送请求失败")
	}
	defer result.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, gerror.New("读取响应时发生错误")
	}
	fmt.Println("-------------")
	fmt.Println(string(body))

	// 解析 JSON 数据
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		fmt.Println(err)
		return nil, gerror.New("解析JSON数据时发生错误")
	}
	fmt.Println("-------------")
	fmt.Println(response)

	if len(response.Geocodes) > 0 {
		fmt.Println("-------------")
		fmt.Println("Formatted Address:", response.Geocodes[0])
	}

	return response, nil
}

// 逆向地理编码
// 经纬度 -> 结构化地址
// func ReGeocoding()
