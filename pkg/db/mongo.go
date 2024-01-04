package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	MongoMaxTimeExec = 10
)

type (
	MongoDB struct {
		Client    *mongo.Client
		DBName    string
		CloseFunc context.CancelFunc
	}

	MongoDBConfig struct {
		Host     string `json:"host"`
		Username string `json:"username"`
		Password string `json:"-"`
		Schema   string `json:"schema"`
	}
)

func setupOptions(config *MongoDBConfig) *options.ClientOptions {
	credential := options.Credential{
		Username: config.Username,
		Password: config.Password,
	}

	clientOptions := options.Client()
	clientOptions.SetAuth(credential)
	clientOptions.ApplyURI(config.Host)

	return clientOptions
}

func NewMongo(config *MongoDBConfig) *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, setupOptions(config))
	if err != nil {
		fmt.Println(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println(err.Error(), err)
	}

	return &MongoDB{
		Client:    client,
		DBName:    config.Schema,
		CloseFunc: cancel,
	}
}
