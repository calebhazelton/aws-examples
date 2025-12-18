package main

import (
	"context"
	"log"
	"os"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/calebhazelton/aws-examples/s3/sdk/s3utils"
)

func main() {
	
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	bucket := os.Getenv("BUCKET_NAME")
	region := os.Getenv("REGION_NAME")

	//Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	s3utils.S3CreateBucket(ctx, client, bucket, region)
	fmt.Println("success")
}