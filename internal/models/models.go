package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoModel interface {
	InsertOne(*mongo.Database) (*mongo.InsertOneResult, error)
	UpdateOne(*mongo.Database) (*mongo.UpdateResult, error)
}

type Employee struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name,omitempty"`
	Surname string             `json:"surname" bson:"surname,omitempty"`
}

type TimeFrame struct {
	From primitive.DateTime `json:"from" bson:"from,omitempty"`
	To   primitive.DateTime `json:"to" bson:"to,omitempty"`
}

type WorkTimes struct {
	Mo []TimeFrame `json:"mo" bson:"mo,omitempty"`
	Tu []TimeFrame `json:"tu" bson:"tu,omitempty"`
	We []TimeFrame `json:"we" bson:"we,omitempty"`
	Th []TimeFrame `json:"th" bson:"th,omitempty"`
	Fr []TimeFrame `json:"fr" bson:"fr,omitempty"`
	Sa []TimeFrame `json:"sa" bson:"sa,omitempty"`
	Su []TimeFrame `json:"su" bson:"su,omitempty"`
}

type EmployeeInfo struct {
	EmployeeID primitive.ObjectID   `json:"_id" bson:"_id,omitempty"`
	WorkTimes  WorkTimes            `json:"work_times" bson:"work_times,omitempty"`
	Competence []primitive.ObjectID `json:"competence" bson:"competence,omitempty"`
}
