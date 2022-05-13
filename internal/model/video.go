package model

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model

	Userid       int64
	Title        string
	Description  string
	Playurl      string
	Coverurl     string
	Commentcount int64
	Likecount    int64

	Likes    []User    `gorm:"many2many:user_likes"`
	Comments []Comment `gorm:"foreignkey:Videoid"`
}

type VideoService struct {
	db  *gorm.DB
	rdb *redis.Client
}

// Todo: 实现redis缓存

func NewVideoService(db *gorm.DB, rdb *redis.Client) *UserService {
	return &UserService{
		db:  db,
		rdb: rdb,
	}
}

// 创建一个新的视频，注意：这里不需要检查用户是否存在
func (v *VideoService) CreateVideo(video *Video) error {
	if err := v.db.Create(video).Error; err != nil {
		return err
	}
	return nil
}

// 更新视频的播放地址，通常用于视频上传完成后
func (u *UserService) UpdateVideoPlayUrl(id uint, playUrl string) error {
	if err := u.db.Exec("update videos set play_url = ? where id = ?", playUrl, id).Error; err != nil {
		return err
	}
	return nil
}

// 获取最新的视频条目，按照时间降序排列，按照文档中的要求，这里只返回前30条
func (v *VideoService) GetNewVideos() ([]*Video, error) {
	var videos []*Video
	if err := v.db.Limit(30).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

// 获取视频的详情
func (v *VideoService) GetVideo(id uint) (*Video, error) {
	var video Video
	if err := v.db.First(&video, id).Error; err != nil {
		return nil, err
	}
	return &video, nil
}

// 获取用户的视频列表
func (v *VideoService) GetVideosByUser(userId uint) ([]*Video, error) {
	var videos []*Video
	if err := v.db.Where("user_id = ?", userId).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

// 获取用户的视频点赞列表
func (v *VideoService) GetUserLikeVideos(userId uint) ([]*Video, error) {
	var videos []*Video
	if err := v.db.Where("id in (select video_id from likes where user_id = ?)", userId).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

// 获取视频的评论列表
func (v *VideoService) GetVideoComments(id uint) ([]*Comment, error) {
	var comments []*Comment
	if err := v.db.Where("video_id = ?", id).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

// 创建一条评论
func (v *VideoService) CreateComment(videoId uint, userId uint, content string) error {
	if err := v.db.Exec("insert into comments (video_id, user_id, content) values (?, ?, ?)", videoId, userId, content).Error; err != nil {
		return err
	}
	return nil
}

// 点赞视频
func (v *VideoService) LikeVideo(userId uint, videoId uint) error {
	if err := v.db.Exec("insert into likes (user_id, video_id) values (?, ?)", userId, videoId).Error; err != nil {
		return err
	}
	return nil
}

// 取消点赞视频
func (v *VideoService) UnLikeVideo(userId uint, videoId uint) error {
	if err := v.db.Exec("delete from likes where user_id = ? and video_id = ?", userId, videoId).Error; err != nil {
		return err
	}
	return nil
}

// 获取视频的点赞数
func (v *VideoService) GetVideoLikesCount(id uint) (int, error) {
	var count int64
	if err := v.db.Model(&Video{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
