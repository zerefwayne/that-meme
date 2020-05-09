package config

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ConnectS3 ...
func (gc *GlobalConfig) ConnectS3() {

	awsAccessKey := gc.Env.AWSConfig.AccessKey
	awsSecret := gc.Env.AWSConfig.Secret
	token := ""

	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecret, token)

	_, err := creds.Get()

	if err != nil {
		log.Fatal(err)
	}

	cfg := aws.NewConfig().WithRegion("ap-south-1").WithCredentials(creds)

	svc := s3.New(session.New(), cfg)

	gc.S3 = svc

	log.Printf("s3		| connected successfully\n")

}
