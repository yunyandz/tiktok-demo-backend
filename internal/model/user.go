package model

import (
	"errors"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"size:32;unique_index"`
	Password string `gorm:"size:256"`

	FollowCount   int64 `gorm:"type:int"`
	FollowerCount int64 `gorm:"type:int"`

	Videos    []*Video `gorm:"many2many:user_videos"`
	Followers []*User  `gorm:"many2many:user_follows"`
	Likes     []*Video `gorm:"many2many:user_likes"`
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
	if err := u.db.Where("id in (select follower_id from user_follows where user_id = ?)", userId).Find(&users).Error; err != nil {
		return nil, err
	}
	// Todo: redis缓存
	return users, nil
}

// 获取用户的粉丝列表
func (u *UserModel) GetFollowerList(userId uint64) ([]*User, error) {
	var users []*User
	if err := u.db.Where("id in (select user_id from user_follows where follower_id = ?)", userId).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserModel) IsFollow(userId uint64, followerId uint64) (bool, error) {
	var count int64
	if err := u.db.Raw("SELECT COUNT(*) FROM user_follows WHERE user_id = ? AND follower_id = ?", userId, followerId).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// 关注一个用户
func (u *UserModel) CreateFollow(userId uint64, followId uint64) error {
	// 使用事务保证一致性
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("insert into user_follows (user_id, follower_id) values (?, ?)", userId, followId).Error; err != nil {
			return err
		}
		if err := tx.Model(User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count + 1")).Error; err != nil {
			return err
		}
		if err := tx.Model(User{}).Where("id = ?", followId).Update("follower_count", gorm.Expr("follower_count + 1")).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	// Todo: Redis缓存
	return nil
}

// 取消关注一个用户
func (u *UserModel) DeleteFollow(userId uint64, followId uint64) error {
	// 使用事务保证一致性
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if rows := tx.Exec("delete from user_follows where user_id = ? and follower_id = ?", userId, followId).RowsAffected; rows == 0 {
			return errors.New("关系不存在")
		}
		if err := tx.Model(User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count - 1")).Error; err != nil {
			return err
		}
		if err := tx.Model(User{}).Where("id = ?", followId).Update("follower_count", gorm.Expr("follower_count - 1")).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	// Todo: Redis缓存
	return nil
}
