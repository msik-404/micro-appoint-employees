package main

import (
	"github.com/gin-gonic/gin"

	"github.com/msik-404/micro-appoint-employees/internal/database"
	"github.com/msik-404/micro-appoint-employees/internal/middleware/employees"
)

func main() {
	mongoClient, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	db := mongoClient.Database("micro-appoint-employees")
	// _, err = database.CreateDBIndexes(db)
	// if err != nil {
	// 	panic(err)
	// }
	// testInsert(db)

	r := gin.Default()

	r.POST("/employees", employees.AddEmployeeEndPoint(db))

	r.Run() // listen and serve on 0.0.0.0:8080
}
