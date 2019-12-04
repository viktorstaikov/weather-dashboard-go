package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viktorstaikov/weather-dashboard-go/services"
)

// WeatherController ...
type WeatherController struct{}

// TempSeries ...
func (h WeatherController) TempSeries(c *gin.Context) {
	resp, err := services.GetTempSeries()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get Temprature series", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, resp)

}

// RainSeries ...
func (h WeatherController) RainSeries(c *gin.Context) {
	c.String(http.StatusOK, "Rain series")
}

// PressureSeries ...
func (h WeatherController) PressureSeries(c *gin.Context) {
	c.String(http.StatusOK, "Pressure series")
}

// HumiditySeries ...
func (h WeatherController) HumiditySeries(c *gin.Context) {
	c.String(http.StatusOK, "Humidity series")
}

// Forecast ...
func (h WeatherController) Forecast(c *gin.Context) {
	c.String(http.StatusOK, "Forecast")
}
