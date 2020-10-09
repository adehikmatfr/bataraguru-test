package config

import (
	"batara/src/helpers/env"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Connect() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI(env.GetEnv("MONGO_URL"))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	ctx,cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("database Connected...")
	return client.Database(env.GetEnv("MONGO_DATABASE")), nil
}
