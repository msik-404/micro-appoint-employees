package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/msik-404/micro-appoint-employees/internal/database"
	"github.com/msik-404/micro-appoint-employees/internal/controllers/employees"
	"github.com/msik-404/micro-appoint-employees/internal/models"
	"github.com/msik-404/micro-appoint-employees/internal/scheduling"
)

func main() {
	mongoClient, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	db := mongoClient.Database("micro-appoint-employees")
	_, err = database.CreateDBIndexes(db)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// public
	r.GET("/employees", employees.GetEmployeesEndPoint(db))
	// private
	r.GET("/employees/:id", employees.GetEmployeeEndPoint(db))
	r.POST("/employees", employees.AddEmployeeEndPoint(db))
	r.PUT("/employees/:id", employees.UpdateEmployeeEndPoint(db))
	r.DELETE("/employees/:id", employees.DeleteEmployeeEndPoint(db))

	r.Run() // listen and serve on 0.0.0.0:8080
}

func test(db *mongo.Database) {
	value, err := primitive.ObjectIDFromHex("646e64f55e41c9bf2c95fbfa")
	if err != nil {
		panic(err)
	}
	date := scheduling.Appointment{
		ServiceID: value,
		WeekDay:   "mo",
		TimeFrame: models.TimeFrame{
			From: 510,
			To:   600,
		},
	}
	cursor, err := scheduling.GetAvailableEmployees(db.Collection("employees"), date)
	if err != nil {
		panic(err)
	}
	var results []bson.D
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := cursor.All(ctx, &results); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", results)
}
