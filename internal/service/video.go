package service

import "time"

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

func (s *Service) PublishVideo(userId uint64, video Video) Response {
	return Response{StatusCode: 0}
}

func (s *Service) LikeDisliakeVideo(userId uint64, videoId uint64, like bool) Response {
	return Response{StatusCode: 0}
}

func (s *Service) GetVideoList(userId uint64) VideoListResponse {
	return VideoListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: []Video{},
	}
}

func (s *Service) GetLikeList(userId uint64) VideoListResponse {
	return VideoListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: []Video{},
	}
}
