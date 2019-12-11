package openweatherapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/viktorstaikov/weather-dashboard-go/config"
	"github.com/viktorstaikov/weather-dashboard-go/services"
)

// OpenWeatherAPI ...
type OpenWeatherAPI struct {
	appID            string
	baseURL          string
	forecastEndpoint string
	uvEndpoint       string
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

// UVResponse ...
type UVResponse struct {
	Timestamp uint    `json:"date"`
	Value     float64 `json:"value"`
}

// MakeOpenWeatherAPI init
func MakeOpenWeatherAPI() *OpenWeatherAPI {
	config.Init("development")
	c := config.GetConfig()

	api := new(OpenWeatherAPI)

	api.appID = c.GetString("openWeather.appId")
	api.baseURL = c.GetString("openWeather.baseUrl")
	api.forecastEndpoint = c.GetString("openWeather.forecastEndpoint")
	api.uvEndpoint = c.GetString("openWeather.uvEndpoint")

	return api
}

// GetTempSeries from OpenWeatherAPI
func (api *OpenWeatherAPI) GetTempSeries() ([]services.TempData, error) {
	meta, err := api.makeForecastRequest()

	if err != nil {
		return nil, err
	}

	var mapped []services.TempData
	for _, item := range meta {
		entry := item.ToTempData()
		mapped = append(mapped, entry)
	}
	return mapped, nil
}

// GetRainSeries from OpenWeatherAPI
func (api *OpenWeatherAPI) GetRainSeries() ([]services.StatsData, error) {
	meta, err := api.makeForecastRequest()

	if err != nil {
		return nil, err
	}

	var mapped []services.StatsData
	for _, item := range meta {
		var entry services.StatsData
		entry.Timestamp = item.Timestamp
		entry.Value = item.Rain
		mapped = append(mapped, entry)
	}
	return mapped, nil
}

// GetPressureSeries from OpenWeatherAPI
func (api *OpenWeatherAPI) GetPressureSeries() ([]services.StatsData, error) {
	meta, err := api.makeForecastRequest()

	if err != nil {
		return nil, err
	}

	var mapped []services.StatsData
	for _, item := range meta {
		var entry services.StatsData
		entry.Timestamp = item.Timestamp
		entry.Value = float64(item.Pressure)
		mapped = append(mapped, entry)
	}
	return mapped, nil
}

// GetHumiditySeries from OpenWeatherAPI
func (api *OpenWeatherAPI) GetHumiditySeries() ([]services.StatsData, error) {
	meta, err := api.makeForecastRequest()

	if err != nil {
		return nil, err
	}

	var mapped []services.StatsData
	for _, item := range meta {
		entry := services.StatsData{
			Timestamp: item.Timestamp,
			Value:     float64(item.Humidity),
		}
		mapped = append(mapped, entry)
	}
	return mapped, nil
}

// GetForecast from OpenWeatherAPI
func (api *OpenWeatherAPI) GetForecast(date *time.Time) (*services.MetaForecast, error) {
	data, err := api.makeForecastRequest()
	if err != nil {
		return nil, err
	}

	uvData, uvErr := api.makeUVIndexRequest()
	if uvErr != nil {
		return nil, err
	}

	filteredData := filterSameDate(data, date)
	if len(filteredData) == 0 {
		return nil, errors.New(fmt.Sprintf(`no data for date %s`, date))
	}

	filteredUVData := filterSameDate(uvData, date)
	if len(filteredUVData) == 0 {
		return nil, errors.New(fmt.Sprintf(`no un data for date %s`, date))
	}

	avgData := averageForecast(filteredData)
	avgUvData := averageForecast(filteredUVData)

	res := &services.MetaForecast{
		Timestamp: avgUvData.Timestamp,
		Temp:      avgData.Temp,
		TempMin:   avgData.TempMin,
		TempMax:   avgData.TempMax,
		Pressure:  avgData.Pressure,
		Humidity:  avgData.Humidity,
		Weather:   avgData.Weather,
		Clouds:    avgData.Clouds,
		Wind:      avgData.Wind,
		Rain:      avgData.Rain,
		Snow:      avgData.Snow,
		UVIndex:   avgUvData.UVIndex,
	}
	return res, nil
}

func (api *OpenWeatherAPI) makeForecastRequest() ([]services.MetaForecast, error) {
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

func (api *OpenWeatherAPI) makeUVIndexRequest() ([]services.MetaForecast, error) {

	url := fmt.Sprintf("%s%s?lat=42.6979&lon=23.3222&appid=%s", api.baseURL, api.uvEndpoint, api.appID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	var forecast []UVResponse
	json.Unmarshal(bytes, &forecast)

	meta := parseUVResponse(forecast)
	return meta, nil
}

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

func parseUVResponse(r []UVResponse) []services.MetaForecast {
	var list []services.MetaForecast
	for _, respItem := range r {
		var f services.MetaForecast
		f.UVIndex = respItem.Value
		f.Timestamp = respItem.Timestamp * 1000

		list = append(list, f)
	}
	return list
}

func averageForecast(arr []services.MetaForecast) *services.MetaForecast {
	count := len(arr)

	total := &services.MetaForecast{
		Timestamp: arr[0].Timestamp,
		Temp:      0,
		TempMin:   0,
		TempMax:   0,
		Pressure:  0,
		Humidity:  0,
		Weather:   arr[count/2].Weather,
		Clouds:    0,
		Wind: services.WindCondition{
			Speed: 0,
			Deg:   0,
		},
		Rain:    0,
		Snow:    0,
		UVIndex: 0,
	}

	for _, item := range arr {
		total.Temp += item.Temp
		total.TempMin += item.TempMin
		total.TempMax += item.TempMax
		total.Pressure += item.Pressure
		total.Humidity += item.Humidity
		total.Clouds += item.Clouds
		total.Wind.Speed += item.Wind.Speed
		total.Wind.Deg += item.Wind.Deg
		total.Rain += item.Rain
		total.Snow += item.Snow
		total.UVIndex += item.UVIndex
	}

	average := &services.MetaForecast{
		Timestamp: total.Timestamp,
		Temp:      total.Temp / float64(count),
		TempMin:   total.TempMin / float64(count),
		TempMax:   total.TempMax / float64(count),
		Pressure:  total.Pressure / count,
		Humidity:  total.Humidity / count,
		Weather:   total.Weather,
		Clouds:    total.Clouds / count,
		Wind: services.WindCondition{
			Speed: total.Wind.Speed / float64(count),
			Deg:   total.Wind.Deg / count,
		},
		Rain:    total.Rain / float64(count),
		Snow:    total.Snow / float64(count),
		UVIndex: total.UVIndex / float64(count),
	}
	return average
}

func filterSameDate(arr []services.MetaForecast, date *time.Time) []services.MetaForecast {
	var result []services.MetaForecast
	for _, item := range arr {
		d := time.Unix(int64(item.Timestamp/1000), 0)
		if sameDate(&d, date) {
			result = append(result, item)
		}
	}
	return result
}

func sameDate(date1, date2 *time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
