package employees

import (
	"context"
    "time"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/msik-404/micro-appoint-employees/internal/middleware"
	"github.com/msik-404/micro-appoint-employees/internal/models"
)

func GetEmployeesEndPoint(db *mongo.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		startValue, err := middleware.GetStartValue(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		nPerPage, err := middleware.GetNPerPageValue(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		cursor, err := models.FindManyEmployees(db, startValue, nPerPage)

		type Employee struct {
			ID      primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
			Name    string             `json:"name" bson:"name,omitempty"`
			Surname string             `json:"surname" bson:"surname,omitempty"`
		}
		var employees []Employee

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := cursor.All(ctx, &employees); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if len(employees) == 0 {
			err = errors.New("No documents in the result")
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusOK, employees)
	}
	return gin.HandlerFunc(fn)
}
