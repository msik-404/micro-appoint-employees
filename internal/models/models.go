package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
type Employee struct {
	ID         primitive.ObjectID   `json:"_id" bson:"_id,omitempty"`
	Name       string               `json:"name" bson:"name,omitempty"`
	Surname    string               `json:"surname" bson:"surname,omitempty"`
	WorkTimes  WorkTimes            `json:"work_times" bson:"work_times,omitempty"`
	Competence []primitive.ObjectID `json:"competence" bson:"competence,omitempty"`
}

func (employee *Employee) InsertOne(
	db *mongo.Database,
) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := db.Collection("employees")
	return coll.InsertOne(ctx, employee)
}

type EmployeeUpdate struct {
	Name       string               `json:"name" bson:"name,omitempty"`
	Surname    string               `json:"surname" bson:"surname,omitempty"`
	WorkTimes  *WorkTimes           `json:"work_times" bson:"work_times,omitempty"`
	Competence []primitive.ObjectID `json:"competence" bson:"competence,omitempty"`
}

func (employeeUpdate *EmployeeUpdate) UpdateOne(
	db *mongo.Database,
	employeeID primitive.ObjectID,
) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := db.Collection("employees")

	update := bson.M{"$set": employeeUpdate}
	return coll.UpdateByID(ctx, employeeID, update)
}

func FindOneEmployee(
	db *mongo.Database,
	employeeID primitive.ObjectID,
) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.FindOne()
	opts.SetProjection(bson.D{
		{"_id", 0},
	})

	coll := db.Collection("employees")
	filter := bson.M{"_id": employeeID}
	return coll.FindOne(ctx, filter, opts)
}

func FindManyEmployees(
	db *mongo.Database,
	startValue primitive.ObjectID,
	nPerPage int64,
) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Find()
	opts.SetSort(bson.M{"_id": -1})
	opts.SetLimit(nPerPage)
	opts.SetProjection(bson.D{
		{"work_times", 0},
		{"competence", 0},
	})

	filter := bson.M{}
	if !startValue.IsZero() {
		filter = bson.M{"_id": bson.M{"$lt": startValue}}
	}
	coll := db.Collection("employees")
	return coll.Find(ctx, filter, opts)
}

func DeleteOneEmployee(
	db *mongo.Database,
	employeeID primitive.ObjectID,
) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := db.Collection("employees")
	filter := bson.M{"_id": employeeID}
	return coll.DeleteOne(ctx, filter)
}
