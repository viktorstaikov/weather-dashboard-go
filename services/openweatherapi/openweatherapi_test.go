package openweatherapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/viktorstaikov/weather-dashboard-go/config"
	"github.com/viktorstaikov/weather-dashboard-go/services"
)

func TestGetTempSeries(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonString := `{"list":[{"dt":1575579600,"main":{"temp":-1,"temp_min":-3,"temp_max":0,"pressure":1031,"sea_level":1031,"grnd_level":913,"humidity":89,"temp_kf":1.58},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.39,"deg":233},"rain":{"3h":0},"snow":{"3h":0},"sys":{"pod":"n"},"dt_txt":"2019-12-05 21:00:00"},{"dt":1575590400,"main":{"temp":3,"temp_min":1,"temp_max":5,"pressure":1030,"sea_level":1030,"grnd_level":913,"humidity":89,"temp_kf":1.19},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.07,"deg":234},"rain":{"3h":20},"snow":{"3h":10},"sys":{"pod":"n"},"dt_txt":"2019-12-06 00:00:00"}]}`

		w.Write([]byte(jsonString))
	}))
	defer ts.Close()

	config.Init("development")
	c := config.GetConfig()
	c.Set("openWeather.baseUrl", ts.URL)

	api := MakeOpenWeatherAPI(c)

	data, err := api.GetTempSeries()
	assert.Nil(t, err)
	assert.Len(t, data, 2)
}
func TestGetRainSeries(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonString := `{"list":[{"dt":1575579600,"main":{"temp":-1,"temp_min":-3,"temp_max":0,"pressure":1031,"sea_level":1031,"grnd_level":913,"humidity":89,"temp_kf":1.58},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.39,"deg":233},"rain":{"3h":0},"snow":{"3h":0},"sys":{"pod":"n"},"dt_txt":"2019-12-05 21:00:00"},{"dt":1575590400,"main":{"temp":3,"temp_min":1,"temp_max":5,"pressure":1030,"sea_level":1030,"grnd_level":913,"humidity":89,"temp_kf":1.19},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.07,"deg":234},"rain":{"3h":20},"snow":{"3h":10},"sys":{"pod":"n"},"dt_txt":"2019-12-06 00:00:00"}]}`

		w.Write([]byte(jsonString))
	}))
	defer ts.Close()

	config.Init("development")
	c := config.GetConfig()
	c.Set("openWeather.baseUrl", ts.URL)

	api := MakeOpenWeatherAPI(c)

	data, err := api.GetRainSeries()
	assert.Nil(t, err)
	assert.Len(t, data, 2)
}
func TestGetPressureSeries(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonString := `{"list":[{"dt":1575579600,"main":{"temp":-1,"temp_min":-3,"temp_max":0,"pressure":1031,"sea_level":1031,"grnd_level":913,"humidity":89,"temp_kf":1.58},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.39,"deg":233},"rain":{"3h":0},"snow":{"3h":0},"sys":{"pod":"n"},"dt_txt":"2019-12-05 21:00:00"},{"dt":1575590400,"main":{"temp":3,"temp_min":1,"temp_max":5,"pressure":1030,"sea_level":1030,"grnd_level":913,"humidity":89,"temp_kf":1.19},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.07,"deg":234},"rain":{"3h":20},"snow":{"3h":10},"sys":{"pod":"n"},"dt_txt":"2019-12-06 00:00:00"}]}`

		w.Write([]byte(jsonString))
	}))
	defer ts.Close()

	config.Init("development")
	c := config.GetConfig()
	c.Set("openWeather.baseUrl", ts.URL)

	api := MakeOpenWeatherAPI(c)

	data, err := api.GetPressureSeries()
	assert.Nil(t, err)
	assert.Len(t, data, 2)
}
func TestGetHumiditySeries(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonString := `{"list":[{"dt":1575579600,"main":{"temp":-1,"temp_min":-3,"temp_max":0,"pressure":1031,"sea_level":1031,"grnd_level":913,"humidity":89,"temp_kf":1.58},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.39,"deg":233},"rain":{"3h":0},"snow":{"3h":0},"sys":{"pod":"n"},"dt_txt":"2019-12-05 21:00:00"},{"dt":1575590400,"main":{"temp":3,"temp_min":1,"temp_max":5,"pressure":1030,"sea_level":1030,"grnd_level":913,"humidity":89,"temp_kf":1.19},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.07,"deg":234},"rain":{"3h":20},"snow":{"3h":10},"sys":{"pod":"n"},"dt_txt":"2019-12-06 00:00:00"}]}`

		w.Write([]byte(jsonString))
	}))
	defer ts.Close()

	config.Init("development")
	c := config.GetConfig()
	c.Set("openWeather.baseUrl", ts.URL)

	api := MakeOpenWeatherAPI(c)

	data, err := api.GetHumiditySeries()
	assert.Nil(t, err)
	assert.Len(t, data, 2)
}
func TestGetForecast(t *testing.T) {
	config.Init("development")
	c := config.GetConfig()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jsonString string

		path := r.URL.EscapedPath()
		forecastEndpoint := c.GetString("openWeather.forecastEndpoint")
		if path == forecastEndpoint {
			jsonString = `{"list":[{"dt":1573065600,"main":{"temp":-1,"temp_min":-3,"temp_max":0,"pressure":1031,"sea_level":1031,"grnd_level":913,"humidity":89,"temp_kf":1.58},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.39,"deg":233},"rain":{"3h":0},"snow":{"3h":0},"sys":{"pod":"n"},"dt_txt":"2019-12-05 21:00:00"},{"dt":1575590400,"main":{"temp":3,"temp_min":1,"temp_max":5,"pressure":1030,"sea_level":1030,"grnd_level":913,"humidity":89,"temp_kf":1.19},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.07,"deg":234},"rain":{"3h":20},"snow":{"3h":10},"sys":{"pod":"n"},"dt_txt":"2019-12-06 00:00:00"}]}`
		} else {
			jsonString = `[{"lat":42.6979,"lon":23.3222,"date_iso":"2019-12-11T12:00:00Z","date":1573065600,"value":1.2},{"lat":42.6979,"lon":23.3222,"date_iso":"2019-12-12T12:00:00Z","date":1576152000,"value":1.09},{"lat":42.6979,"lon":23.3222,"date_iso":"2019-12-13T12:00:00Z","date":1576238400,"value":1.05},{"lat":42.6979,"lon":23.3222,"date_iso":"2019-12-14T12:00:00Z","date":1576324800,"value":1.05},{"lat":42.6979,"lon":23.3222,"date_iso":"2019-12-15T12:00:00Z","date":1576411200,"value":1.11},{"lat":42.6979,"lon":23.3222,"date_iso":"2019-12-16T12:00:00Z","date":1576497600,"value":1.09},{"lat":42.6979,"lon":23.3222,"date_iso":"2019-12-17T12:00:00Z","date":1576584000,"value":1.08},{"lat":42.6979,"lon":23.3222,"date_iso":"2019-12-18T12:00:00Z","date":1576670400,"value":1.08}]`
		}

		w.Write([]byte(jsonString))
	}))
	defer ts.Close()

	c.Set("openWeather.baseUrl", ts.URL)

	api := MakeOpenWeatherAPI(c)

	date := time.Unix(1573065600, 0)
	data, err := api.GetForecast(&date)

	assert.Nil(t, err)
	assert.NotNil(t, data)
}

