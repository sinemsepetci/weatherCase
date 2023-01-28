package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoUri() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error .env")
	}
	mongoUri := os.Getenv("MONGOURI")
	return mongoUri
}
func EnvWeatherApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error .env")
	}
	mongoUri := os.Getenv("WEATHERAPIKEY")
	return mongoUri
}
func EnvWeatherStackApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error .env")
	}
	mongoUri := os.Getenv("WEATHERSTACKAPIKEY")
	return mongoUri
}
