package appmongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const name = "mongo"

type Config struct {
	URI     string
	Timeout int
}

func NewDriver(ctx context.Context, config Config) (*mongo.Database, error) {
	if len(config.URI) == 0 {
		return nil, errors.New("invalid_mongo_uri")
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(config.Timeout)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.URI))
	if err != nil {
		return nil, err
	}

	return client.Database("backend_challenge"), nil

	//return client, nil
}
