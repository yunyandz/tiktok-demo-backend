package service

import (
	"context"
	"time"

	"github.com/yunyandz/tiktok-demo-backend/internal/errorx"
	"github.com/yunyandz/tiktok-demo-backend/internal/model"
)

type Video struct {
	Id            uint64 `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount uint32 `json:"favorite_count,omitempty"`
	CommentCount  uint32 `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

func (s *Service) Feed(userId uint64, lasttime time.Time) FeedResponse {
	return FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: []Video{},
		NextTime:  time.Now().Unix(),
	}
}

func (s *Service) LikeDisliakeVideo(userId uint64, videoId uint64, like bool) Response {
	return Response{StatusCode: 0}
}

func (s *Service) GetVideoList(ctx context.Context, UserID uint64) (*VideoListResponse, error) {
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

func (s *Service) GetLikeList(userId uint64) VideoListResponse {
	return VideoListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: []Video{},
	}
}
