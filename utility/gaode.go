package utility

import (
	"SheeDrive/internal/consts"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 封装返回结果
type Location struct {
	latitude         float64
	longitude        float64
	formattedAddress string
	Province         string
	City             string
	District         string
	Street           string
	StreetNumber     string
	PoiName          string
}

// 地理编码
// 结构化地址 -> 经纬度
func Geocoding(address, city string) (location *Location, err error) {
	// 设置高德地图地理编码API的URL和参数
	url := "https://restapi.amap.com/v3/geocode/geo"
	parameters := "?key=" + consts.GaodeKey + "&address=" + address + "&city=" + city

	fmt.Println(url + parameters)

	response, err := http.Get(url + parameters)
	if err != nil {
		fmt.Println("发送请求时发生错误:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应时发生错误:", err)
		return
	}
	fmt.Println(string(body))
	return
}

// 逆向地理编码
// 经纬度 -> 结构化地址
// func ReGeocoding()
