package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username  string
	Password  string
	Follows   []User `gorm:"many2many:follows"`
	Followers []User `gorm:"many2many:followers"`
	Likes     []Like `gorm:"foreignkey:Userid"`
}
