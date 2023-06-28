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
	return fmt.Sprintf(
        "mongodb://%s:%s@%s:27017", 
        os.Getenv("DB_USER"), 
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOSTNAME"),
    )
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
                {Key: "id", Value: 1},
                {Key: "company_id", Value: 1},
            },
        },
        {
            Keys: bson.D{
                {Key: "company_id", Value: 1},
                {Key: "competence", Value: 1},
            },
        },
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return coll.Indexes().CreateMany(ctx, index)
}
