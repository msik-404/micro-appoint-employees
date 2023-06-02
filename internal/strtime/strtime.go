package strtime

import (
	"fmt"
	"strconv"

	"github.com/msik-404/micro-appoint-employees/internal/models"
)

// String representation of TimeFrame.
// like "1630" is 16:30 is 16*60+30
type TimeFrameStr struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type WorkTimesStr struct {
	Mo []TimeFrameStr `json:"mo"`
	Tu []TimeFrameStr `json:"tu"`
	We []TimeFrameStr `json:"we"`
	Th []TimeFrameStr `json:"th"`
	Fr []TimeFrameStr `json:"fr"`
	Sa []TimeFrameStr `json:"sa"`
	Su []TimeFrameStr `json:"su"`
}

type TimeParseError struct {
    message string
}

func (e TimeParseError) Error() string {
    return e.message
}

// Transforms string representation to numeric(number of minutes since 00:00).
func strTimeToInt(stringDate string) (int, error) {
	if len(stringDate) == 0 || len(stringDate) > 4 {
		return -1, TimeParseError{"Wrong string format"}
	}
	hours, err := strconv.Atoi(stringDate[0:2])
	if err != nil {
		return hours, TimeParseError{err.Error()}
	}
	if hours < 0 || hours > 23 {
		return hours, TimeParseError{"Wrong string format"}
	}
	minutes, err := strconv.Atoi(stringDate[2:4])
	if err != nil {
		return hours, TimeParseError{err.Error()}
	}
	if minutes < 0 || minutes > 59 {
		return hours, TimeParseError{"Wrong string format"}
	}
	return hours*60 + minutes, nil
}

func intToStrTime(i int) (string, error) {
	if i < 0 || i > (23*60+59) {
		return "", TimeParseError{"Wrong int value to transform to string time"}
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

func toTimeFrameStr(timeFrame *models.TimeFrame) (TimeFrameStr, error) {
	fromTimeStr, err := intToStrTime(timeFrame.From)
	if err != nil {
		return TimeFrameStr{fromTimeStr, ""}, err 
	}
	toTimeStr, err := intToStrTime(timeFrame.To)
	return TimeFrameStr{fromTimeStr, toTimeStr}, err 
}

func toTimeFrame(timeFrameStr *TimeFrameStr) (models.TimeFrame, error) {
	fromTime, err := strTimeToInt(timeFrameStr.From)
	if err != nil {
		return models.TimeFrame{From: fromTime, To: -1}, err
	}
	toTime, err := strTimeToInt(timeFrameStr.To)
	return models.TimeFrame{From: fromTime, To: toTime}, err
}

func ToWorkTimesStr(workTimes *models.WorkTimes) (WorkTimesStr, error) {
	workTimesStr := WorkTimesStr{}
	for _, timeFrame := range workTimes.Mo {
		timeFrameStr, err := toTimeFrameStr(&timeFrame)
		if err != nil {
			return workTimesStr, err
		}
		workTimesStr.Mo = append(workTimesStr.Mo, timeFrameStr)
	}
	for _, timeFrame := range workTimes.Tu {
		timeFrameStr, err := toTimeFrameStr(&timeFrame)
		if err != nil {
			return workTimesStr, err
		}
		workTimesStr.Tu = append(workTimesStr.Tu, timeFrameStr)
	}
	for _, timeFrame := range workTimes.We {
		timeFrameStr, err := toTimeFrameStr(&timeFrame)
		if err != nil {
			return workTimesStr, err
		}
		workTimesStr.We = append(workTimesStr.We, timeFrameStr)
	}
	for _, timeFrame := range workTimes.Th {
		timeFrameStr, err := toTimeFrameStr(&timeFrame)
		if err != nil {
			return workTimesStr, err
		}
		workTimesStr.Th = append(workTimesStr.Th, timeFrameStr)
	}
	for _, timeFrame := range workTimes.Fr {
		timeFrameStr, err := toTimeFrameStr(&timeFrame)
		if err != nil {
			return workTimesStr, err
		}
		workTimesStr.Fr = append(workTimesStr.Fr, timeFrameStr)
	}
	for _, timeFrame := range workTimes.Sa {
		timeFrameStr, err := toTimeFrameStr(&timeFrame)
		if err != nil {
			return workTimesStr, err
		}
		workTimesStr.Sa = append(workTimesStr.Sa, timeFrameStr)
	}
	for _, timeFrame := range workTimes.Su {
		timeFrameStr, err := toTimeFrameStr(&timeFrame)
		if err != nil {
			return workTimesStr, err
		}
		workTimesStr.Su = append(workTimesStr.Su, timeFrameStr)
	}
	return workTimesStr, nil
}

func ToWorkTimes(workTimesStr *WorkTimesStr) (models.WorkTimes, error) {
	workTimes := models.WorkTimes{}
	for _, timeFrameStr := range workTimesStr.Mo {
		timeFrame, err := toTimeFrame(&timeFrameStr)
		if err != nil {
			return workTimes, err
		}
		workTimes.Mo = append(workTimes.Mo, timeFrame)
	}
	for _, timeFrameStr := range workTimesStr.Tu {
		timeFrame, err := toTimeFrame(&timeFrameStr)
		if err != nil {
			return workTimes, err
		}
		workTimes.Tu = append(workTimes.Tu, timeFrame)
	}
	for _, timeFrameStr := range workTimesStr.We {
		timeFrame, err := toTimeFrame(&timeFrameStr)
		if err != nil {
			return workTimes, err
		}
		workTimes.We = append(workTimes.We, timeFrame)
	}
	for _, timeFrameStr := range workTimesStr.Th {
		timeFrame, err := toTimeFrame(&timeFrameStr)
		if err != nil {
			return workTimes, err
		}
		workTimes.Th = append(workTimes.Th, timeFrame)
	}
	for _, timeFrameStr := range workTimesStr.Fr {
		timeFrame, err := toTimeFrame(&timeFrameStr)
		if err != nil {
			return workTimes, err
		}
		workTimes.Fr = append(workTimes.Fr, timeFrame)
	}
	for _, timeFrameStr := range workTimesStr.Sa {
		timeFrame, err := toTimeFrame(&timeFrameStr)
		if err != nil {
			return workTimes, err
		}
		workTimes.Sa = append(workTimes.Sa, timeFrame)
	}
	for _, timeFrameStr := range workTimesStr.Su {
		timeFrame, err := toTimeFrame(&timeFrameStr)
		if err != nil {
			return workTimes, err
		}
		workTimes.Su = append(workTimes.Su, timeFrame)
	}
	return workTimes, nil
}
