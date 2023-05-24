package main

import (
	"github.com/gin-gonic/gin"

    "github.com/msik-404/micro-appoint-employees/internal/database"
)

func main() {
	mongoClient, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	db := mongoClient.Database("micro-appoint-companies")
	_, err = database.CreateDBIndexes(db)
	if err != nil {
		panic(err)
	}
	// testInsert(db)

	r := gin.Default()

	r.Run() // listen and serve on 0.0.0.0:8080
}
