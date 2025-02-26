package driver

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDriver struct {
	Hostname string
	Username string
	Password string
	PoolSize uint64
}

func NewMongoClient(config MongoDriver) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := config.Hostname

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).
		SetServerAPIOptions(serverAPI).
		SetMaxPoolSize(config.PoolSize).
		SetAuth(options.Credential{
			Username: config.Username,
			Password: config.Password,
		})

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("driver.NewMongoClient: %w", err)
	}

	return client, nil
}
