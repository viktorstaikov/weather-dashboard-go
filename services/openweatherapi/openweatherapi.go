package openweatherapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/viktorstaikov/weather-dashboard-go/config"
)

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

// MakeForecastRequest ...
func MakeForecastRequest() (*ForecastResponse, error) {
	config := config.GetConfig()
	appID := config.GetString("openWeather.appId")
	baseURL := config.GetString("openWeather.baseUrl")
	endpoint := config.GetString("openWeather.forecastEndpoint")

	url := fmt.Sprintf("%s%s?lat=42.6979&lon=23.3222&appid=%s&units=metric", baseURL, endpoint, appID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	var forecast ForecastResponse
	json.Unmarshal(bytes, &forecast)

	return &forecast, nil
}
