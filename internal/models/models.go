package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoModel interface {
	InsertOne(*mongo.Database) (*mongo.InsertOneResult, error)
	UpdateOne(*mongo.Database) (*mongo.UpdateResult, error)
}
