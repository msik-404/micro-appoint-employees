package models

import (
	"errors"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (employeeInfo *EmployeeInfo) InsertOne(db *mongo.Database) (*mongo.InsertOneResult, error) {
	coll := db.Collection("employee_infos")
	return GenericInsertOne(coll, employeeInfo)
}

func (employeeInfo *EmployeeInfo) UpdateOne(db *mongo.Database) (*mongo.UpdateResult, error) {
	coll := db.Collection("employee_infos")
	filter := bson.D{{"_id", employeeInfo.EmployeeID}}
	return GenericUpdateOne(coll, filter, employeeInfo)
}

func intToStringTime(i int) (string, error) {
	if i < 0 || i > (23*60+59) {
		return "", errors.New("Wrong int value to transform to string time")
	}
	hoursInt := i / 60
	minutesInt := i % 60
	hours := strconv.Itoa(hoursInt)
	minutes := strconv.Itoa(minutesInt)
	// if is single digit add 0 prefix
	if hoursInt < 10 {
		hours = fmt.Sprintf("0%s", hours)
	}
	if minutesInt < 10 {
		minutes = fmt.Sprintf("0%s", minutes)
	}
	return fmt.Sprintf("%s%s", hours, minutes), nil
}

func (timeFrame *TimeFrame) toTimeFrameString() (TimeFrameString, error) {
	fromTimeString, err := intToStringTime(timeFrame.From)
	if err != nil {
		return TimeFrameString{fromTimeString, ""}, err
	}
	toTimeString, err := intToStringTime(timeFrame.To)
	return TimeFrameString{fromTimeString, toTimeString}, err
}

func (workTimes *WorkTimes) toWorkTimesString() (WorkTimesString, error) {
	workTimesString := WorkTimesString{}
	for _, timeFrame := range workTimes.Mo {
		timeFrameString, err := timeFrame.toTimeFrameString()
		if err != nil {
			return workTimesString, err
		}
		workTimesString.Mo = append(workTimesString.Mo, timeFrameString)
	}
	for _, timeFrame := range workTimes.Tu {
		timeFrameString, err := timeFrame.toTimeFrameString()
		if err != nil {
			return workTimesString, err
		}
		workTimesString.Tu = append(workTimesString.Tu, timeFrameString)
	}
	for _, timeFrame := range workTimes.We {
		timeFrameString, err := timeFrame.toTimeFrameString()
		if err != nil {
			return workTimesString, err
		}
		workTimesString.We = append(workTimesString.We, timeFrameString)
	}
	for _, timeFrame := range workTimes.Th {
		timeFrameString, err := timeFrame.toTimeFrameString()
		if err != nil {
			return workTimesString, err
		}
		workTimesString.Th = append(workTimesString.Th, timeFrameString)
	}
	for _, timeFrame := range workTimes.Fr {
		timeFrameString, err := timeFrame.toTimeFrameString()
		if err != nil {
			return workTimesString, err
		}
		workTimesString.Fr = append(workTimesString.Fr, timeFrameString)
	}
	for _, timeFrame := range workTimes.Sa {
		timeFrameString, err := timeFrame.toTimeFrameString()
		if err != nil {
			return workTimesString, err
		}
		workTimesString.Sa = append(workTimesString.Sa, timeFrameString)
	}
	for _, timeFrame := range workTimes.Su {
		timeFrameString, err := timeFrame.toTimeFrameString()
		if err != nil {
			return workTimesString, err
		}
		workTimesString.Su = append(workTimesString.Su, timeFrameString)
	}
	return workTimesString, nil
}

type EmployeeInfoRepr struct {
	WorkTimes  WorkTimesString      `json:"work_times"`
	Competence []primitive.ObjectID `json:"competence"`
}

func (employeeInfo *EmployeeInfo) ToEmployeeInfoRepr() (EmployeeInfoRepr, error) {
	worktimes, err := employeeInfo.WorkTimes.toWorkTimesString()
	if err != nil {
		return EmployeeInfoRepr{}, err
	}
	return EmployeeInfoRepr{worktimes, employeeInfo.Competence}, nil
}
