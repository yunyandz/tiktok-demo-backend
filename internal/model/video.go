package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model

	Userid       int64
	Title        string
	Description  string
	Playurl      string
	Coverurl     string
	Commentcount int64
	Likecount    int64

	Likes    []Like    `gorm:"foreignkey:Videoid"`
	Comments []Comment `gorm:"foreignkey:Videoid"`
}
