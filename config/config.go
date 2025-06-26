package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
)

var Cfg config.Config

func LoadAWSConfig() {
	awsRegion := os.Getenv("AWS_REGION")
	c, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
	if err != nil {
		log.Fatalf("Unable to load AWS config: %v", err)
	}
	Cfg = c
}
