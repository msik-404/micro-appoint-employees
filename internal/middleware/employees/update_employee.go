package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msik-404/micro-appoint-employees/internal/middleware"
	"github.com/msik-404/micro-appoint-employees/internal/models"
	"github.com/msik-404/micro-appoint-employees/internal/strtime"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateEmployeeEndPoint(db *mongo.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		employeeID, err := middleware.GetObjectId(c.Param("id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		type EmployeeUpdateStr struct {
            Name       string                `json:"name" binding:"max=30"`
            Surname    string                `json:"surname" binding:"max=30"`
			WorkTimes  *strtime.WorkTimesStr `json:"work_times"`
			Competence []primitive.ObjectID  `json:"competence"`
		}
		var employeeUpdateStr EmployeeUpdateStr
		if err := c.BindJSON(&employeeUpdateStr); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		var workTimesUpdate *models.WorkTimes = nil
		if employeeUpdateStr.WorkTimes != nil {
            workTimes, err := strtime.ToWorkTimes(employeeUpdateStr.WorkTimes)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
            workTimesUpdate = &workTimes
		}
		employeeUpdate := models.EmployeeUpdate{
			Name:       employeeUpdateStr.Name,
			Surname:    employeeUpdateStr.Surname,
			WorkTimes:  workTimesUpdate,
			Competence: employeeUpdateStr.Competence,
		}
		results, err := employeeUpdate.UpdateOne(db, employeeID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, results)
	}
	return gin.HandlerFunc(fn)
}
