package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (employeeInfo *EmployeeInfo) InsertOne(db *mongo.Database) (*mongo.InsertOneResult, error) {
    coll := db.Collection("employee_infos")
    return GenericInsertOne(coll, employeeInfo)
}

func (employeeInfo *EmployeeInfo ) UpdateOne(db *mongo.Database) (*mongo.UpdateResult, error) {
	coll := db.Collection("employee_infos")
	filter := bson.D{{"_id", employeeInfo.EmployeeID}}
	return GenericUpdateOne(coll, filter, employeeInfo)
}