func TestGetForecastToday(t *testing.T) {
	config.Init("development")
	c := config.GetConfig()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jsonString string

		path := r.URL.EscapedPath()
		forecastEndpoint := c.GetString("openWeather.forecastEndpoint")

		timestamp := time.Now().Unix()
		if path == forecastEndpoint {
			jsonString = fmt.Sprintf(`{"list":[{"dt":%s,"main":{"temp":-1,"temp_min":-3,"temp_max":0,"pressure":1031,"sea_level":1031,"grnd_level":913,"humidity":89,"temp_kf":1.58},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.39,"deg":233},"rain":{"3h":0},"snow":{"3h":0},"sys":{"pod":"n"},"dt_txt":"2019-12-05 21:00:00"},{"dt":1575590400,"main":{"temp":3,"temp_min":1,"temp_max":5,"pressure":1030,"sea_level":1030,"grnd_level":913,"humidity":89,"temp_kf":1.19},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.07,"deg":234},"rain":{"3h":20},"snow":{"3h":10},"sys":{"pod":"n"},"dt_txt":"2019-12-06 00:00:00"}]}`, strconv.FormatInt(timestamp, 10))
		} else {
			jsonString = fmt.Sprintf(`{"lat":42.6979,"lon":23.3222,"date_iso":"2019-12-11T12:00:00Z","date":%s,"value":1.2}`, strconv.FormatInt(timestamp, 10))
			fmt.Println(jsonString)
		}

		w.Write([]byte(jsonString))
	}))
	defer ts.Close()

	c.Set("openWeather.baseUrl", ts.URL)

	api := MakeOpenWeatherAPI(c)

	date := time.Now()
	data, err := api.GetForecast(&date)

	assert.Nil(t, err)
	assert.NotNil(t, data)
}

