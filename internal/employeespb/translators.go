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
