package communication

import (
	"golang.org/x/exp/constraints"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func verifyString(value *string, maxLength int) (*string, error) {
	if value != nil {
		if len(*value) > int(maxLength) {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Value should be shorter than %d",
				maxLength,
			)
		}
	}
	return value, nil
}

func verifyInteger[T constraints.Integer](value *T, low T, high T) (*T, error) {
	if value != nil {
		if *value > high || *value <= low {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Value should be smaller than %d and greater than %d",
				high,
				low,
			)
		}
	}
	return value, nil
}

func verifyTimeFrame(timeFrame *TimeFrame) (*TimeFrame, error) {
	if timeFrame != nil {
		if *timeFrame.From < 0 ||
			*timeFrame.To < 0 ||
			*timeFrame.From > 23*60+59 ||
			*timeFrame.To > 23*60+59 {
			return nil, status.Error(
				codes.InvalidArgument,
				"Value should be smaller than 1439 and greater than 0",
			)
		}
		if *timeFrame.From >= *timeFrame.To {
			return nil, status.Error(
				codes.InvalidArgument,
				"From value should be smaller than To value",
			)
		}
	}
	return timeFrame, nil
}

func verifyWorkTimes(workTimes *WorkTimes) (*WorkTimes, error) {
	if workTimes != nil {
		for _, timeFrame := range workTimes.Mo {
			if _, err := verifyTimeFrame(timeFrame); err != nil {
				return nil, err
			}
		}
		for _, timeFrame := range workTimes.Tu {
			if _, err := verifyTimeFrame(timeFrame); err != nil {
				return nil, err
			}
		}
		for _, timeFrame := range workTimes.We {
			if _, err := verifyTimeFrame(timeFrame); err != nil {
				return nil, err
			}
		}
		for _, timeFrame := range workTimes.Th {
			if _, err := verifyTimeFrame(timeFrame); err != nil {
				return nil, err
			}
		}
		for _, timeFrame := range workTimes.Fr {
			if _, err := verifyTimeFrame(timeFrame); err != nil {
				return nil, err
			}
		}
		for _, timeFrame := range workTimes.Sa {
			if _, err := verifyTimeFrame(timeFrame); err != nil {
				return nil, err
			}
		}
		for _, timeFrame := range workTimes.Su {
			if _, err := verifyTimeFrame(timeFrame); err != nil {
				return nil, err
			}
		}
	}
	return workTimes, nil
}

func verifyCompetence(competence []string) ([]string, error) {
	if len(competence) != 0 {
		for _, hex := range competence {
			if len(hex) != 24 {
				return nil, status.Error(
					codes.InvalidArgument,
					"This is not proper hex value for objectID",
				)
			}
		}
	}
	return competence, nil
}
