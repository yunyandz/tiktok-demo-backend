package service

type Comment struct {
	Id         uint64 `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

func (s *Service) PublishComment(userId uint64, videoId uint64, content string) Response {
	return Response{}
}

func (s *Service) DeleteComment(userId uint64, commentId uint64) Response {
	return Response{}
}

func (s *Service) GetCommentList(userId uint64, videoId uint64) CommentListResponse {
	return CommentListResponse{}
}
