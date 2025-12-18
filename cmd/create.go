/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/calebhazelton/aws-examples/s3/sdk"
)

var (
	bucketName string
	region string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an S3 bucket",
	Long: `Creat an S3 bucket using AWS SDK for Go v2.
	
	Parameters required --bucket, --region`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if bucketName == "" {
			return fmt.Errorf("bucket name is required (use --bucket)")
		}
		if region == "" {
			return fmt.Errorf("region is required (use --region)")
		}

		ctx := context.TODO()

		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			log.Fatal(err)
		}

		//Create an Amazon S3 service client
		client := s3.NewFromConfig(cfg)

		if err := s3utils.S3CreateBucket(ctx, client, bucketName, region); err != nil {
            return fmt.Errorf("create bucket failed: %w", err)
        }

        fmt.Printf("Created bucket %q in %q\n", bucketName, region)
        return nil
	},


}

func init() {
    rootCmd.AddCommand(createCmd)

    createCmd.Flags().StringVarP(&bucketName, "bucket", "b", "", "Bucket name (required)")
    createCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region (required)")
    _ = createCmd.MarkFlagRequired("bucket")
    _ = createCmd.MarkFlagRequired("region")
}