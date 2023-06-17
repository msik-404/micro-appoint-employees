package employeespb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/msik-404/micro-appoint-employees/internal/database"
	"github.com/msik-404/micro-appoint-employees/internal/models"
)

type Server struct {
	UnimplementedApiServer
	Client mongo.Client
}

func (s *Server) AddEmployee(
	ctx context.Context,
	request *AddEmployeeRequest,
) (*emptypb.Empty, error) {
	companyID, err := hexToObjectID(request.GetCompanyId())
	if err != nil {
		return nil, err
	}
	err = verifyString(request.Name, 30)
	if err != nil {
		return nil, err
	}
	err = verifyString(request.Surname, 30)
	if err != nil {
		return nil, err
	}
	err = verifyWorkTimes(request.WorkTimes)
	if err != nil {
		return nil, err
	}
	err = verifyCompetence(request.Competence)
	if err != nil {
		return nil, err
	}
	var competence []primitive.ObjectID
	for _, hex := range request.GetCompetence() {
		serviceID, err := hexToObjectID(hex)
		if err != nil {
			return nil, err
		}
		competence = append(competence, serviceID)
	}
	var workTimesModel models.WorkTimes
	if request.WorkTimes != nil {
		workTimesModel = workTimesGRPCToModel(request.GetWorkTimes())
	}
	newEmployee := models.Employee{
		CompanyID:  companyID,
		Name:       request.GetName(),
		Surname:    request.GetSurname(),
		WorkTimes:  workTimesModel,
		Competence: competence,
	}
	db := s.Client.Database(database.DBName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = newEmployee.InsertOne(ctx, db)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) UpdateEmployee(
	ctx context.Context,
	request *UpdateEmployeeRequest,
) (*emptypb.Empty, error) {
	employeeID, err := hexToObjectID(request.GetId())
	if err != nil {
		return nil, err
	}
	companyID, err := hexToObjectID(request.GetCompanyId())
	if err != nil {
		return nil, err
	}
	err = verifyString(request.Name, 30)
	if err != nil {
		return nil, err
	}
	err = verifyString(request.Surname, 30)
	if err != nil {
		return nil, err
	}
	err = verifyWorkTimes(request.WorkTimes)
	if err != nil {
		return nil, err
	}
	err = verifyCompetence(request.Competence)
	if err != nil {
		return nil, err
	}
	var competence []primitive.ObjectID
	for _, hex := range request.GetCompetence() {
		serviceID, err := hexToObjectID(hex)
		if err != nil {
			return nil, err
		}
		competence = append(competence, serviceID)
	}
	var workTimesModel *models.WorkTimes
	if request.WorkTimes != nil {
		workTimes := workTimesGRPCToModel(request.GetWorkTimes())
		workTimesModel = &workTimes
	}
	employeeUpdate := models.EmployeeUpdate{
		Name:       request.Name,
		Surname:    request.Surname,
		WorkTimes:  workTimesModel,
		Competence: competence,
	}
	db := s.Client.Database(database.DBName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := employeeUpdate.UpdateOne(ctx, db, companyID, employeeID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if result.MatchedCount == 0 {
		return nil, status.Error(
			codes.NotFound,
			"Employee with that companyID and EmployeeID was not found",
		)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteEmployee(
	ctx context.Context,
	request *DeleteEmployeeRequest,
) (*emptypb.Empty, error) {
	employeeID, err := hexToObjectID(request.GetId())
	if err != nil {
		return nil, err
	}
	companyID, err := hexToObjectID(request.GetCompanyId())
	if err != nil {
		return nil, err
	}
	db := s.Client.Database(database.DBName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := models.DeleteOneEmployee(ctx, db, companyID, employeeID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if result.DeletedCount == 0 {
		return nil, status.Error(
			codes.NotFound,
			"Employee with that companyID and EmployeeID was not found",
		)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) FindOneEmployee(
	ctx context.Context,
	request *EmployeeRequest,
) (*EmployeeReply, error) {
	employeeID, err := hexToObjectID(request.GetId())
	if err != nil {
		return nil, err
	}
	db := s.Client.Database(database.DBName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var employeeModel models.Employee
	err = models.FindOneEmployee(ctx, db, employeeID).Decode(&employeeModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	var workTimes *WorkTimes
	workTimes = workTimesModelToGRPC(&employeeModel.WorkTimes)
	employeeProto := &EmployeeReply{
		Name:      &employeeModel.Name,
		Surname:   &employeeModel.Surname,
		WorkTimes: workTimes,
	}
	for _, serviceID := range employeeModel.Competence {
		employeeProto.Competence = append(employeeProto.Competence, serviceID.Hex())
	}
	return employeeProto, nil
}

func (s *Server) FindManyEmployees(
	ctx context.Context,
	request *EmployeesRequest,
) (reply *EmployeesReply, err error) {
	companyID, err := hexToObjectID(request.GetCompanyId())
	if err != nil {
		return nil, err
	}
	startValue := primitive.NilObjectID
	if request.StartValue != nil {
		startValue, err = primitive.ObjectIDFromHex(request.GetStartValue())
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	var nPerPage int64 = 30
	if request.NPerPage != nil {
		nPerPage = *request.NPerPage
	}
	db := s.Client.Database(database.DBName)
	cursor, err := models.FindManyEmployees(ctx, db, companyID, startValue, nPerPage)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	defer cursor.Close(ctx)
	reply = &EmployeesReply{}
	for cursor.Next(ctx) {
		var employeeModel models.Employee
		if err := cursor.Decode(&employeeModel); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		employeeID := employeeModel.ID.Hex()
		employeeProto := &EmployeeShort{
			Id:      &employeeID,
			Name:    &employeeModel.Name,
			Surname: &employeeModel.Surname,
		}
		reply.Employees = append(reply.Employees, employeeProto)
	}
	if len(reply.Employees) == 0 {
		return nil, status.Error(
			codes.NotFound,
			"There aren't any companies",
		)
	}
	return reply, nil
}

func (s *Server) FindManyTimeFrames(
	ctx context.Context,
	request *TimeSlotsRequest,
) (*TimeSlotsReply, error) {
	companyID, err := hexToObjectID(request.GetCompanyId())
	if err != nil {
		return nil, err
	}
	serviceID, err := hexToObjectID(request.GetServiceId())
	if err != nil {
		return nil, err
	}
	if request.Day == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"Day field should be set to integer in range 0-6",
		)
	}
	day, err := IntToDay(request.GetDay())
	if err != nil {
		return nil, err
	}
	if request.ServiceDuration == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"Service Duration field should be set to integer in range 0-480",
		)
	}
	serviceDuration := request.GetServiceDuration()
	err = verifyInteger(&serviceDuration, 0, 480)
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"Service Duration field should be set to integer in range 0-480",
		)
	}
	startValue := primitive.NilObjectID
	if request.StartValue != nil {
		startValue, err = primitive.ObjectIDFromHex(request.GetStartValue())
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	var nPerPage int64 = 30
	if request.NPerPage != nil {
		nPerPage = *request.NPerPage
	}
	db := s.Client.Database(database.DBName)
	cursor, err := models.FindManyTimeFrames(
		ctx,
		db,
		companyID,
		serviceID,
		day,
		startValue,
		nPerPage,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var reply TimeSlotsReply 
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	defer cursor.Close(ctx)
	// for each employee
	for cursor.Next(ctx) {
		var employee models.Employee
		if err := cursor.Decode(&employee); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		// slice of work time frames
		timeFrames, err := IntToTimeFrame(&employee, request.GetDay())
		if err != nil {
			return nil, err
		}
		id := employee.ID.Hex()
		employeeInfo := EmployeeShort{
			Id:      &id,
			Name:    &employee.Name,
			Surname: &employee.Surname,
		}
		employeeTimeSlots := EmployeeTimeSlots{
			EmployeeInfo: &employeeInfo,
		}
		// for each work time frame, get all service slots
		for i := range timeFrames {
			serviceSlots := GetServiceSlots(
				&timeFrames[i],
				serviceDuration,
				TimeSlotLength,
				BrakeDuration,
			)
			// flatten service slots
			for j := range serviceSlots {
				employeeTimeSlots.TimeSlots = append(
					employeeTimeSlots.TimeSlots,
					serviceSlots[j],
				)
			}
		}
		reply.EmployeeTimeSlots = append(
			reply.EmployeeTimeSlots,
			&employeeTimeSlots,
		)
	}
	if len(reply.EmployeeTimeSlots) == 0 {
		return nil, status.Error(
			codes.NotFound,
			"There aren't any time slots",
		)
	}
	return &reply, nil
}
