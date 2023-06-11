package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBName = os.Getenv("DB_NAME")
const CollName string = "employees"

func getURI() string {
	return fmt.Sprintf("mongodb://%s:%s@employees-db:27017", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
}

func ConnectDB() (*mongo.Client, error) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(getURI()).SetServerAPIOptions(serverAPI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Create a new client and connect to the server
	return mongo.Connect(ctx, opts)
}

func CreateDBIndexes(client *mongo.Client) ([]string, error) {
    db := client.Database(DBName)
	coll := db.Collection("employees")
	index := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "work_times.mo.from", Value: 1},
				{Key: "work_times.mo.to", Value: -1},
			},
		},
		{
			Keys: bson.D{
				{Key: "work_times.tu.from", Value: 1},
				{Key: "work_times.tu.to", Value: -1},
			},
		},
		{
			Keys: bson.D{
				{Key: "work_times.we.from", Value: 1},
				{Key: "work_times.we.to", Value: -1},
			},
		},
		{
			Keys: bson.D{
				{Key: "work_times.th.from", Value: 1},
				{Key: "work_times.th.to", Value: -1},
			},
		},
		{
			Keys: bson.D{
				{Key: "work_times.fr.from", Value: 1},
				{Key: "work_times.fr.to", Value: -1},
			},
		},
		{
			Keys: bson.D{
				{Key: "work_times.sa.from", Value: 1},
				{Key: "work_times.sa.to", Value: -1},
			},
		},
		{
			Keys: bson.D{
				{Key: "work_times.su.from", Value: 1},
				{Key: "work_times.su.to", Value: -1},
			},
		},
		{
			Keys: bson.M{"competence": 1},
		},
		{
			Keys: bson.M{"company_id": 1},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return coll.Indexes().CreateMany(ctx, index)
}
