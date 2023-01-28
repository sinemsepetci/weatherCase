package main

import (
	"github.com/gofiber/fiber/v2"
	"weatherCase/app"
	"weatherCase/configs"
	"weatherCase/repository"
	"weatherCase/services"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()
	dbClient := configs.GetCollection(configs.DB, "weather_queries")

	weatherRepository := repository.NewModelRepositoryDb(dbClient)
	weatherHandler := app.WeatherController{Service: services.NewWeatherService(weatherRepository)}

	appRoute.Get("/weather/:location", weatherHandler.GetWeather)

	appRoute.Listen(":8080")
}
