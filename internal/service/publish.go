package service

import (
	"context"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/yunyandz/tiktok-demo-backend/internal/errorx"
	"github.com/yunyandz/tiktok-demo-backend/internal/model"
)

// 上传video的部分，上传到s3是异步的，所以这里不需要等待立刻返回
func (s *Service) PublishVideo(ctx context.Context, UserID uint64, filename string, videodata io.Reader) Response {
	video := model.Video{
		AuthorID: UserID,
		Playurl:  "",
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
			if err != nil {
				s.logger.Sugar().Errorf("presign video to s3 failed: %s", err.Error())
				return
			}
			s.logger.Sugar().Debugf("upload video to s3 success: %s", filename)
			// 更新video的playurl
			pr, err := s.s3.PresignGetObject(context.TODO(), &s3.GetObjectInput{
				Bucket: aws.String(s.cfg.S3.Bucket),
				Key:    &filename,
			})
			if err != nil {
				s.logger.Sugar().Errorf("presign video to s3 failed: %s", err.Error())
				return
			}
			// 转换为https
			video.Playurl = strings.Replace(pr.URL, "http://", "https://", 1)
			if err := vm.UpdateVideo(vid, video.Playurl); err != nil {
				s.logger.Sugar().Errorf("update video playurl failed: %s", err.Error())
				return
			}
			s.logger.Sugar().Debugf("update video playurl success: %s", video.Playurl)
		}()
	}
	s.logger.Sugar().Debugf("create video success: %d", vid)
	return Response{StatusCode: 0}
}

/**
获取用户所有投稿过的视频
*/
func (s *Service) PublicList(ctx context.Context, UserID uint64) (*VideoListResponse, error) {
	vm := model.NewVideoModel(s.db, s.rds)
	videosModel, err := vm.GetVideosByUser(UserID)
	userModel := model.NewUserModel(s.db, s.rds)
	if err != nil {
		s.logger.Sugar().Errorf("get video failed: %s", err.Error())
		return nil, errorx.ErrUserOffline
	}
	// 封装需要的信息
	var videos []Video
	for _, arr := range videosModel {
		var video Video
		// 根据获取到的authorId去
		user, err := userModel.GetUser(arr.AuthorID)
		if err != nil {
			return nil, errorx.ErrUserDoesNotExists
		}
		// 封装一个Video
		video.Id = uint64(arr.ID)

		video.Author.ID = uint64(user.ID)
		video.Author.Username = user.Username
		video.Author.FollowCount = user.FollowCount
		video.Author.FollowerCount = user.FollowerCount
		// 暂时先设置为false
		// TODO 需要设置一个SQL查询是否关注
		video.Author.IsFollow = false

		video.PlayUrl = arr.Playurl
		video.CoverUrl = arr.Coverurl
		video.FavoriteCount = uint32(arr.Likecount)
		video.CommentCount = uint32(arr.Commentcount)
		video.Title = arr.Title

		videos = append(videos, video)

	}
	rsp := VideoListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		VideoList: videos,
	}
	return &rsp, nil
}
