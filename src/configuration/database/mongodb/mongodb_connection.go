package mongodb

import (
	"api-user-golang/src/configuration/logger"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL    = "MONGO_URL"
	MONGO_DATABASE = "MONGO_DB_NAME"
)

func NewMongoDbConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGO_DATABASE)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("MongoDB connection established successfully")

	return client.Database(mongodb_database), nil
}
