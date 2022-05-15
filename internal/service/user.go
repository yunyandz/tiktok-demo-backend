package service

type User struct {
	Id            uint64 `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   uint32 `json:"follow_count,omitempty"`
	FollowerCount uint32 `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

type UserLoginResponse struct {
	Response
	UserId uint64 `json:"user_id,omitempty"`
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

func (s *Service) Follow(userId uint64, to_userId uint64) Response {
	return Response{}
}

func (s *Service) GetFollowList(userId uint64) UserListResponse {
	return UserListResponse{}
}

func (s *Service) GetFollowerList(userId uint64) UserListResponse {
	return UserListResponse{}
}

func (s *Service) GetUserInfo(userId uint64) UserResponse {
	return UserResponse{}
}
