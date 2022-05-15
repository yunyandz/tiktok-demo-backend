package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model

	UserID  uint64 `gorm:"index"`
	VideoID uint64 `gorm:"index"`

	Content string `gorm:"size:1024"`
}

// 获取视频的评论列表
func (v *VideoModel) GetVideoComments(id uint64) ([]*Comment, error) {
	var comments []*Comment
	if err := v.db.Where("video_id = ?", id).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

// 创建一条评论
func (v *VideoModel) CreateComment(videoId uint64, userId uint64, content string) error {
	if err := v.db.Exec("insert into comments (video_id, user_id, content) values (?, ?, ?)", videoId, userId, content).Error; err != nil {
		return err
	}
	return nil
}
