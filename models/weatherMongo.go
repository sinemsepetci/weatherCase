package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type WeatherMongo struct {
	Id                    primitive.ObjectID `json:"id,omitempty"`
	Location              string             `json:"location"`
	Service_1_Temperature float64            `json:"service_1_temperature"`
	Service_2_Temperature float64            `json:"service_2_temperature"`
	Request_Count         int                `json:"request_count"`
	Created_At            primitive.DateTime `json:"created_at,omitempty"`
}
