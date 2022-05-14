package service

import "time"

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
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

func (s *Service) Feed(userId int64, lasttime time.Time) FeedResponse {
	return FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: []Video{},
		NextTime:  time.Now().Unix(),
	}
}

func (s *Service) PublishVideo(userId int64, video Video) Response {
	return Response{StatusCode: 0}
}

func (s *Service) LikeDisliakeVideo(userId int64, videoId int64, like bool) Response {
	return Response{StatusCode: 0}
}

func (s *Service) GetVideoList(userId int64) VideoListResponse {
	return VideoListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: []Video{},
	}
}

func (s *Service) GetLikeList(userId int64) VideoListResponse {
	return VideoListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: []Video{},
	}
}
