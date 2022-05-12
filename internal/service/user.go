package service

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"

	"github.com/yunyandz/tiktok-demo-backend/internal/model"
)

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

func (u *UserService) GetUser(id uint) (*model.User, error) {
	var user model.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserService) GetFollowList(userId uint) ([]*model.User, error) {
	var users []*model.User
	if err := u.db.Where("id in (select followed_id from follows where follower_id = ?)", userId).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) GetFollowerList(userId uint) ([]*model.User, error) {
	var users []*model.User
	if err := u.db.Where("id in (select follower_id from followers where user_id = ?)", userId).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) CreateFollow(userId uint, followId uint) error {
	if err := u.db.Exec("insert into followers (user_id, follower_id) values (?, ?)", userId, followId).Error; err != nil {
		return err
	}
	return nil
}
