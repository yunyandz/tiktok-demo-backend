package service

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

func (s *Service) PublishComment(userId int64, videoId int64, content string) Response {
	return Response{}
}

func (s *Service) DeleteComment(userId int64, commentId int64) Response {
	return Response{}
}

func (s *Service) GetCommentList(userId int64, videoId int64) CommentListResponse {
	return CommentListResponse{}
}
