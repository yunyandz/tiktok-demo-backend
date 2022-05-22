package service

import (
	"github.com/yunyandz/tiktok-demo-backend/internal/model"
	"time"
)

type Video struct {
	Id            uint64 `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount uint32 `json:"favorite_count,omitempty"`
	CommentCount  uint32 `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
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
	rsp := Response{}
	vid := &model.VideoModel{}
	if like == true {
		err := vid.LikeVideo(userId, videoId)
		if err != nil {
			rsp.StatusCode = -1
			rsp.StatusMsg = err.Error()
			return rsp
		}
	}
	err := vid.UnLikeVideo(userId, videoId)
	if err != nil {
		rsp.StatusCode = -1
		rsp.StatusMsg = err.Error()
		return rsp
	}
	return Response{StatusCode: 0}
}

func (s *Service) GetVideoList(userId uint64) VideoListResponse {
	return VideoListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: []Video{},
	}
}

func (s *Service) GetLikeList(userId uint64) VideoListResponse {
	// 这儿我蒙了，在哪的API给我调啊
	vid := &model.VideoModel{}
	videos, err := vid.GetUserLikeVideos(userId)
	resvideos := make([]Video, len(videos))
	// TODO server.Video与Model.Video转换
	if err != nil {
		return VideoListResponse{
			Response:  Response{StatusCode: -1, StatusMsg: err.Error()},
			VideoList: []Video{},
		}
	}
	return VideoListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: resvideos,
	}
}
