package models

type WeatherStackResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Current struct {
	Temperature float64 `json:"temperature,omitempty"`
}
