package service

import (
	"context"
	"encoding/json"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-redis/redis/v8"

	"github.com/yunyandz/tiktok-demo-backend/internal/constant"
	"github.com/yunyandz/tiktok-demo-backend/internal/model"
)

// 上传video的部分，上传到s3是异步的，所以这里不需要等待立刻返回
func (s *Service) PublishVideo(ctx context.Context, UserID uint64, filename string, videodata io.Reader, title string) Response {
	playurl := ""
	if s.cfg.S3.Vaild {
		pr, err := s.s3.PresignGetObject(context.TODO(), &s3.GetObjectInput{
			Bucket: aws.String(s.cfg.S3.Bucket),
			Key:    &filename,
		})
		if err != nil {
			s.logger.Sugar().Errorf("presign video to s3 failed: %s", err.Error())
			return Response{StatusCode: 1, StatusMsg: err.Error()}
		}
		playurl = strings.Replace(pr.URL, "http://", "https://", 1)
	}
	// 转换为https
	video := model.Video{
		AuthorID: UserID,
		Title:    title,
		Playurl:  playurl,
	}
	vm := model.NewVideoModel(s.db, s.rds)
	vid, err := vm.CreateVideo(&video)
	if err != nil {
		s.logger.Sugar().Errorf("create video failed: %s", err.Error())
		return Response{StatusCode: 1, StatusMsg: err.Error()}
	}

	if s.cfg.S3.Vaild {
		go func() {
			_, err := s.s3.PutObject(context.TODO(), &s3.PutObjectInput{
				Bucket:      aws.String(s.cfg.S3.Bucket),
				Key:         &filename,
				Body:        videodata,
				ContentType: aws.String("video/mp4"),
			})
			if err != nil {
				s.logger.Sugar().Errorf("upload video to s3 failed: %s", err.Error())
				return
			}
			s.logger.Sugar().Debugf("upload video to s3 success: %s", filename)
		}()
	}
	if s.cfg.Redis.Vaild {
		go func() {
			su, err := s.GetUserInfo(video.AuthorID)
			if err != nil {
				s.logger.Sugar().Errorf("get user info failed: %s", err.Error())
				return
			}
			sv := Video{
				Id:      uint64(video.ID),
				Author:  su.User,
				Title:   video.Title,
				PlayUrl: video.Playurl,
			}
			mb, err := json.Marshal(sv)
			if err != nil {
				s.logger.Sugar().Errorf("marshal video failed: %s", err.Error())
				return
			}
			s.rds.ZAdd(ctx, constant.RedisVideoZSetKey, &redis.Z{
				Score:  float64(video.CreatedAt.Unix()),
				Member: mb,
			})
			s.logger.Sugar().Debugf("create video cache success: %d", vid)
		}()
	}
	s.logger.Sugar().Debugf("create video success: %d", vid)
	return Response{StatusCode: 0}
}
