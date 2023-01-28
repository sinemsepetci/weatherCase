package services

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"time"
	"weatherCase/configs"
	"weatherCase/models"
	"weatherCase/repository"
)

type DefaultService struct {
	repo repository.WeatherRepository
}

type WeatherService interface {
	WeatherInsert(weather models.WeatherMongo) (bool, error)
}

func (w DefaultService) WeatherInsert(location string) models.Response {

	weatherApiResp := CallWeatherApi(location)
	weatherStackApiResp := CallWeatherStackApi(location)

	mongoModel := models.WeatherMongo{
		Id:                    primitive.ObjectID{},
		Location:              weatherApiResp.Location.Name,
		Service_1_Temperature: weatherApiResp.Current.Temp_C,
		Service_2_Temperature: weatherStackApiResp.Current.Temperature,
		Request_Count:         0,
		Created_At:            primitive.NewDateTimeFromTime(time.Now()),
	}

	result, err := w.repo.Insert(mongoModel)
	if err != nil || result == false {
		return models.Response{}
	}

	avg := (mongoModel.Service_1_Temperature + mongoModel.Service_2_Temperature) / 2

	return models.Response{Location: location, Temperature: fmt.Sprintf("%v", avg)}
}

func CallWeatherApi(location string) models.WeatherApiResponse {
	var response models.WeatherApiResponse
	//WeatherApi
	weatherApiResponse, error := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + configs.EnvWeatherApiKey() + "&q=" + location + "&days=1&aqi=no&alerts=no")
	if error != nil {
		fmt.Println(error)
	}
	body, err := io.ReadAll(weatherApiResponse.Body)
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(body)
	json.Unmarshal([]byte(bodyString), &response)
	return response
}

func CallWeatherStackApi(location string) models.WeatherStackResponse {
	var response models.WeatherStackResponse
	//WeatherStack
	weatherStackResponse, error := http.Get("http://api.weatherstack.com/current?access_key=" + configs.EnvWeatherStackApiKey() + "&query=Istanbul")
	if error != nil {
		fmt.Println(error)
	}
	body, err := io.ReadAll(weatherStackResponse.Body)
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(body)
	json.Unmarshal([]byte(bodyString), &response)
	return response
}

func NewWeatherService(weatherRepository repository.WeatherRepository) DefaultService {
	return DefaultService{weatherRepository}

}
