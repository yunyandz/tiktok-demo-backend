package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string
	Password string
	Follows  []User `gorm:"many2many:follows"`
	Followed []User `gorm:"many2many:followed"`
}
