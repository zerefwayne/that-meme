package config

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

// GlobalConfig ...
type GlobalConfig struct {
	DB    *mongo.Client
	Cache redis.Conn
	S3    *s3.S3
	Env
}

// Env ...
type Env struct {
	AWSConfig
}

// AWSConfig ...
type AWSConfig struct {
	Secret     string
	AccessKey  string
	BucketName string
	RegionCode string
}

var (
	// Config ...
	Config GlobalConfig
)

// LoadEnv ...
func (gc *GlobalConfig) LoadEnv() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	gc.Env.AWSConfig.AccessKey = os.Getenv("AWS_ACCESS_KEY")
	gc.Env.AWSConfig.BucketName = os.Getenv("AWS_S3_BUCKET_NAME")
	gc.Env.AWSConfig.RegionCode = os.Getenv("AWS_S3_REGION_CODE")
	gc.Env.AWSConfig.Secret = os.Getenv("AWS_SECRET")

}
