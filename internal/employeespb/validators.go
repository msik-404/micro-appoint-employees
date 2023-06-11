package employeespb 

import (
	"golang.org/x/exp/constraints"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func verifyString(value *string, maxLength int) error {
	if value != nil {
		if len(*value) > int(maxLength) {
			return status.Errorf(
				codes.InvalidArgument,
				"Value should be shorter than %d",
				maxLength,
			)
		}
	}
	return nil
}

func verifyInteger[T constraints.Integer](value *T, low T, high T) error {
	if value != nil {
		if *value > high || *value <= low {
			return status.Errorf(
				codes.InvalidArgument,
				"Value should be smaller than %d and greater than %d",
				high,
				low,
			)
		}
	}
	return nil
}

func verifyTimeFrame(timeFrame *TimeFrame) error {
	if timeFrame != nil {
		if *timeFrame.From < 0 ||
			*timeFrame.To < 0 ||
			*timeFrame.From > 23*60+59 ||
			*timeFrame.To > 23*60+59 {
			return status.Error(
				codes.InvalidArgument,
				"Value should be smaller than 1439 and greater than 0",
			)
		}
		if *timeFrame.From >= *timeFrame.To {
			return status.Error(
				codes.InvalidArgument,
				"From value should be smaller than To value",
			)
		}
	}
	return nil
}

func verifyWorkTimes(workTimes *WorkTimes) error {
	if workTimes != nil {
		for _, timeFrame := range workTimes.Mo {
			if err := verifyTimeFrame(timeFrame); err != nil {
				return err
			}
		}
		for _, timeFrame := range workTimes.Tu {
			if err := verifyTimeFrame(timeFrame); err != nil {
				return err
			}
		}
		for _, timeFrame := range workTimes.We {
			if err := verifyTimeFrame(timeFrame); err != nil {
				return err
			}
		}
		for _, timeFrame := range workTimes.Th {
			if err := verifyTimeFrame(timeFrame); err != nil {
				return err
			}
		}
		for _, timeFrame := range workTimes.Fr {
			if err := verifyTimeFrame(timeFrame); err != nil {
				return err
			}
		}
		for _, timeFrame := range workTimes.Sa {
			if err := verifyTimeFrame(timeFrame); err != nil {
				return  err
			}
		}
		for _, timeFrame := range workTimes.Su {
			if err := verifyTimeFrame(timeFrame); err != nil {
				return err
			}
		}
	}
	return nil
}

func verifyCompetence(competence []string) error {
	if len(competence) != 0 {
		for _, hex := range competence {
			if len(hex) != 24 {
				return status.Error(
					codes.InvalidArgument,
					"This is not proper hex value for objectID",
				)
			}
		}
	}
	return nil
}
