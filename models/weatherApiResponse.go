package models

type WeatherApiResponse struct {
	Location Location    `json:"location"`
	Current  CurrentTemp `json:"current"`
}

type CurrentTemp struct {
	Temp_C float64 `json:"temp_c,omitempty"`
}

type Location struct {
	Name string `json:"name,omitempty"`
}
