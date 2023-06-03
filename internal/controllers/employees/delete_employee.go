package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/msik-404/micro-appoint-employees/internal/middleware"
	"github.com/msik-404/micro-appoint-employees/internal/models"
)

// cascade delete(this will also delete employeeInfo)
func DeleteEmployeeEndPoint(db *mongo.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		employeeID, err := middleware.GetObjectId(c.Param("id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
        result, err := models.DeleteOneEmployee(db, employeeID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, result)
	}
	return gin.HandlerFunc(fn)
}
