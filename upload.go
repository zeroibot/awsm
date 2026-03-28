package awsm

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager"
	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/zeroibot/pack/lang"
	"github.com/zeroibot/pack/str"
)

type UploadConfig struct {
	Profile     string
	Region      string
	Bucket      string
	FilePath    string
	BucketPath  string
	ACL         types.ObjectCannedACL
	ContentType string
}

// UploadFile uploads a file to an S3 bucket
func UploadFile(cfg *UploadConfig) error {
	// Load AWS configuration
	profile := lang.Ternary(str.IsEmpty(cfg.Profile), "default", cfg.Profile)
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(profile),
		config.WithRegion(cfg.Region),
	)
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}

	// Create S3 client
	client := s3.NewFromConfig(awsCfg)

	// Open local file
	file, err := os.Open(cfg.FilePath)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %w", cfg.FilePath, err)
	}
	defer file.Close()

	// Create S3 uploader
	uploader := transfermanager.New(client)

	// Upload file
	_, err = uploader.UploadObject(context.TODO(), &transfermanager.UploadObjectInput{
		Bucket:      new(cfg.Bucket),
		Key:         new(cfg.BucketPath),
		Body:        file,
		ACL:         cfg.ACL,
		ContentType: new(cfg.ContentType),
	})
	if err != nil {
		return fmt.Errorf("failed to upload file %q to S3: %w", cfg.FilePath, err)
	}

	return nil
}

// PublicURL returns the public URL of the uploaded file
func (cfg UploadConfig) PublicURL() string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", cfg.Bucket, cfg.Region, cfg.BucketPath)
}
