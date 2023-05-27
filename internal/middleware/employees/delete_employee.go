package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
		coll := db.Collection("employees")
		filter := bson.D{{"_id", employeeID}}
		var results []*mongo.DeleteResult
		result, err := models.GenericDeleteOne(coll, filter)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		results = append(results, result)
		coll = db.Collection("employee_infos")
		result, err = models.GenericDeleteOne(coll, filter)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		results = append(results, result)
		c.JSON(http.StatusOK, result)
	}
	return gin.HandlerFunc(fn)
}
