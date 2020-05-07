package config

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

// Connections ...
type Connections struct {
	DB    *mongo.Client
	Cache redis.Conn
	S3    *s3.S3
}

var (
	// Config ...
	Config Connections
)
