package model

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yunyandz/tiktok-demo-backend/internal/constant"
	"github.com/yunyandz/tiktok-demo-backend/internal/errorx"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model

	AuthorID uint64 `gorm:"column:author_id"`

	Title       string `gorm:"size:128"`
	Description string `gorm:"size:1024"`
	Playurl     string `gorm:"size:1024"`
	Coverurl    string `gorm:"size:1024"`

	Commentcount uint64
	Likecount    uint64 `gorm:"default:0"`

	Likes    []*User    `gorm:"many2many:user_likes"`
	Comments []*Comment `gorm:"many2many:video_comments"`
}

type VideoModel struct {
	db  *gorm.DB
	rdb *redis.Client
}

// Todo: 实现redis缓存

func NewVideoModel(db *gorm.DB, rdb *redis.Client) *VideoModel {
	return &VideoModel{
		db:  db,
		rdb: rdb,
	}
}

// 创建一个新的视频。
func (v *VideoModel) CreateVideo(video *Video) (uint64, error) {
	//事务
	err := v.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(video).Error; err != nil {
			return err
		}
		var user User
		if err := tx.First(&user, video.AuthorID).Error; err != nil {
			return err
		}
		if err := tx.Model(&user).Association("Videos").Append(video); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return uint64(video.ID), nil
}

// 更新视频的播放地址，通常用于视频上传完成后
func (u *VideoModel) UpdateVideo(id uint64, playurl string) error {
	if err := u.db.Model(&Video{}).Where("id = ?", id).Update("playurl", playurl).Error; err != nil {
		return err
	}
	return nil
}

// 获取最新的视频条目，按照时间降序排列，按照文档中的要求，这里只返回前30条
func (v *VideoModel) GetNewVideos() ([]*Video, error) {
	var videos []*Video
	if err := v.db.Limit(constant.FeedLimit).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

// 获取指定时间戳之前创建时间的视频列表。这里依然是最多返回30条。
func (v *VideoModel) GetVideosBeforeTime(time time.Time) ([]*Video, error) {
	var videos []*Video
	if err := v.db.Where("created_at < ?", time).Limit(constant.FeedLimit).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

// 获取视频的详情
func (v *VideoModel) GetVideo(videoId uint64) (*Video, error) {
	var video Video
	if err := v.db.First(&video, videoId).Error; err != nil {
		return nil, err
	}
	return &video, nil
}

// 获取用户的视频列表
func (v *VideoModel) GetVideosByUser(userId uint64) ([]*Video, error) {
	var videos []*Video
	if err := v.db.Where("author_id = ?", userId).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

// 获取用户的视频点赞列表
func (v *VideoModel) GetUserLikeVideos(userId uint64) ([]*Video, error) {
	var videos []*Video
	if err := v.db.Raw("SELECT * FROM videos WHERE id IN (SELECT video_id FROM user_likes WHERE user_id = ?)", userId).Scan(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

// 点赞视频
func (v *VideoModel) LikeVideo(userId uint64, videoId uint64) error {
	err := v.db.Transaction(func(tx *gorm.DB) error {
		if row := v.db.Exec("INSERT INTO user_likes (user_id, video_id) VALUES (?, ?)", userId, videoId).RowsAffected; row != 1 {
			return errorx.ErrUserAlreadyLikeVideo
		}
		if err := v.db.Exec("UPDATE videos SET likecount = likecount + 1 WHERE id = ?", videoId).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// 取消点赞视频
func (v *VideoModel) UnLikeVideo(userId uint64, videoId uint64) error {
	err := v.db.Transaction(func(tx *gorm.DB) error {
		if v.db.Exec("DELETE FROM user_likes WHERE user_id = ? AND video_id = ?", userId, videoId).RowsAffected == 0 {
			return errorx.ErrUserNotLikeVideo
		}
		if err := v.db.Exec("UPDATE videos SET likecount = likecount - 1 WHERE id = ?", videoId).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// 获取视频的点赞数
func (v *VideoModel) GetVideoLikesCount(id uint64) (int64, error) {
	var count int64
	if err := v.db.Model(&Video{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

//查询视频点赞
func (v *VideoModel) IsFavorite(userId uint64, videoId uint64) (bool, error) {
	var count int64
	if err := v.db.Raw("SELECT COUNT(*) FROM user_likes WHERE user_id = ? AND video_id = ?", userId, videoId).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
