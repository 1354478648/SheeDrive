package dashboard

import (
	apiDashboard "SheeDrive/api/dashboard"
	"SheeDrive/utility"
	"context"
	"fmt"
)

var DashboardController = &cDashboard{}

type cDashboard struct{}

func (c *cDashboard) GetWeather(ctx context.Context, req *apiDashboard.GetWeatherReq) (res *apiDashboard.GetWeatherRes, err error) {
	geocode, err := utility.Geocoding(fmt.Sprintf("%s%s%s%s", req.Province, req.City, req.District, req.DetailAddress), req.City)
	if err != nil {
		return nil, err
	}

	weather, err := utility.GetWeather(geocode.Adcode)
	if err != nil {
		return nil, err
	}

	res = &apiDashboard.GetWeatherRes{
		Province:      weather.Province,
		City:          weather.City,
		Weather:       weather.Weather,
		Temperature:   weather.Temperature,
		Winddirection: weather.Winddirection,
		Windpower:     weather.Windpower,
		Humidity:      weather.Humidity,
		Reporttime:    weather.Reporttime,
	}

	return
}
