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
		var newEmployee models.EmployeeCombRepr
		if err := c.BindJSON(&newEmployee); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		newEmployee.ID = employeeID
		results, err := newEmployee.UpdateCombRepr(db)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, results)
	}
	return gin.HandlerFunc(fn)
}
