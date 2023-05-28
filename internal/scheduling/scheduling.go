package scheduling

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

type AppointDate struct {
	WeekDay   Day
	TimeFrame models.TimeFrame
}

func GetAvailableEmployees(coll mongo.Collection, appointDate AppointDate) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	options := options.Find().SetProjection(bson.M{"_id": 1})
    day := fmt.Sprintf("work_times.%s", appointDate.WeekDay)
    filter := bson.M{
        day: bson.M{
            "$elemMatch": bson.M{
                "$and": bson.A{
                    bson.M{"from": bson.M{"$lte": appointDate.TimeFrame.From}},
                    bson.M{"to": bson.M{"$gte": appointDate.TimeFrame.To}},
                },
            },
        },
    }
    return coll.Find(ctx, filter, options)
}
