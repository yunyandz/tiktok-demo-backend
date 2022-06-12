package service

import (
	"context"
	"time"

	"github.com/yunyandz/tiktok-demo-backend/internal/errorx"
	"github.com/yunyandz/tiktok-demo-backend/internal/model"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

func (s *Service) GetFeed(ctx context.Context, userId uint64, isnew bool, isTour bool, lasttime time.Time) (*FeedResponse, error) {
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
	s.sortVideosByTime(videosModel)
	videos, nt := s.convertVideoModeltoVideoWithNextTime(userId, videosModel, isTour, lasttime)
	rsp := FeedResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		VideoList: videos,
		NextTime:  nt.UnixMilli(),
	}
	return &rsp, nil
}
