package model

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string
	Password string

	Videos    []Video `gorm:"many2many:user_videos"`
	Followers []User  `gorm:"many2many:user_follows"`
	Likes     []Video `gorm:"many2many:user_likes"`
}

type UserService struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewUserService(db *gorm.DB, rdb *redis.Client) *UserService {
	return &UserService{
		db:  db,
		rdb: rdb,
	}
}

// Todo: 实现redis缓存

// 获取用户信息
func (u *UserService) GetUser(id uint) (*User, error) {
	var user User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 获取用户的关注列表
func (u *UserService) GetFollowList(userId uint) ([]*User, error) {
	var users []*User
	if err := u.db.Where("id in (select followed_id from follows where follower_id = ?)", userId).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// 获取用户的粉丝列表
func (u *UserService) GetFollowerList(userId uint) ([]*User, error) {
	var users []*User
	if err := u.db.Where("id in (select follower_id from followers where user_id = ?)", userId).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// 关注一个用户
func (u *UserService) CreateFollow(userId uint, followId uint) error {
	if err := u.db.Exec("insert into followers (user_id, follower_id) values (?, ?)", userId, followId).Error; err != nil {
		return err
	}
	return nil
}

// 取消关注一个用户
func (u *UserService) DeleteFollow(userId uint, followId uint) error {
	if err := u.db.Exec("delete from followers where user_id = ? and follower_id = ?", userId, followId).Error; err != nil {
		return err
	}
	return nil
}
