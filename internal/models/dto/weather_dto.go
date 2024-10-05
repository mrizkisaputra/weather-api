package dto

type WeatherResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"resolvedAddress"`
	Timezone  string  `json:"timezone"`
	Days      []Day   `json:"days"`
}

type Day struct {
	Date          string  `json:"datetime"`
	DatetimeEpoch int64   `json:"datetimeEpoch"`
	TempMax       float64 `json:"tempMax"`
	TempMin       float64 `json:"tempMin"`
	Temp          float64 `json:"temp"`
	Hours         []Hour  `json:"hours"`
}

type Hour struct {
	Time          string  `json:"datetime"`
	DatetimeEpoch int64   `json:"datetimeEpoch"`
	Temp          float64 `json:"temp"`
	FeelsLike     float64 `json:"feelslike"`
	Humidity      float64 `json:"humidity"`
	Dew           float64 `json:"dew"`
}

type WeatherRequest struct {
	Location string
	Date1    string
	Date2    string
}
