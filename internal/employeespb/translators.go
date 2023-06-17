package employeespb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/msik-404/micro-appoint-employees/internal/models"
)

func timeFrameGRPCToModel(timeFrame *TimeFrame) models.TimeFrame {
	return models.TimeFrame{
		From: timeFrame.GetFrom(),
		To:   timeFrame.GetTo(),
	}
}

func workTimesGRPCToModel(workTimes *WorkTimes) models.WorkTimes {
	modelWorkTimes := models.WorkTimes{}
	for _, timeFrame := range workTimes.Mo {
		modelTimeFrame := timeFrameGRPCToModel(timeFrame)
		modelWorkTimes.Mo = append(modelWorkTimes.Mo, modelTimeFrame)
	}
	for _, timeFrame := range workTimes.Tu {
		modelTimeFrame := timeFrameGRPCToModel(timeFrame)
		modelWorkTimes.Tu = append(modelWorkTimes.Tu, modelTimeFrame)
	}
	for _, timeFrame := range workTimes.We {
		modelTimeFrame := timeFrameGRPCToModel(timeFrame)
		modelWorkTimes.We = append(modelWorkTimes.We, modelTimeFrame)
	}
	for _, timeFrame := range workTimes.Th {
		modelTimeFrame := timeFrameGRPCToModel(timeFrame)
		modelWorkTimes.Th = append(modelWorkTimes.Th, modelTimeFrame)
	}
	for _, timeFrame := range workTimes.Fr {
		modelTimeFrame := timeFrameGRPCToModel(timeFrame)
		modelWorkTimes.Fr = append(modelWorkTimes.Fr, modelTimeFrame)
	}
	for _, timeFrame := range workTimes.Sa {
		modelTimeFrame := timeFrameGRPCToModel(timeFrame)
		modelWorkTimes.Sa = append(modelWorkTimes.Sa, modelTimeFrame)
	}
	for _, timeFrame := range workTimes.Su {
		modelTimeFrame := timeFrameGRPCToModel(timeFrame)
		modelWorkTimes.Su = append(modelWorkTimes.Su, modelTimeFrame)
	}
	return modelWorkTimes
}

func hexToObjectID(hex string) (primitive.ObjectID, error) {
	ID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return primitive.NilObjectID, status.Error(codes.InvalidArgument, err.Error())
	}
	return ID, nil
}

func timeFrameModelToGRPC(timeFrame *models.TimeFrame) TimeFrame {
	return TimeFrame{
		From: &timeFrame.From,
		To:   &timeFrame.To,
	}
}

func workTimesModelToGRPC(workTimes *models.WorkTimes) *WorkTimes {
	grpcWorkTimes := WorkTimes{}
	for _, timeFrame := range workTimes.Mo {
		grpcTimeFrame := timeFrameModelToGRPC(&timeFrame)
		grpcWorkTimes.Mo = append(grpcWorkTimes.Mo, &grpcTimeFrame)
	}
	for _, timeFrame := range workTimes.Tu {
		grpcTimeFrame := timeFrameModelToGRPC(&timeFrame)
		grpcWorkTimes.Tu = append(grpcWorkTimes.Tu, &grpcTimeFrame)
	}
	for _, timeFrame := range workTimes.We {
		grpcTimeFrame := timeFrameModelToGRPC(&timeFrame)
		grpcWorkTimes.We = append(grpcWorkTimes.We, &grpcTimeFrame)
	}
	for _, timeFrame := range workTimes.Th {
		grpcTimeFrame := timeFrameModelToGRPC(&timeFrame)
		grpcWorkTimes.Th = append(grpcWorkTimes.Th, &grpcTimeFrame)
	}
	for _, timeFrame := range workTimes.Fr {
		grpcTimeFrame := timeFrameModelToGRPC(&timeFrame)
		grpcWorkTimes.Fr = append(grpcWorkTimes.Fr, &grpcTimeFrame)
	}
	for _, timeFrame := range workTimes.Sa {
		grpcTimeFrame := timeFrameModelToGRPC(&timeFrame)
		grpcWorkTimes.Sa = append(grpcWorkTimes.Sa, &grpcTimeFrame)
	}
	for _, timeFrame := range workTimes.Su {
		grpcTimeFrame := timeFrameModelToGRPC(&timeFrame)
		grpcWorkTimes.Su = append(grpcWorkTimes.Su, &grpcTimeFrame)
	}
	return &grpcWorkTimes
}

func IntToDay(intDay int32) (string, error) {
	switch intDay {
	case 0:
		return "mo", nil
	case 1:
		return "tu", nil
	case 2:
		return "we", nil
	case 3:
		return "th", nil
	case 4:
		return "fr", nil
	case 5:
		return "sa", nil
	case 6:
		return "su", nil
	default:
		return "", status.Error(
			codes.InvalidArgument,
			"Integer representing day is invalid, should be in range of 0-6",
		)
	}
}

func IntToTimeFrame(
	employee *models.Employee,
	intDay int32,
) ([]models.TimeFrame, error) {
	switch intDay {
	case 0:
		return employee.WorkTimes.Mo, nil
	case 1:
		return employee.WorkTimes.Tu, nil
	case 2:
		return employee.WorkTimes.We, nil
	case 3:
		return employee.WorkTimes.Th, nil
	case 4:
		return employee.WorkTimes.Fr, nil
	case 5:
		return employee.WorkTimes.Sa, nil
	case 6:
		return employee.WorkTimes.Su, nil
	default:
		return nil, status.Error(
			codes.InvalidArgument,
			"Integer representing day is invalid, should be in range of 0-6",
		)
	}
}
