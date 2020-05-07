package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connections ...
type Connections struct {
	DB *mongo.Client
}

// ConnectDatabase ...
func (c *Connections) ConnectDatabase() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	host := "localhost"
	port := "27017"
	database := "thatmeme"

	mongoURI := fmt.Sprintf("mongodb://%s:%s/%s", host, port, database)
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongoURI,
	))

	if err != nil {
		log.Fatal(err)
	}

	c.DB = client

	log.Printf("database	| connected successfully: %s\n", mongoURI)

}