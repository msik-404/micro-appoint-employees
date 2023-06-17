package employeespb 

import (
	"math"

	"github.com/msik-404/micro-appoint-employees/internal/models"
)

// Employee working TimeFrame's are broken into units with that duration.
const TimeSlotLength int32 = 10
// Breke between services
const BrakeDuration int32 = 15

func GetServiceSlots(
	workTimeFrame *models.TimeFrame,
	serviceDuration int32,
	timeSlotLength int32,
    brakeDuration int32,
) []*TimeFrame {
	serviceSlotLength := int32(math.Ceil(
		float64(serviceDuration) / float64(timeSlotLength),
	)) * timeSlotLength
	var serviceSlots []*TimeFrame
    // Employee should have some time between providing services.
    // This includes preparation for initial service in given workTimeFrame.
    startTime := workTimeFrame.From + brakeDuration
    endTime := workTimeFrame.To - serviceSlotLength - brakeDuration 
	for i := startTime; i <= endTime; i += serviceSlotLength + brakeDuration {
        start := i
        end := start + serviceSlotLength 
		serviceSlot := TimeFrame{
			From: &start,
			To:   &end,
		}
		serviceSlots = append(serviceSlots, &serviceSlot)
	}
	return serviceSlots
}
