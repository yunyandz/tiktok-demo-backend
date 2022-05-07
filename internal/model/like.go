package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	Userid  int64
	Videoid int64
}
