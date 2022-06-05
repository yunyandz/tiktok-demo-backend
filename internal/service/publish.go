package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-redis/redis/v8"

	"github.com/yunyandz/tiktok-demo-backend/internal/constant"
	"github.com/yunyandz/tiktok-demo-backend/internal/model"
	"github.com/yunyandz/tiktok-demo-backend/internal/util"
)

// 上传video的部分，上传到s3是异步的，所以这里不需要等待立刻返回
func (s *Service) PublishVideo(ctx context.Context, UserID uint64, filename string, videodata io.Reader, title string) Response {
	playurl := ""
	coverurl := ""
	filename = strings.Join([]string{s.Hash([]byte(filename + title)), filename}, "-")
	coverfilename := s.GetCoverFileName(filename)
	if s.cfg.S3.Vaild {
		var err error
		pus, err := s.PreSignUrl(&filename)
		if err != nil {
			return Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			}
		}
		cus, err := s.PreSignUrl(&coverfilename)
		if err != nil {
			return Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			}
		}
		purl, _ := url.Parse(pus)
		curl, _ := url.Parse(cus)
		playurl = util.GetRawUrl(purl)
		coverurl = util.GetRawUrl(curl)
	}
	video := model.Video{
		AuthorID: UserID,
		Title:    title,
		Playurl:  playurl,
		Coverurl: coverurl,
	}
	vm := model.NewVideoModel(s.db, s.rds)
	vid, err := vm.CreateVideo(&video)
	if err != nil {
		s.logger.Sugar().Errorf("create video failed: %s", err.Error())
		return Response{StatusCode: 1, StatusMsg: err.Error()}
	}

	if s.cfg.S3.Vaild {
		go s.UploadVideoToS3(ctx, filename, videodata)
		go func() {
			coverdata, err := s.GetCoverFromVideoFile(videodata)
			if err != nil {
				s.logger.Sugar().Errorf("get cover from video failed: %s", err.Error())
				return
			}
			s.logger.Sugar().Debugf("get cover from video success")
			s.UploadCoverToS3(ctx, coverfilename, coverdata)
		}()
	}
	if s.cfg.Redis.Vaild {
		go s.PutVideoInfoToRedis(ctx, &video)
	}
	s.logger.Sugar().Debugf("create video success: %d", vid)
	return Response{StatusCode: 0}
}

func (s *Service) UploadVideoToS3(ctx context.Context, filename string, videodata io.Reader) {
	s.UploadToS3(ctx, filename, videodata, "video/mp4")
}

func (s *Service) UploadCoverToS3(ctx context.Context, filename string, coverdata io.Reader) {
	s.UploadToS3(ctx, filename, coverdata, "image/jpeg")
}

func (s *Service) UploadToS3(ctx context.Context, filename string, data io.Reader, contenttype string) {
	_, err := s.s3.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.cfg.S3.Bucket),
		Key:         &filename,
		Body:        data,
		ContentType: &contenttype,
	})
	if err != nil {
		s.logger.Sugar().Errorf("upload to s3 failed: %s", err.Error())
		return
	}
	s.logger.Sugar().Debugf("upload to s3 success: %s", filename)
}

func (s *Service) PreSignUrl(filename *string) (string, error) {
	pr, err := s.s3.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s.cfg.S3.Bucket),
		Key:    filename,
	})
	if err != nil {
		s.logger.Sugar().Errorf("presign url to s3 failed: %s", err.Error())
		return "", err
	}
	//转换为https
	playurl := strings.Replace(pr.URL, "http://", "https://", 1)
	playurl = strings.TrimRight(playurl, "/")
	s.logger.Sugar().Debugf("presign url to s3 success: %s", playurl)
	return playurl, nil
}

func (s *Service) PutVideoInfoToRedis(ctx context.Context, video *model.Video) {
	su, err := s.GetUserInfo(video.AuthorID)
	if err != nil {
		s.logger.Sugar().Errorf("get user info failed: %s", err.Error())
		return
	}
	sv := Video{
		Id:       uint64(video.ID),
		Author:   su.User,
		Title:    video.Title,
		PlayUrl:  video.Playurl,
		CoverUrl: video.Coverurl,
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
	s.logger.Sugar().Debugf("create video cache success: %d", video.ID)
}

func (s *Service) GetCoverFromVideoFile(data io.Reader) (io.Reader, error) {
	tmpf, err := os.CreateTemp("", "tiktok-*.mp4")
	if err != nil {
		s.logger.Sugar().Errorf("create temp file failed: %s", err.Error())
		return nil, err
	}
	defer tmpf.Close()
	_, err = io.Copy(tmpf, data)
	if err != nil {
		s.logger.Sugar().Errorf("copy video data failed: %s", err.Error())
		return nil, err
	}
	coverfile, err := util.ReadFrameAsJpeg(tmpf.Name())
	if err != nil {
		s.logger.Sugar().Errorf("read frame as jpeg failed: %s", err.Error())
		return nil, err
	}
	s.logger.Sugar().Debugf("read frame as jpeg success")
	return coverfile, nil
}

func (s *Service) GetCoverFileName(filename string) string {
	return strings.TrimSuffix(filename, ".mp4") + ".jpg"
}

func (s *Service) Hash(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
