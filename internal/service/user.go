package service

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func (s *Service) Register(username string, password string) UserLoginResponse {
	return UserLoginResponse{}
}

func (s *Service) Login(username string, password string) UserLoginResponse {
	return UserLoginResponse{}
}

func (s *Service) Follow(userId int64, to_userId int64) Response {
	return Response{}
}

func (s *Service) GetFollowList(userId int64) UserListResponse {
	return UserListResponse{}
}

func (s *Service) GetFollowerList(userId int64) UserListResponse {
	return UserListResponse{}
}

func (s *Service) GetUserInfo(userId int64) UserResponse {
	return UserResponse{}
}
