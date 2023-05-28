package models

import (
	"errors"
	"strconv"

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

// String representation of TimeFrame.
// like "1630" is 16:30 is 16*60+30
type TimeFrameString struct {
	From string `json:"from" bson:"from,omitempty"`
	To   string `json:"to" bson:"to,omitempty"`
}

// Transforms string representation to numeric(number of minutes since 00:00).
func stringTimeToInt(stringDate string) (int, error) {
	if len(stringDate) == 0 || len(stringDate) > 4 {
		return -1, errors.New("Wrong string format")
	}
	hours, err := strconv.Atoi(stringDate[0:2])
	if err != nil {
		return hours, err
	}
	if hours < 0 || hours > 23 {
		return hours, errors.New("Wrong string fromat")
	}
	minutes, err := strconv.Atoi(stringDate[2:4])
	if err != nil {
		return hours, err
	}
	if minutes < 0 || minutes > 59 {
		return hours, errors.New("Wrong string fromat")
	}
	return hours*60 + minutes, err
}

func (timeFrameString *TimeFrameString) toTimeFrame() (TimeFrame, error) {
	fromTime, err := stringTimeToInt(timeFrameString.From)
	if err != nil {
		return TimeFrame{fromTime, -1}, err
	}
	toTime, err := stringTimeToInt(timeFrameString.To)
	return TimeFrame{fromTime, toTime}, err
}

type WorkTimesString struct {
	Mo []TimeFrameString `json:"mo" bson:"mo,omitempty"`
	Tu []TimeFrameString `json:"tu" bson:"tu,omitempty"`
	We []TimeFrameString `json:"we" bson:"we,omitempty"`
	Th []TimeFrameString `json:"th" bson:"th,omitempty"`
	Fr []TimeFrameString `json:"fr" bson:"fr,omitempty"`
	Sa []TimeFrameString `json:"sa" bson:"sa,omitempty"`
	Su []TimeFrameString `json:"su" bson:"su,omitempty"`
}

func (workTimesString *WorkTimesString) toWorkTimes() (*WorkTimes, error) {
	workTimes := WorkTimes{}
	isEmpty := true
	for _, timeFrameString := range workTimesString.Mo {
		isEmpty = false
		timeFrame, err := timeFrameString.toTimeFrame()
		if err != nil {
			return nil, err
		}
		workTimes.Mo = append(workTimes.Mo, timeFrame)
	}
	for _, timeFrameString := range workTimesString.Tu {
		isEmpty = false
		timeFrame, err := timeFrameString.toTimeFrame()
		if err != nil {
			return nil, err
		}
		workTimes.Tu = append(workTimes.Tu, timeFrame)
	}
	for _, timeFrameString := range workTimesString.We {
		isEmpty = false
		timeFrame, err := timeFrameString.toTimeFrame()
		if err != nil {
			return nil, err
		}
		workTimes.We = append(workTimes.We, timeFrame)
	}
	for _, timeFrameString := range workTimesString.Th {
		isEmpty = false
		timeFrame, err := timeFrameString.toTimeFrame()
		if err != nil {
			return nil, err
		}
		workTimes.Th = append(workTimes.Th, timeFrame)
	}
	for _, timeFrameString := range workTimesString.Fr {
		isEmpty = false
		timeFrame, err := timeFrameString.toTimeFrame()
		if err != nil {
			return nil, err
		}
		workTimes.Fr = append(workTimes.Fr, timeFrame)
	}
	for _, timeFrameString := range workTimesString.Sa {
		isEmpty = false
		timeFrame, err := timeFrameString.toTimeFrame()
		if err != nil {
			return nil, err
		}
		workTimes.Sa = append(workTimes.Sa, timeFrame)
	}
	for _, timeFrameString := range workTimesString.Su {
		isEmpty = false
		timeFrame, err := timeFrameString.toTimeFrame()
		if err != nil {
			return nil, err
		}
		workTimes.Su = append(workTimes.Su, timeFrame)
	}
    if isEmpty {
        return nil, nil
    }
	return &workTimes, nil
}

// Data representation which will be used for comunication with the front.
// Here WorkTime is represented as strings like "10:30".
type EmployeeCombRepr struct {
	ID              primitive.ObjectID   `json:"_id" bson:"_id,omitempty"`
	Name            string               `json:"name" bson:"name,omitempty"`
	Surname         string               `json:"surname" bson:"surname,omitempty"`
	WorkTimesString WorkTimesString      `json:"work_times" bson:"work_times,omitempty"`
	Competence      []primitive.ObjectID `json:"competence" bson:"competence,omitempty"`
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
	workTimes, err := employeeCombRepr.WorkTimesString.toWorkTimes()
	if err != nil {
		return nil, err
	}
	employee_info := EmployeeInfo{
		EmployeeID: result.InsertedID.(primitive.ObjectID),
		WorkTimes:  workTimes,
		Competence: employeeCombRepr.Competence,
	}
	result, err = employee_info.InsertOne(db)
	if err != nil {
		return nil, err
	}
	results = append(results, result)
	return results, nil
}

func (employeeCombRepr *EmployeeCombRepr) UpdateCombRepr(db *mongo.Database) ([]*mongo.UpdateResult, error) {
	employee := Employee{
		ID:      employeeCombRepr.ID,
		Name:    employeeCombRepr.Name,
		Surname: employeeCombRepr.Surname,
	}
	var results []*mongo.UpdateResult
	result, err := employee.UpdateOne(db)
	if err != nil {
		return nil, err
	}
	results = append(results, result)
	workTimes, err := employeeCombRepr.WorkTimesString.toWorkTimes()
	if err != nil {
		return nil, err
	}
	employee_info := EmployeeInfo{
		EmployeeID: employeeCombRepr.ID,
		WorkTimes:  workTimes,
		Competence: employeeCombRepr.Competence,
	}
	result, err = employee_info.UpdateOne(db)
	if err != nil {
		return nil, err
	}
	results = append(results, result)
	return results, nil
}
