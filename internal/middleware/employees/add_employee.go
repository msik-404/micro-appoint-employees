package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/msik-404/micro-appoint-employees/internal/models"
	"github.com/msik-404/micro-appoint-employees/internal/strtime"
)

func AddEmployeeEndPoint(db *mongo.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		type EmployeeStr struct {
			Name       string               `json:"name"`
			Surname    string               `json:"surname"`
			WorkTimes  strtime.WorkTimesStr `json:"work_times"`
			Competence []primitive.ObjectID `json:"competence"`
		}
		var newEmployeeStr EmployeeStr
		if err := c.BindJSON(&newEmployeeStr); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
        workTimes, err := strtime.ToWorkTimes(&newEmployeeStr.WorkTimes)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
        newEmployee := models.Employee{
            Name: newEmployeeStr.Name,
            Surname: newEmployeeStr.Surname,
            WorkTimes: workTimes,
            Competence: newEmployeeStr.Competence,
        }
		result, err := newEmployee.InsertOne(db)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, result)
	}
	return gin.HandlerFunc(fn)
}
