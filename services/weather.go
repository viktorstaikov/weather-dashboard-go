package services

// "github.com/viktorstaikov/weather-dashboard-go/services/openweatherapi"

// WeatherAPI ...
type WeatherAPI interface {
	MakeForecastRequest() ([]MetaForecast, error)
}

// WeatherService ...
type WeatherService struct {
	api WeatherAPI
}

// WeatherCondition ...
type WeatherCondition struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// WindCondition ...
type WindCondition struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

// MetaForecast ...
type MetaForecast struct {
	Timestamp uint             `json:"timestamp"`
	Temp      float64          `json:"temp"`
	TempMin   float64          `json:"temp_min"`
	TempMax   float64          `json:"temp_max"`
	Pressure  int              `json:"pressure"`
	Humidity  int              `json:"humidity"`
	Weather   WeatherCondition `json:"weather"`
	Clouds    int              `json:"clouds"`
	Wind      WindCondition    `json:"wind"`
	Rain      float64          `json:"rain"`
	Snow      float64          `json:"snow"`
	UVIndex   float64          `json:"uvindex"`
}

// TempData represents temperature forecast at given time
type TempData struct {
	Timestamp uint    `json:"timestamp"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
}

// MakeWeatherService ...
func MakeWeatherService(weatherAPI WeatherAPI) *WeatherService {
	s := new(WeatherService)
	s.api = weatherAPI

	return s
}

// ToTempData ...
func (m MetaForecast) ToTempData() TempData {
	var item TempData
	item.TempMax = m.TempMax
	item.TempMin = m.TempMin
	item.Timestamp = m.Timestamp
	return item
}

// GetTempSeries returns data series for min and max temperature
func (w *WeatherService) GetTempSeries() ([]TempData, error) {

	meta, err := w.api.MakeForecastRequest()

	if err != nil {
		return nil, err
	}

	var mapped []TempData
	for _, item := range meta {
		entry := item.ToTempData()
		mapped = append(mapped, entry)
	}
	return mapped, nil
}
