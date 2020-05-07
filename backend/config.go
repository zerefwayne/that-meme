package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connections ...
type Connections struct {
	DB *mongo.Client
	Cache redis.Conn
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

// ConnectCache ...
func (c *Connections) ConnectCache() {

	redisURI := "redis://localhost"

	conn, err := redis.DialURL(redisURI)

	if err != nil {
		log.Fatal(err)
	}

	c.Cache = conn

	log.Printf("cache	| connected successfully: %s\n", redisURI)
}