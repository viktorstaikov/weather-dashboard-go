package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/viktorstaikov/weather-dashboard-go/config"
)

// GetTempSeries returns data series for min and max temperature
func GetTempSeries() ([]TempData, error) {

	config := config.GetConfig()
	appID := config.GetString("openWeatherAppId")
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?lat=42.6979&lon=23.3222&appid=%s&units=metric", appID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	var forecast forecastResponse
	json.Unmarshal(bytes, &forecast)

	meta := forecast.toMetaForecast()
	var mapped []TempData
	for _, item := range meta {
		entry := item.toTempData()
		mapped = append(mapped, entry)
	}
	return mapped, nil
}
