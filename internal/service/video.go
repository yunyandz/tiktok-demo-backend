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
	FavoriteCount uint64 `json:"favorite_count,omitempty"`
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
	rsp := Response{}
	vid := model.NewVideoModel(s.db, s.rds)
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
	vido := model.NewVideoModel(s.db, s.rds)
	videos, err := vido.GetUserLikeVideos(userId)
	// TODO is_follow comment_count 字段未实现
	userModel := model.NewUserModel(s.db, s.rds)
	var res = make([]Video, len(videos))
	for i, v := range videos {
		vid := Video{}
		vid.Id = uint64(v.ID)
		vid.PlayUrl = v.Playurl
		vid.CoverUrl = v.Coverurl
		vid.Title = v.Title
		vid.IsFavorite = vido.IsFavorite(userId, uint64(v.ID))
		likeCount, err := vido.GetVideoLikesCount(uint64(v.ID))
		vid.FavoriteCount = uint64(likeCount)
		user, err := userModel.GetUser(v.AuthorID)
		if err != nil {
			res[i] = Video{}
			continue
		}
		vid.Author.ID = uint64(user.ID)
		vid.Author.Username = user.Username
		vid.Author.FollowCount = user.FollowCount
		vid.Author.FollowerCount = user.FollowerCount
		res[i] = vid
	}
	if err != nil {
		return VideoListResponse{
			Response:  Response{StatusCode: -1, StatusMsg: err.Error()},
			VideoList: []Video{},
		}
	}
	return VideoListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: res,
	}
}
