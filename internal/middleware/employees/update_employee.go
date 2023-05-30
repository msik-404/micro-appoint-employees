package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msik-404/micro-appoint-employees/internal/middleware"
	"github.com/msik-404/micro-appoint-employees/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateEmployeeEndPoint(db *mongo.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		employeeID, err := middleware.GetObjectId(c.Param("id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		var employeeUpdate models.EmployeeUpdate
		if err := c.BindJSON(&employeeUpdate); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
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
