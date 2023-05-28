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

// Stores work time intervals. Time is stored as number of minutes since 00:00.
type TimeFrame struct {
	From int `json:"from" bson:"from,omitempty"`
	To   int `json:"to" bson:"to,omitempty"`
}

// At each day there may be many work time intervals.
type WorkTimes struct {
	Mo []TimeFrame `json:"mo" bson:"mo,omitempty"`
	Tu []TimeFrame `json:"tu" bson:"tu,omitempty"`
	We []TimeFrame `json:"we" bson:"we,omitempty"`
	Th []TimeFrame `json:"th" bson:"th,omitempty"`
	Fr []TimeFrame `json:"fr" bson:"fr,omitempty"`
	Sa []TimeFrame `json:"sa" bson:"sa,omitempty"`
	Su []TimeFrame `json:"su" bson:"su,omitempty"`
}

// Employees not only have personal work times but also competence:
// set of services which they can perform.
type EmployeeInfo struct {
	EmployeeID primitive.ObjectID   `json:"_id" bson:"_id,omitempty"`
	WorkTimes  *WorkTimes            `json:"work_times" bson:"work_times,omitempty"`
	Competence []primitive.ObjectID `json:"competence" bson:"competence,omitempty"`
}
