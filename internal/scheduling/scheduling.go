package scheduling

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/msik-404/micro-appoint-employees/internal/models"
)

type Day string

const (
	MO Day = "mo"
	TU     = "tu"
	WE     = "we"
	TH     = "th"
	FR     = "fr"
	SA     = "sa"
	SU     = "su"
)

type Appointment struct {
	ServiceID primitive.ObjectID
	WeekDay   Day
	TimeFrame models.TimeFrame
}

// Returns all employees that can perform specific service and are
// available at certain time and week day.
func GetAvailableEmployees(coll mongo.Collection, appointment Appointment) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
    // We want to just return matching ids.
	options := options.Find().SetProjection(bson.M{"_id": 1})
	competenceFilter := bson.M{
		"competence": appointment.ServiceID,
	}
    day := fmt.Sprintf("work_times.%s", appointment.WeekDay)
	dateFilter := bson.M{
		day: bson.M{
			"$elemMatch": bson.M{
				"$and": bson.A{
					bson.M{"from": bson.M{"$lte": appointment.TimeFrame.From}},
					bson.M{"to": bson.M{"$gte": appointment.TimeFrame.To}},
				},
			},
		},
	}
	filter := bson.M{
		"$and": bson.A{competenceFilter, dateFilter},
	}
	return coll.Find(ctx, filter, options)
}