func TestGetForecastNoData(t *testing.T) {
	config.Init("development")
	c := config.GetConfig()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jsonString string

		path := r.URL.EscapedPath()
		forecastEndpoint := c.GetString("openWeather.forecastEndpoint")
		if path == forecastEndpoint {
			jsonString = `{"list":[{"dt":1576065600,"main":{"temp":-1,"temp_min":-3,"temp_max":0,"pressure":1031,"sea_level":1031,"grnd_level":913,"humidity":89,"temp_kf":1.58},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.39,"deg":233},"rain":{"3h":0},"snow":{"3h":0},"sys":{"pod":"n"},"dt_txt":"2019-12-05 21:00:00"},{"dt":1575590400,"main":{"temp":3,"temp_min":1,"temp_max":5,"pressure":1030,"sea_level":1030,"grnd_level":913,"humidity":89,"temp_kf":1.19},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.07,"deg":234},"rain":{"3h":20},"snow":{"3h":10},"sys":{"pod":"n"},"dt_txt":"2019-12-06 00:00:00"}]}`
		} else {
			jsonString = `[{"lat":42.6979,"lon":23.3222,"date_iso":"2019-12-11T12:00:00Z","date":1574064600,"value":1.2}]`
		}

		w.Write([]byte(jsonString))
	}))
	defer ts.Close()

	c.Set("openWeather.baseUrl", ts.URL)

	api := MakeOpenWeatherAPI(c)

	date := time.Unix(1574064600, 0)
	data, err := api.GetForecast(&date)

	assert.NotNil(t, err)
	assert.Nil(t, data)
}

func TestGetForecastNoUV(t *testing.T) {
	config.Init("development")
	c := config.GetConfig()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jsonString string

		path := r.URL.EscapedPath()
		forecastEndpoint := c.GetString("openWeather.forecastEndpoint")
		if path == forecastEndpoint {
			jsonString = `{"list":[{"dt":1573065600,"main":{"temp":-1,"temp_min":-3,"temp_max":0,"pressure":1031,"sea_level":1031,"grnd_level":913,"humidity":89,"temp_kf":1.58},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.39,"deg":233},"rain":{"3h":0},"snow":{"3h":0},"sys":{"pod":"n"},"dt_txt":"2019-12-05 21:00:00"},{"dt":1575590400,"main":{"temp":3,"temp_min":1,"temp_max":5,"pressure":1030,"sea_level":1030,"grnd_level":913,"humidity":89,"temp_kf":1.19},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.07,"deg":234},"rain":{"3h":20},"snow":{"3h":10},"sys":{"pod":"n"},"dt_txt":"2019-12-06 00:00:00"}]}`
		} else {
			jsonString = `[{"lat":42.6979,"lon":23.3222,"date_iso":"2019-12-11T12:00:00Z","date":1574064600,"value":1.2}]`

		}

		w.Write([]byte(jsonString))
	}))
	defer ts.Close()

	c.Set("openWeather.baseUrl", ts.URL)

	api := MakeOpenWeatherAPI(c)

	date := time.Unix(1573065600, 0)
	data, err := api.GetForecast(&date)

	assert.NotNil(t, err)
	assert.Nil(t, data)
}

