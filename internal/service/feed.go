package service

import (
	"context"
	"time"

	"github.com/yunyandz/tiktok-demo-backend/internal/errorx"
	"github.com/yunyandz/tiktok-demo-backend/internal/model"
)

func (s *Service) GetFeed(ctx context.Context, isnew bool, lasttime time.Time) (*VideoListResponse, error) {
	vm := model.NewVideoModel(s.db, s.rds)
	var videosModel []*model.Video
	var err error
	if isnew {
		videosModel, err = vm.GetNewVideos()
	} else {
		videosModel, err = vm.GetVideosBeforeTime(lasttime)
	}
	if err != nil {
		s.logger.Sugar().Errorf("get video failed: %s", err.Error())
		return nil, errorx.ErrUserOffline
	}
	userModel := model.NewUserModel(s.db, s.rds)
	// TODO 封装需要的信息
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
