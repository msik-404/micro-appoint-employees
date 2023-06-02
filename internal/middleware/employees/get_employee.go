package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/msik-404/micro-appoint-employees/internal/middleware"
	"github.com/msik-404/micro-appoint-employees/internal/models"
	"github.com/msik-404/micro-appoint-employees/internal/strtime"
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
			Name       string               `bson:"name,omitempty"`
			Surname    string               `bson:"surname,omitempty"`
			WorkTimes  models.WorkTimes     `bson:"work_times,omitempty"`
			Competence []primitive.ObjectID `bson:"competence,omitempty"`
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

		type EmployeeStr struct {
			Name       string               `json:"name"`
			Surname    string               `json:"surname"`
			WorkTimes  strtime.WorkTimesStr `json:"work_times"`
			Competence []primitive.ObjectID `json:"competence"`
		}
        workTimeStr, err := strtime.ToWorkTimesStr(&employee.WorkTimes)
        if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
            return
        }
        employeeStr := EmployeeStr {
            employee.Name,
            employee.Surname,
            workTimeStr,
            employee.Competence,
        }
		c.JSON(http.StatusOK, employeeStr)
	}
	return gin.HandlerFunc(fn)
}
