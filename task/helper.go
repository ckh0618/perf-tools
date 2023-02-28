package task

import (
	"context"
	"github.com/Netflix/go-env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	MongoDBUri string `env:"MONGODB_URI"`
}

func NewConfig() (*Config, error) {

	c := new(Config)
	_, err := env.UnmarshalFromEnviron(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func GetConnection(ctx context.Context) (*mongo.Client, error) {

	cfg, err := NewConfig()
	if err != nil {
		return nil, err
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDBUri))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}
