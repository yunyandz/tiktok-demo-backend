package service

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"

	"github.com/yunyandz/tiktok-demo-backend/internal/model"
)

type VideoService struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewVideoService(db *gorm.DB, rdb *redis.Client) *UserService {
	return &UserService{
		db:  db,
		rdb: rdb,
	}
}

func (u *UserService) CreateVideo(userId uint, video *model.Video) error {
	if err := u.db.Create(video).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserService) UpdateVideoPlayUrl(id uint, playUrl string) error {
	if err := u.db.Exec("update videos set play_url = ? where id = ?", playUrl, id).Error; err != nil {
		return err
	}
	return nil
}

func (v *VideoService) GetNewVideos() ([]*model.Video, error) {
	var videos []*model.Video
	if err := v.db.Limit(30).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func (v *VideoService) GetVideo(id uint) (*model.Video, error) {
	var video model.Video
	if err := v.db.First(&video, id).Error; err != nil {
		return nil, err
	}
	return &video, nil
}

func (v *VideoService) GetVideosByUser(userId uint) ([]*model.Video, error) {
	var videos []*model.Video
	if err := v.db.Where("user_id = ?", userId).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func (v *VideoService) GetUserLikeVideos(userId uint) ([]*model.Video, error) {
	var videos []*model.Video
	if err := v.db.Where("id in (select video_id from likes where user_id = ?)", userId).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func (v *VideoService) GetVideoComments(id uint) ([]*model.Comment, error) {
	var comments []*model.Comment
	if err := v.db.Where("video_id = ?", id).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (v *VideoService) CreateComment(videoId uint, userId uint, content string) error {
	if err := v.db.Exec("insert into comments (video_id, user_id, content) values (?, ?, ?)", videoId, userId, content).Error; err != nil {
		return err
	}
	return nil
}
