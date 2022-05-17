package service

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/yunyandz/tiktok-demo-backend/internal/constant"
	"github.com/yunyandz/tiktok-demo-backend/internal/model"
)

func (s *Service) PublishVideo(UserID uint64, filename string, videodata io.Reader) Response {
	go func() {
		_, err := s.s3.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(constant.BucketName),
			Key:    &filename,
			Body:   videodata,
		})
		if err != nil {
			s.logger.Sugar().Errorf("upload video to s3 failed: %s", err.Error())
			return
		}
	}()
	video := model.Video{
		AuthorID: UserID,
		Playurl:  "https://" + constant.BucketName + ".s3.amazonaws.com/" + filename,
	}
	vm := model.NewVideoModel(s.db, s.rds)
	err := vm.CreateVideo(&video)
	if err != nil {
		s.logger.Sugar().Errorf("create video failed: %s", err.Error())
		return Response{StatusCode: 1, StatusMsg: err.Error()}
	}
	return Response{StatusCode: 0}
}
