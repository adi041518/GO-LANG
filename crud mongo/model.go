package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id"`
	Title   string             `json:"title"`
	Watched bool               `json:"watched"`
}
