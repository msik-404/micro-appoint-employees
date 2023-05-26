package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (employee *Employee) InsertOne(db *mongo.Database) (*mongo.InsertOneResult, error) {
	coll := db.Collection("employees")
	return GenericInsertOne(coll, employee)
}

func (employee *Employee) UpdateOne(db *mongo.Database) (*mongo.UpdateResult, error) {
	coll := db.Collection("employees")
	filter := bson.D{{"_id", employee.ID}}
	return GenericUpdateOne(coll, filter, employee)
}

type EmployeeCombRepr struct {
	ID         primitive.ObjectID   `json:"_id" bson:"_id,omitempty"`
	Name       string               `json:"name" bson:"name,omitempty"`
	Surname    string               `json:"surname" bson:"surname,omitempty"`
	WorkTimes  WorkTimes            `json:"work_times" bson:"work_times,omitempty"`
	Competence []primitive.ObjectID `json:"competence" bson:"competence,omitempty"`
}

func (employeeCombRepr *EmployeeCombRepr) InsertCombRepr(db *mongo.Database) ([]*mongo.InsertOneResult, error) {
	employee := Employee{
		Name:    employeeCombRepr.Name,
		Surname: employeeCombRepr.Surname,
	}
	var results []*mongo.InsertOneResult
	result, err := employee.InsertOne(db)
	if err != nil {
		return nil, err
	}
	results = append(results, result)
	employee_info := EmployeeInfo{
		WorkTimes:  employeeCombRepr.WorkTimes,
		Competence: employeeCombRepr.Competence,
	}
	result, err = employee_info.InsertOne(db)
	if err != nil {
		return nil, err
	}
	results = append(results, result)
	return results, nil
}
