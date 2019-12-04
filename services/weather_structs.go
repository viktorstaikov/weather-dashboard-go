package services

type forecastResponse struct {
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

type metaForecast struct {
	Timestamp uint    `json:"timestamp"`
	Temp      float64 `json:"temp"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	Weather   struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Clouds int `json:"clouds"`
	Wind   struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Rain    float64 `json:"rain"`
	Snow    float64 `json:"snow"`
	UVIndex float64 `json:"uvindex"`
}

func (r forecastResponse) toMetaForecast() []metaForecast {
	var list []metaForecast
	for _, item := range r.List {
		var metaItem metaForecast
		metaItem.Timestamp = item.Dt * 1000
		metaItem.Temp = item.Main.Temp
		metaItem.TempMin = item.Main.TempMin
		metaItem.TempMax = item.Main.TempMax
		metaItem.Pressure = item.Main.Pressure
		metaItem.Humidity = item.Main.Humidity
		metaItem.Weather = item.Weather[0]
		metaItem.Wind = item.Wind

		metaItem.Clouds = 0
		if item.Clouds.All > 0 {
			metaItem.Clouds = item.Clouds.All
		}

		metaItem.Rain = 0
		if item.Rain.ThreeH > 0 {
			metaItem.Rain = item.Rain.ThreeH
		}

		metaItem.Snow = 0
		if item.Snow.ThreeH > 0 {
			metaItem.Snow = item.Snow.ThreeH
		}

		list = append(list, metaItem)
	}
	return list
}

// TempData represents temperature forecast at given time
type TempData struct {
	Timestamp uint    `json:"timestamp"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
}

func (m metaForecast) toTempData() TempData {
	var item TempData
	item.TempMax = m.TempMax
	item.TempMin = m.TempMin
	item.Timestamp = m.Timestamp
	return item
}
