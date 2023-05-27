package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/msik-404/micro-appoint-employees/internal/models"
)

func AddEmployeeEndPoint(db *mongo.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var newEmployee models.EmployeeCombRepr
		if err := c.BindJSON(&newEmployee); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		result, err := newEmployee.InsertCombRepr(db)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, result)
	}
	return gin.HandlerFunc(fn)
}
