package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Userid  int64
	Videoid int64
	Content string
}
