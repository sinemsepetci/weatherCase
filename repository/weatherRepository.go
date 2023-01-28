package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"weatherCase/models"
)

type WeatherRepository struct {
	weather_queries *mongo.Collection
}

type repository interface {
	Insert(weatherMongo models.WeatherMongo) (bool, error)
}

func (r WeatherRepository) Insert(model models.WeatherMongo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.weather_queries.InsertOne(ctx, model)

	if result.InsertedID == nil || err != nil {
		errors.New("failed add")
		return false, err
	}
	return true, nil
}

func NewModelRepositoryDb(dbClient *mongo.Collection) WeatherRepository {
	return WeatherRepository{dbClient}
}
