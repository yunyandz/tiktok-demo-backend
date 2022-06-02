package service

import (
	"github.com/yunyandz/tiktok-demo-backend/internal/errorx"
	"github.com/yunyandz/tiktok-demo-backend/internal/jwtx"
	"github.com/yunyandz/tiktok-demo-backend/internal/logger"
	"github.com/yunyandz/tiktok-demo-backend/internal/model"
	"github.com/yunyandz/tiktok-demo-backend/internal/scryptx"
)

type User struct {
	//Id            uint64 `json:"id,omitempty"`
	//Name          string `json:"name,omitempty"`
	//FollowCount   uint32 `json:"follow_count,omitempty"`
	//FollowerCount uint32 `json:"follower_count,omitempty"`

	ID            uint64 `json:"id"`
	Username      string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

type UserLoginResponse struct {
	Response
	UserID uint64 `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserRegisterResponse struct {
	Response
	UserID uint64 `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func (s *Service) Register(username string, password string) (*UserRegisterResponse, error) {

	var res UserRegisterResponse
	userModel := model.NewUserModel(s.db, s.rds)
	user, err := userModel.GetUserByName(username)
	if err == nil {
		if user.ID > 0 {
			return nil, errorx.ErrUserAlreadyExists
		}
		return nil, err
	}

	u := model.User{
		Username: username,
		Password: scryptx.PasswordEncrypt(password),
	}

	id, err := userModel.CreateUser(&u)
	if err != nil {
		logger.Suger().Errorw("Register save failed", "err", err.Error())
		return nil, errorx.ErrInternalBusy
	}

	token, err := jwtx.CreateUserClaims(jwtx.UserInfo{
		Username: username,
		UserID:   uint64(id),
	})
	if err != nil {
		logger.Suger().Errorw("Register CreateUserClaims failed", "err", err.Error())
		return nil, errorx.ErrInternalBusy
	}

	res.Response = Response{
		StatusCode: 0,
		StatusMsg:  "ok",
	}
	res.UserID = id
	res.Token = token

	return &res, nil
}

func (s *Service) Login(username string, password string) (*UserLoginResponse, error) {

	userModel := model.NewUserModel(s.db, s.rds)
	user, err := userModel.GetUserByName(username)
	if err != nil {
		return nil, errorx.ErrUserDoesNotExists
	}

	if !scryptx.PasswordValidate(password, user.Password) {
		return nil, errorx.ErrUserPassword
	}

	token, err := jwtx.CreateUserClaims(jwtx.UserInfo{
		Username: username,
		UserID:   uint64(user.ID),
	})
	if err != nil {
		logger.Suger().Errorw("Login CreateUserClaims failed", "err", err.Error())
		return nil, errorx.ErrInternalBusy
	}

	rsp := UserLoginResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		UserID: uint64(user.ID),
		Token:  token,
	}

	return &rsp, nil
}

func (s *Service) Follow(userId uint64, to_userId uint64) Response {
	userModel := model.NewUserModel(s.db, s.rds)
	err := userModel.CreateFollow(userId, to_userId)
	if err != nil {
		return Response{StatusCode: 1, StatusMsg: err.Error()}
	}
	return Response{StatusCode: 0, StatusMsg: "Follow succeed"}
}

func (s *Service) UnFollow(userId uint64, to_userId uint64) Response {
	userModel := model.NewUserModel(s.db, s.rds)
	err := userModel.DeleteFollow(userId, to_userId)
	if err != nil {
		return Response{StatusCode: 1, StatusMsg: err.Error()}
	}
	return Response{StatusCode: 0, StatusMsg: "UnFollow succeed"}
}

func (s *Service) GetFollowList(userId uint64) UserListResponse {
	userModel := model.NewUserModel(s.db, s.rds)
	followList, err := userModel.GetFollowList(userId)
	if err != nil {
		return UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		}
	}
	var users []User
	var user *User
	for _, item := range followList {
		user = &User{}
		user.ID = uint64(item.ID)
		user.Username = item.Username
		user.FollowCount = item.FollowCount
		user.FollowerCount = item.FollowerCount
		user.IsFollow = true
		users = append(users, *user)
	}
	return UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: users,
	}
}

func (s *Service) GetFollowerList(userId uint64) UserListResponse {
	userModel := model.NewUserModel(s.db, s.rds)
	followList, err := userModel.GetFollowerList(userId)
	if err != nil {
		return UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		}
	}
	var users []User
	var user *User
	for _, item := range followList {
		user = &User{}
		user.ID = uint64(item.ID)
		user.Username = item.Username
		user.FollowCount = item.FollowCount
		user.FollowerCount = item.FollowerCount
		user.IsFollow = userModel.IsFollow(uint64(item.ID), userId)
		users = append(users, *user)
	}
	return UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: users,
	}
}

func (s *Service) GetUserInfo(UserID uint64) (*UserResponse, error) {
	userModel := model.NewUserModel(s.db, s.rds)
	user, err := userModel.GetUser(UserID)
	if err != nil {
		return nil, errorx.ErrUserDoesNotExists
	}

	// TODO 需要查follow表
	// favoriteModel.IsFollow()

	rsp := UserResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		User: User{
			ID:            uint64(user.Model.ID),
			Username:      user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			// TODO 需要查follow表
			IsFollow: false,
		},
	}
	return &rsp, nil
}
