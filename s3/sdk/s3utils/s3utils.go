package s3utils

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// Creates an S3 bucket with the given name.
func S3CreateBucket(ctx context.Context, client *s3.Client, bucket string, region string) error {
	fmt.Println("Create Bucket")
	if bucket == "" {
		return fmt.Errorf("bucket name is required")
	}

	input := &s3.CreateBucketInput{
		Bucket: &bucket,
	}

	// Determine the region from the client	
	// If the client's region is anything other than us-east-1, LocationConstraint is required.
	if region != "us-east-1" && region != "" {
		input.CreateBucketConfiguration = &types.CreateBucketConfiguration {
			LocationConstraint: types.BucketLocationConstraint(region),
		}
	}

    _, err := client.CreateBucket(ctx, input)
    if err != nil {
        return fmt.Errorf("create bucket %q: %w", bucket, err)
    }

	return nil
}