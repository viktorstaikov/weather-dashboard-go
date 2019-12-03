package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WeatherController ...
type WeatherController struct{}

// TempSeries ...
func (h WeatherController) TempSeries(c *gin.Context) {
	c.String(http.StatusOK, "Temp series")
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
