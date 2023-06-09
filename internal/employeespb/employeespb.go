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
	companyID, err := hexToObjectID(request.CompanyId)
	if err != nil {
		return nil, err
	}
	name, err := verifyString(request.Name, 30)
	if err != nil {
		return nil, err
	}
	surname, err := verifyString(request.Surname, 30)
	if err != nil {
		return nil, err
	}
	workTimes, err := verifyWorkTimes(request.WorkTimes)
	if err != nil {
		return nil, err
	}
	competencePlain, err := verifyCompetence(request.Competence)
	if err != nil {
		return nil, err
	}
	var competence []primitive.ObjectID
	for _, hex := range competencePlain {
		serviceID, err := hexToObjectID(hex)
		if err != nil {
			return nil, err
		}
		competence = append(competence, serviceID)
	}
    var workTimesModel *models.WorkTimes
    if workTimes != nil {
        workTimesModel = workTimesGRPCToModel(workTimes)
    }
	newEmployee := models.Employee{
		CompanyID:  companyID,
		Name:       name,
		Surname:    surname,
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
	employeeID, err := hexToObjectID(request.Id)
	if err != nil {
		return nil, err
	}
	// unimplemented
	// companyID, err := grpcHexToObjectID(*request.CompanyId)
	// if err != nil {
	// 	return nil, err
	// }
	name, err := verifyString(request.Name, 30)
	if err != nil {
		return nil, err
	}
	surname, err := verifyString(request.Surname, 30)
	if err != nil {
		return nil, err
	}
	workTimes, err := verifyWorkTimes(request.WorkTimes)
	if err != nil {
		return nil, err
	}
	competencePlain, err := verifyCompetence(request.Competence)
	if err != nil {
		return nil, err
	}
	var competence []primitive.ObjectID
	for _, hex := range competencePlain {
		serviceID, err := hexToObjectID(hex)
		if err != nil {
			return nil, err
		}
		competence = append(competence, serviceID)
	}
	var workTimesModel *models.WorkTimes
	if workTimes != nil {
		workTimesModel = workTimesGRPCToModel(workTimes)
	}
	employeeUpdate := models.Employee{
		Name:       name,
		Surname:    surname,
		WorkTimes:  workTimesModel,
		Competence: competence,
	}
	db := s.Client.Database(database.DBName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := employeeUpdate.UpdateOne(ctx, db, employeeID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if result.MatchedCount == 0 {
		return nil, status.Error(
			codes.NotFound,
			"Employee with that id was not found",
		)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteEmployee(
	ctx context.Context,
	request *DeleteEmployeeRequest,
) (*emptypb.Empty, error) {
	employeeID, err := hexToObjectID(request.Id)
	if err != nil {
		return nil, err
	}
	db := s.Client.Database(database.DBName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := models.DeleteOneEmployee(ctx, db, employeeID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if result.DeletedCount == 0 {
		return nil, status.Error(
			codes.NotFound,
			"Employee with that id was not found",
		)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) FindOneEmployee(
	ctx context.Context,
	request *EmployeeRequest,
) (*EmployeeReply, error) {
	employeeID, err := hexToObjectID(request.Id)
	if err != nil {
		return nil, err
	}
	db := s.Client.Database(database.DBName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	employeeModel := models.Employee{}
	err = models.FindOneEmployee(ctx, db, employeeID).Decode(&employeeModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	var workTimes *WorkTimes
	if employeeModel.WorkTimes != nil {
		workTimes = workTimesModelToGRPC(employeeModel.WorkTimes)
	}
	employeeProto := &EmployeeReply{
		Name:      employeeModel.Name,
		Surname:   employeeModel.Surname,
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
	companyID, err := hexToObjectID(request.CompanyId)
	if err != nil {
		return nil, err
	}
	startValue := primitive.NilObjectID
	if request.StartValue != nil {
		startValue, err = primitive.ObjectIDFromHex(*request.StartValue)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	var nPerPage int64 = 10
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
	reply = &EmployeesReply{}
	for cursor.Next(ctx) {
		var employeeModel models.Employee
		if err := cursor.Decode(&employeeModel); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		employeeProto := &EmployeeShort{
			Id:      employeeModel.ID.Hex(),
			Name:    employeeModel.Name,
			Surname: employeeModel.Surname,
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