func TestMakeForecastRequest(t *testing.T) {
	config.Init("development")
	c := config.GetConfig()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		assert.Equal(t, r.Method, "GET")

		path := r.URL.EscapedPath()
		defaultPath := c.GetString("openWeather.forecastEndpoint")
		assert.Equal(t, defaultPath, path)

		query := r.URL.Query()

		lat := query.Get("lat")
		assert.Equal(t, "42.6979", lat)
		lon := query.Get("lon")
		assert.Equal(t, "23.3222", lon)

		defaultAppID := c.GetString("openWeather.appId")
		appID := query.Get("appid")
		assert.Equal(t, defaultAppID, appID)

		units := query.Get("units")
		assert.Equal(t, "metric", units)

	}))
	defer ts.Close()

	c.Set("openWeather.baseUrl", ts.URL)

	api := MakeOpenWeatherAPI(c)

	_, err := api.makeForecastRequest()
	if err != nil {
		t.Errorf("makeForecastRequest() returned an error: %s", err)
	}
}

func TestMakeUVIndexForecastRequest(t *testing.T) {
	config.Init("development")
	c := config.GetConfig()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		assert.Equal(t, r.Method, "GET")

		path := r.URL.EscapedPath()
		defaultPath := c.GetString("openWeather.uvForecastEndpoint")
		assert.Equal(t, defaultPath, path)

		query := r.URL.Query()

		lat := query.Get("lat")
		assert.Equal(t, "42.6979", lat)
		lon := query.Get("lon")
		assert.Equal(t, "23.3222", lon)

		defaultAppID := c.GetString("openWeather.appId")
		appID := query.Get("appid")
		assert.Equal(t, defaultAppID, appID)

	}))
	defer ts.Close()

	c.Set("openWeather.baseUrl", ts.URL)

	api := MakeOpenWeatherAPI(c)

	_, err := api.makeUVIndexForecastRequest()
	if err != nil {
		t.Errorf("makeUVIndexForecastRequest() returned an error: %s", err)
	}
}

func TestParseForecastResponse(t *testing.T) {
	jsonString := `{"list":[{"dt":1575579600,"main":{"temp":-1,"temp_min":-3,"temp_max":0,"pressure":1031,"sea_level":1031,"grnd_level":913,"humidity":89,"temp_kf":1.58},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.39,"deg":233},"rain":{"3h":0},"snow":{"3h":0},"sys":{"pod":"n"},"dt_txt":"2019-12-05 21:00:00"},{"dt":1575590400,"main":{"temp":3,"temp_min":1,"temp_max":5,"pressure":1030,"sea_level":1030,"grnd_level":913,"humidity":89,"temp_kf":1.19},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.07,"deg":234},"rain":{"3h":20},"snow":{"3h":10},"sys":{"pod":"n"},"dt_txt":"2019-12-06 00:00:00"}]}`

	var forecast ForecastResponse
	json.Unmarshal([]byte(jsonString), &forecast)

	metaForecastSlice := parseForecastResponse(&forecast)

	assert.Equal(t, 2, len(metaForecastSlice))
}

func TestParseUVResponse(t *testing.T) {
	jsonString := `{"lat":42.6979,"lon":23.3222,"date_iso":"2019-12-11T12:00:00Z","date":1576065600,"value":1.2}`

	var uvResponse UVResponse
	json.Unmarshal([]byte(jsonString), &uvResponse)

	uvData := parseUVResponse(&uvResponse)

	assert.NotNil(t, uvData)
	assert.Equal(t, 1.2, uvData.UVIndex)
}

