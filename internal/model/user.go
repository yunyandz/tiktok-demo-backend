package model

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"size:32;unique_index"`
	Password string `gorm:"size:256"`

	FollowCount   int64 `gorm:"type:int"`
	FollowerCount int64 `gorm:"type:int"`

	Videos    []Video `gorm:"many2many:user_videos"`
	Followers []User  `gorm:"many2many:user_follows"`
	Likes     []Video `gorm:"many2many:user_likes"`
}

type UserModel struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewUserModel(db *gorm.DB, rdb *redis.Client) *UserModel {
	return &UserModel{
		db:  db,
		rdb: rdb,
	}
}

// Todo: 实现redis缓存

// 获取用户信息
func (u *UserModel) GetUser(id uint64) (*User, error) {
	var user User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserModel) GetUserByName(username string) (*User, error) {
	var user User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserModel) CreateUser(user *User) (id uint64, err error) {
	err = u.db.Model(&User{}).Save(user).Error

	return uint64(user.ID), err
}

// 获取用户的关注列表
func (u *UserModel) GetFollowList(userId uint64) ([]*User, error) {
	var users []*User
	if err := u.db.Where("id in (select followed_id from follows where follower_id = ?)", userId).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// 获取用户的粉丝列表
func (u *UserModel) GetFollowerList(userId uint64) ([]*User, error) {
	var users []*User
	if err := u.db.Where("id in (select follower_id from followers where user_id = ?)", userId).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// 关注一个用户
func (u *UserModel) CreateFollow(userId uint64, followId uint64) error {
	if err := u.db.Exec("insert into followers (user_id, follower_id) values (?, ?)", userId, followId).Error; err != nil {
		return err
	}
	return nil
}

// 取消关注一个用户
func (u *UserModel) DeleteFollow(userId uint64, followId uint64) error {
	if err := u.db.Exec("delete from followers where user_id = ? and follower_id = ?", userId, followId).Error; err != nil {
		return err
	}
	return nil
}
