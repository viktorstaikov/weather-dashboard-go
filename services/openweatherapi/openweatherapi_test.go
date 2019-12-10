package openweatherapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/viktorstaikov/weather-dashboard-go/config"
)

func TestParseForecastResponse(t *testing.T) {
	jsonString := `{"list":[{"dt":1575579600,"main":{"temp":-1,"temp_min":-3,"temp_max":0,"pressure":1031,"sea_level":1031,"grnd_level":913,"humidity":89,"temp_kf":1.58},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.39,"deg":233},"rain":{"3h":0},"snow":{"3h":0},"sys":{"pod":"n"},"dt_txt":"2019-12-05 21:00:00"},{"dt":1575590400,"main":{"temp":3,"temp_min":1,"temp_max":5,"pressure":1030,"sea_level":1030,"grnd_level":913,"humidity":89,"temp_kf":1.19},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.07,"deg":234},"rain":{"3h":20},"snow":{"3h":10},"sys":{"pod":"n"},"dt_txt":"2019-12-06 00:00:00"}]}`

	var forecast ForecastResponse
	json.Unmarshal([]byte(jsonString), &forecast)

	metaForecastSlice := parseForecastResponse(&forecast)

	assert.Equal(t, 2, len(metaForecastSlice))
}

func TestmakeForecastRequest(t *testing.T) {
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

	api := MakeOpenWeatherAPI()

	_, err := api.makeForecastRequest()
	if err != nil {
		t.Errorf("makeForecastRequest() returned an error: %s", err)
	}

}
