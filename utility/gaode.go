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
type GeocodeResponse struct {
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
	Adcode           string      `json:"adcode" dc:"区域编码"`
	Location         string      `json:"location" dc:"经纬度"`
}

type WeatherResponse struct {
	Lives []Lives `json:"lives"`
}

type Lives struct {
	Province      string `json:"province" dc:"省"` // 也可以是市
	City          string `json:"city" dc:"市"`     // 也可以是区
	Weather       string `json:"weather" dc:"天气现象"`
	Temperature   string `json:"temperature" dc:"温度"`
	Winddirection string `json:"winddirection" dc:"风向"`
	Windpower     string `json:"windpower" dc:"风力等级"`
	Humidity      string `json:"humidity" dc:"空气湿度"`
	Reporttime    string `json:"reporttime" dc:"数据发布时间"`
}

// type ReGeocodeResponse struct {
// 	ReGeocode ReGeocode `json:"regeocode"`
// }

// type ReGeocode struct {
// 	AddressComponent AddressComponent `json:"addressComponent"`
// 	FormattedAddress string           `json:"formatted_address"`
// }

// type AddressComponent struct {
// 	Province     string       `json:"province"`
// 	City         interface{}  `json:"city"`
// 	District     string       `json:"district"`
// 	StreetNumber StreetNumber `json:"streetNumber"`
// }

// type StreetNumber struct {
// 	Number string `json:"number"`
// 	Street string `json:"street"`
// }

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
	var response GeocodeResponse
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

// 获取天气数据
func GetWeather(adcode string) (lives *Lives, err error) {
	// 设置高德地图地理编码API的URL和参数
	url := "https://restapi.amap.com/v3/weather/weatherInfo"
	parameters := fmt.Sprintf("?key=%s&city=%s", consts.GaodeKey, adcode)

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
	var response WeatherResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, gerror.New("解析JSON数据时发生错误")
	}

	// 返回的可能有多个数据，选择第一个
	if len(response.Lives) > 0 {
		return &response.Lives[0], nil
	}
	return nil, gerror.New("未找到匹配的地址")
}

// 逆向地理编码
// 经纬度 -> 结构化地址
// func ReGeocoding(lnglat string) (reGeocode *ReGeocode, err error) {
// 	// 设置高德地图地理编码API的URL和参数
// 	url := "https://restapi.amap.com/v3/geocode/regeo"
// 	parameters := fmt.Sprintf("?key=%s&location=%s", consts.GaodeKey, lnglat)
// 	fmt.Println("-------------")
// 	fmt.Println(url + parameters)
// 	// 发送HTTP请求
// 	result, err := http.Get(url + parameters)
// 	if err != nil {
// 		return nil, gerror.New("发送高德API请求时发生错误")
// 	}
// 	defer result.Body.Close()

// 	// 读取响应内容
// 	body, err := io.ReadAll(result.Body)
// 	if err != nil {
// 		return nil, gerror.New("读取高德API响应时发生错误")
// 	}
// 	fmt.Println("-------------")
// 	fmt.Println(string(body))
// 	// 解析JSON数据
// 	var response ReGeocodeResponse
// 	err = json.Unmarshal([]byte(body), &response)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, gerror.New("解析JSON数据时发生错误")
// 	}
// 	fmt.Println("-------------")
// 	fmt.Println(&response.ReGeocode)
// 	return &response.ReGeocode, nil
// }
