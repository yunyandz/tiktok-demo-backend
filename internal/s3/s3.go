package s3

import (
	"context"

	"github.com/yunyandz/tiktok-demo-backend/internal/config"
	"github.com/yunyandz/tiktok-demo-backend/internal/constant"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/zap"
)

func New(config *config.Config, logger *zap.Logger) *s3.Client {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           config.S3.Endpoint,
			SigningRegion: config.S3.Region,
		}, nil
	})
	cfg, err := awscfg.LoadDefaultConfig(context.TODO(), awscfg.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		panic(err)
	}
	client := s3.NewFromConfig(cfg)

	// 检查bucket是否存在，不存在则创建
	burkets, err := client.ListBuckets(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	hastiktokbucket := false
	for _, bucket := range burkets.Buckets {
		logger.Info("bucket", zap.String("name", *bucket.Name))
		if *bucket.Name == constant.BucketName {
			hastiktokbucket = true
		}
	}
	if !hastiktokbucket {
		_, err := client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
			Bucket: aws.String(constant.BucketName),
		})
		if err != nil {
			panic(err)
		}
	}
	return client
}

type S3ObjectAPI interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}
