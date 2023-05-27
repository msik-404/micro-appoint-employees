package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/msik-404/micro-appoint-employees/internal/middleware"
	"github.com/msik-404/micro-appoint-employees/internal/models"
)

func GetEmployeesEndPoint(db *mongo.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		coll := db.Collection("employees")
		employees, err := middleware.PaginHandler[models.Employee](c, coll, nil)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, employees)
	}
	return gin.HandlerFunc(fn)
}
