package openweatherapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/viktorstaikov/weather-dashboard-go/config"
	"github.com/viktorstaikov/weather-dashboard-go/services"
)

// OpenWeatherAPI ...
type OpenWeatherAPI struct {
	appID            string
	baseURL          string
	forecastEndpoint string
}

// ForecastResponse ...
type ForecastResponse struct {
	List []struct {
		Dt   uint `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  int     `json:"pressure"`
			SeaLevel  int     `json:"sea_level"`
			GrndLevel int     `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
		} `json:"wind"`
		Rain struct {
			ThreeH float64 `json:"3h"`
		} `json:"rain"`
		Snow struct {
			ThreeH float64 `json:"3h"`
		} `json:"snow"`
		Sys struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
	} `json:"list"`
}

// parseForecastResponse ...
func parseForecastResponse(r *ForecastResponse) []services.MetaForecast {
	var list []services.MetaForecast
	for _, item := range r.List {
		var metaItem services.MetaForecast
		metaItem.Timestamp = item.Dt * 1000
		metaItem.Temp = item.Main.Temp
		metaItem.TempMin = item.Main.TempMin
		metaItem.TempMax = item.Main.TempMax
		metaItem.Pressure = item.Main.Pressure
		metaItem.Humidity = item.Main.Humidity
		metaItem.Weather = item.Weather[0]
		metaItem.Wind = item.Wind

		metaItem.Clouds = 0
		if item.Clouds.All >= 0 {
			metaItem.Clouds = item.Clouds.All
		}

		metaItem.Rain = 0
		if item.Rain.ThreeH >= 0 {
			metaItem.Rain = item.Rain.ThreeH
		}

		metaItem.Snow = 0
		if item.Snow.ThreeH >= 0 {
			metaItem.Snow = item.Snow.ThreeH
		}

		list = append(list, metaItem)
	}
	return list
}

// MakeForecastRequest ...
func (api *OpenWeatherAPI) MakeForecastRequest() ([]services.MetaForecast, error) {
	url := fmt.Sprintf("%s%s?lat=42.6979&lon=23.3222&appid=%s&units=metric", api.baseURL, api.forecastEndpoint, api.appID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	var forecast ForecastResponse
	json.Unmarshal(bytes, &forecast)

	meta := parseForecastResponse(&forecast)
	return meta, nil
}

// MakeOpenWeatherAPI ...
func MakeOpenWeatherAPI() *OpenWeatherAPI {
	config.Init("development")
	c := config.GetConfig()

	api := new(OpenWeatherAPI)

	api.appID = c.GetString("openWeather.appId")
	api.baseURL = c.GetString("openWeather.baseUrl")
	api.forecastEndpoint = c.GetString("openWeather.forecastEndpoint")

	return api
}
