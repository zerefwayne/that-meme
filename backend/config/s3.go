package config

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ConnectS3 ...
func (c *Connections) ConnectS3() {

	awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
	awsSecret := os.Getenv("AWS_SECRET")
	token := ""

	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecret, token)

	_, err := creds.Get()

	if err != nil {
		log.Fatal(err)
	}

	cfg := aws.NewConfig().WithRegion("ap-south-1").WithCredentials(creds)

	svc := s3.New(session.New(), cfg)

	c.S3 = svc

	log.Printf("s3		| connected successfully\n")

}
