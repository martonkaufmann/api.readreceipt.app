package database

import (
	"context"

	"github.com/readreceipt/api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var c *mongo.Client

func Init() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoURL()))
	if err != nil {
		return err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return err
	}

	c = client

	return nil
}

func Client() *mongo.Client {
	return c
}

func Database() *mongo.Database {
	return c.Database(config.MongoDatabase())
}
