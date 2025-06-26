package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"              // This is the correct type
	awsconfig "github.com/aws/aws-sdk-go-v2/config" // Alias to avoid confusion
)

var Cfg aws.Config // ✅ Correct type

func LoadAWSConfig() {
	awsRegion := os.Getenv("AWS_REGION")
	c, err := awsconfig.LoadDefaultConfig(context.TODO(), awsconfig.WithRegion(awsRegion))
	if err != nil {
		log.Fatalf("Unable to load AWS config: %v", err)
	}
	Cfg = c
}
