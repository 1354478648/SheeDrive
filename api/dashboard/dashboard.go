package dashboard

import "github.com/gogf/gf/v2/frame/g"

type GetWeatherReq struct {
	g.Meta        `path:"/weather" method:"get"`
	Province      string `p:"province" v:"required#请输入省份信息" dc:"省"`
	City          string `p:"city" v:"required#请输入城市信息" dc:"市"`
	District      string `p:"district" v:"required#请输入区县信息" dc:"区县"`
	DetailAddress string `p:"detail_address" v:"required#请输入详细地址" dc:"详细地址"`
}

type GetWeatherRes struct {
	Province      string `json:"province" dc:"省"`
	City          string `json:"city" dc:"市"`
	Weather       string `json:"weather" dc:"天气现象"`
	Temperature   string `json:"temperature" dc:"温度"`
	Winddirection string `json:"winddirection" dc:"风向"`
	Windpower     string `json:"windpower" dc:"风力等级"`
	Humidity      string `json:"humidity" dc:"空气湿度"`
	Reporttime    string `json:"reporttime" dc:"数据发布时间"`
}
