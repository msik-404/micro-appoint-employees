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
	From int `bson:"from,omitempty"`
	To   int `bson:"to,omitempty"`
}

// At each day there may be many work time intervals.
type WorkTimes struct {
	Mo []TimeFrame `bson:"mo,omitempty"`
	Tu []TimeFrame `bson:"tu,omitempty"`
	We []TimeFrame `bson:"we,omitempty"`
	Th []TimeFrame `bson:"th,omitempty"`
	Fr []TimeFrame `bson:"fr,omitempty"`
	Sa []TimeFrame `bson:"sa,omitempty"`
	Su []TimeFrame `bson:"su,omitempty"`
}

// Employees not only have personal work times but also competence:
// set of services which they can perform.
type Employee struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	Name       string               `bson:"name,omitempty"`
	Surname    string               `bson:"surname,omitempty"`
	WorkTimes  WorkTimes            `bson:"work_times,omitempty"`
	Competence []primitive.ObjectID `bson:"competence,omitempty"`
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
	Name       string               `bson:"name,omitempty"`
	Surname    string               `bson:"surname,omitempty"`
	WorkTimes  *WorkTimes           `bson:"work_times,omitempty"`
	Competence []primitive.ObjectID `bson:"competence,omitempty"`
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
