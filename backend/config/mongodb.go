package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDatabase ...
func (gc *GlobalConfig) ConnectDatabase() {

	ctx := context.Background()

	host := gc.Env.MongoDbEnv.Host
	port := gc.Env.MongoDbEnv.Port
	database := gc.Env.MongoDbEnv.Database

	mongoURI := fmt.Sprintf("mongodb://%s:%s/%s", host, port, database)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongoURI,
	))

	if err != nil {
		log.Fatal(err)
	}

	gc.DB = client

	log.Printf("database	| connected successfully: %s\n", mongoURI)

}
