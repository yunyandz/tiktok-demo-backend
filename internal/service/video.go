package service

import (
	"context"
	"sort"
	"time"

	"github.com/yunyandz/tiktok-demo-backend/internal/model"
)

type Video struct {
	Id            uint64 `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

func (s *Service) LikeDisliakeVideo(userId uint64, videoId uint64, like bool) *Response {
	rsp := Response{}
	vid := model.NewVideoModel(s.db, s.rds)
	if like {
		err := vid.LikeVideo(userId, videoId)
		if err != nil {
			rsp.StatusCode = -1
			rsp.StatusMsg = err.Error()
			return &rsp
		}
		return &Response{StatusCode: 0}
	}
	err := vid.UnLikeVideo(userId, videoId)
	if err != nil {
		rsp.StatusCode = -1
		rsp.StatusMsg = err.Error()
		return &rsp
	}
	return &Response{StatusCode: 0}
}

func (s *Service) GetVideoList(ctx context.Context, selfId uint64, userId uint64) *VideoListResponse {
	vm := model.NewVideoModel(s.db, s.rds)
	s.logger.Sugar().Debugf("get video list,selfId:%d,userId:%d", selfId, userId)
	videos, err := vm.GetVideosByUser(userId)
	if err != nil {
		s.logger.Sugar().Errorf("get video failed: %s", err.Error())
		return &VideoListResponse{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		}
	}
	s.sortVideosByTime(videos)
	// 封装需要的信息
	res := s.convertVideoModeltoVideo(selfId, videos, false)
	return &VideoListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		VideoList: res,
	}
}

func (s *Service) GetLikeList(selfId uint64, userId uint64) *VideoListResponse {
	vido := model.NewVideoModel(s.db, s.rds)
	videos, err := vido.GetUserLikeVideos(userId)
	if err != nil {
		s.logger.Sugar().Errorf("get video failed: %s", err.Error())
		return &VideoListResponse{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "get video failed",
			},
			VideoList: []Video{},
		}
	}
	s.logger.Sugar().Debugf("get video success: %v", videos)
	res := s.convertVideoModeltoVideo(selfId, videos, false)
	if err != nil {
		return &VideoListResponse{
			Response:  Response{StatusCode: -1, StatusMsg: err.Error()},
			VideoList: []Video{},
		}
	}
	s.logger.Sugar().Debugf("get like list success")
	return &VideoListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: res,
	}
}

func (s *Service) convertVideoModeltoVideo(selfid uint64, videos []*model.Video, isTour bool) []Video {
	res, _ := s.convertVideoModeltoVideoWithNextTime(selfid, videos, isTour, time.Now())
	return res
}

func (s *Service) convertVideoModeltoVideoWithNextTime(selfid uint64, videos []*model.Video, isTour bool, lasttime time.Time) ([]Video, time.Time) {
	if lasttime.IsZero() {
		lasttime = time.Now()
	}
	nextTime := lasttime
	var res = make([]Video, len(videos))
	for i, v := range videos {
		vid := Video{}

		user, err := s.getUser(selfid, v.AuthorID)
		if err != nil {
			s.logger.Sugar().Errorf("get user failed: %s", err.Error())
			res[i] = Video{}
			continue
		}
		vid.Author = user

		vid.IsFavorite = false
		if !isTour {
			vid.IsFavorite, err = s.isFavorite(selfid, uint64(v.ID))
			if err != nil {
				s.logger.Sugar().Errorf("get video favorite failed: %s", err.Error())
				res[i] = Video{}
				continue
			}
			s.logger.Sugar().Debugf("get video favorite success: %v", vid.IsFavorite)
		}

		vid.Id = uint64(v.ID)
		vid.PlayUrl = v.Playurl
		vid.CoverUrl = v.Coverurl
		vid.Title = v.Title
		vid.FavoriteCount = v.Likecount
		vid.CommentCount = v.Commentcount

		res[i] = vid
		s.logger.Sugar().Debugf("get video success: %v", vid)

		if v.CreatedAt.Before(nextTime) {
			nextTime = v.CreatedAt
		}
	}
	return res, nextTime
}

func (s *Service) isFavorite(userId uint64, videoId uint64) (bool, error) {
	vid := model.NewVideoModel(s.db, s.rds)
	return vid.IsFavorite(userId, videoId)
}

func (s *Service) sortVideosByTime(videos []*model.Video) {
	sort.Slice(videos, func(i, j int) bool {
		return videos[i].CreatedAt.After(videos[j].CreatedAt)
	})
}
