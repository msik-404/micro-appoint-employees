package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/msik-404/micro-appoint-employees/internal/middleware"
	"github.com/msik-404/micro-appoint-employees/internal/models"
)

func GetEmployeeEndPoint(db *mongo.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		employeeId, err := middleware.GetObjectId(c.Param("id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		result := models.FindOneEmployee(db, employeeId)

		type Employee struct {
			Name       string               `json:"name" bson:"name,omitempty"`
			Surname    string               `json:"surname" bson:"surname,omitempty"`
			WorkTimes  models.WorkTimes     `json:"work_times" bson:"work_times,omitempty"`
			Competence []primitive.ObjectID `json:"competence" bson:"competence,omitempty"`
		}
		var employee Employee

		err = result.Decode(&employee)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.AbortWithError(http.StatusNotFound, err)
			} else {
				c.AbortWithError(http.StatusInternalServerError, err)
			}
			return
		}
		c.JSON(http.StatusOK, employee)
	}
	return gin.HandlerFunc(fn)
}
