package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/msik-404/micro-appoint-employees/internal/database"
)

// Stores work time intervals. Time is stored as number of minutes since 00:00.
type TimeFrame struct {
	From int32 `bson:"from,omitempty"`
	To   int32 `bson:"to,omitempty"`
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
	CompanyID  primitive.ObjectID   `bson:"company_id,omitempty"`
	Name       string               `bson:"name,omitempty"`
	Surname    string               `bson:"surname,omitempty"`
	WorkTimes  WorkTimes            `bson:"work_times,omitempty"`
	Competence []primitive.ObjectID `bson:"competence,omitempty"`
}

func (employee *Employee) InsertOne(
	ctx context.Context,
	db *mongo.Database,
) (*mongo.InsertOneResult, error) {

	coll := db.Collection(database.CollName)
	return coll.InsertOne(ctx, employee)
}

type EmployeeUpdate struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	CompanyID  *primitive.ObjectID  `bson:"company_id,omitempty"`
	Name       *string              `bson:"name,omitempty"`
	Surname    *string              `bson:"surname,omitempty"`
	WorkTimes  *WorkTimes           `bson:"work_times,omitempty"`
	Competence []primitive.ObjectID `bson:"competence,omitempty"`
}

func (employeeUpdate *EmployeeUpdate) UpdateOne(
	ctx context.Context,
	db *mongo.Database,
	employeeID primitive.ObjectID,
) (*mongo.UpdateResult, error) {
	coll := db.Collection(database.CollName)

	update := bson.M{"$set": employeeUpdate}
	return coll.UpdateByID(ctx, employeeID, update)
}

func FindOneEmployee(
	ctx context.Context,
	db *mongo.Database,
	employeeID primitive.ObjectID,
) *mongo.SingleResult {
	opts := options.FindOne()
	opts.SetProjection(bson.D{
		{Key: "_id", Value: 0},
	})

	coll := db.Collection(database.CollName)
	filter := bson.M{"_id": employeeID}
	return coll.FindOne(ctx, filter, opts)
}

func FindManyEmployees(
	ctx context.Context,
	db *mongo.Database,
	companyID primitive.ObjectID,
	startValue primitive.ObjectID,
	nPerPage int64,
) (*mongo.Cursor, error) {
	opts := options.Find()
	opts.SetSort(bson.M{"_id": -1})
	opts.SetLimit(nPerPage)
	opts.SetProjection(bson.D{
		{Key: "work_times", Value: 0},
		{Key: "competence", Value: 0},
	})

	filter := bson.M{"company_id": companyID}
	if !startValue.IsZero() {
		filter = bson.M{"_id": bson.M{"$lt": startValue}}
	}
	coll := db.Collection(database.CollName)
	return coll.Find(ctx, filter, opts)
}

func DeleteOneEmployee(
	ctx context.Context,
	db *mongo.Database,
	employeeID primitive.ObjectID,
) (*mongo.DeleteResult, error) {
	coll := db.Collection(database.CollName)
	filter := bson.M{"_id": employeeID}
	return coll.DeleteOne(ctx, filter)
}