func TestAverageForecast(t *testing.T) {
	arr := []services.MetaForecast{
		services.MetaForecast{
			Timestamp: 1576151100000,
			Temp:      -1,
			TempMin:   -2,
			TempMax:   -1,
			Pressure:  1030,
			Humidity:  89,
			Weather: services.WeatherCondition{
				ID:          800,
				Main:        "Clear",
				Description: "clear sky",
				Icon:        "01n",
			},
			Clouds: 0,
			Wind: services.WindCondition{
				Speed: 1.07,
				Deg:   234,
			},
			Rain:    10,
			Snow:    10,
			UVIndex: 0,
		},
		services.MetaForecast{
			Timestamp: 1576237500000,
			Temp:      2,
			TempMin:   1,
			TempMax:   2,
			Pressure:  1030,
			Humidity:  89,
			Weather: services.WeatherCondition{
				ID:          600,
				Main:        "Clouds",
				Description: "partly cloudy",
				Icon:        "02n",
			},
			Clouds: 0,
			Wind: services.WindCondition{
				Speed: 1.07,
				Deg:   234,
			},
			Rain:    0,
			Snow:    10,
			UVIndex: 0,
		},
	}

	res := averageForecast(arr)

	actual := services.MetaForecast{
		Timestamp: 1576151100000,
		Temp:      0.5,
		TempMin:   -0.5,
		TempMax:   0.5,
		Pressure:  1030,
		Humidity:  89,
		Weather: services.WeatherCondition{
			ID:          600,
			Main:        "Clear",
			Description: "clear sky",
			Icon:        "01n",
		},
		Clouds: 0,
		Wind: services.WindCondition{
			Speed: 1.07,
			Deg:   234,
		},
		Rain:    5,
		Snow:    10,
		UVIndex: 0,
	}

	assert.Equal(t, actual.Temp, res.Temp)
	assert.Equal(t, actual.TempMin, res.TempMin)
	assert.Equal(t, actual.TempMax, res.TempMax)
	assert.Equal(t, actual.Pressure, res.Pressure)
	assert.Equal(t, actual.Humidity, res.Humidity)
	assert.Equal(t, actual.Weather.ID, res.Weather.ID)
	assert.Equal(t, actual.Clouds, res.Clouds)
	assert.Equal(t, actual.Wind.Speed, res.Wind.Speed)
	assert.Equal(t, actual.Wind.Deg, res.Wind.Deg)
	assert.Equal(t, actual.Rain, res.Rain)
	assert.Equal(t, actual.Snow, res.Snow)
	assert.Equal(t, actual.UVIndex, res.UVIndex)
}

func TestFilterSameDate(t *testing.T) {
	d, _ := time.Parse(time.RFC3339, "2019-12-12T11:45:00.000Z")
	arr := []services.MetaForecast{
		services.MetaForecast{
			Timestamp: 1576151100000,
			Temp:      -1,
			TempMin:   -2,
			TempMax:   -1,
			Pressure:  1030,
			Humidity:  89,
			Weather: services.WeatherCondition{
				ID:          800,
				Main:        "Clear",
				Description: "clear sky",
				Icon:        "01n",
			},
			Clouds: 0,
			Wind: services.WindCondition{
				Speed: 1.07,
				Deg:   234,
			},
			Rain:    0,
			Snow:    0,
			UVIndex: 0,
		},
		services.MetaForecast{
			Timestamp: 1576237500000,
			Temp:      2,
			TempMin:   1,
			TempMax:   2,
			Pressure:  1030,
			Humidity:  89,
			Weather: services.WeatherCondition{
				ID:          800,
				Main:        "Clear",
				Description: "clear sky",
				Icon:        "01n",
			},
			Clouds: 0,
			Wind: services.WindCondition{
				Speed: 1.07,
				Deg:   234,
			},
			Rain:    0,
			Snow:    0,
			UVIndex: 0,
		},
	}

	res := filterSameDate(arr, &d)

	assert.Len(t, res, 1)
}

func TestSameDatePositive(t *testing.T) {
	d1, _ := time.Parse(time.RFC3339, "2019-12-10T23:57:29.928Z")
	d2, _ := time.Parse(time.RFC3339, "2019-12-10T10:51:25.513Z")

	same := sameDate(&d1, &d2)
	assert.True(t, same)
}

func TestSameDateNegative(t *testing.T) {
	d1, _ := time.Parse(time.RFC3339, "2019-12-10T23:57:29.928Z")
	d2, _ := time.Parse(time.RFC3339, "2007-12-10T10:51:25.513Z")

	same := sameDate(&d1, &d2)
	assert.False(t, same)
}
