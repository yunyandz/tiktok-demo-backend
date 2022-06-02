package service

import (
	"github.com/yunyandz/tiktok-demo-backend/internal/errorx"
	"github.com/yunyandz/tiktok-demo-backend/internal/logger"
	"github.com/yunyandz/tiktok-demo-backend/internal/model"
)

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

type CommentActionResponse struct {
	Response
	Comment Comment `json: "comment"`
}

func (s *Service) PublishComment(userId uint64, videoId uint64, content string) (*CommentActionResponse, error) {
	videoModel := model.NewVideoModel(s.db, s.rds)
	comment, err := videoModel.CreateAComment(videoId, userId, content)

	// Video数据库写入评论失败
	if err != nil {
		logger.Suger().Errorw("PublishComment CreateComment failed", "err", err.Error())
		return nil, errorx.ErrInternalBusy
	}
	// User数据库查找评论用户信息失败
	r, err := s.GetUserInfo(userId)
	if err != nil {
		return nil, errorx.ErrUserDoesNotExists
	}

	rsp := CommentActionResponse{
		Response: r.Response,
		Comment: Comment{
			Id:         uint64(comment.Model.ID),
			User:       r.User,
			Content:    content,
			CreateDate: comment.Model.CreatedAt.Format("01-02"), //格式化为mm-dd
		},
	}
	return &rsp, nil
}

func (s *Service) DeleteComment(commentId uint64) (*CommentActionResponse, error) {
	videoModel := model.NewVideoModel(s.db, s.rds)
	comment, err := videoModel.FindAComment(commentId)
	if err != nil {
		return nil, errorx.ErrCommentDoesNotExists
	}

	err = videoModel.DeleteAComment(commentId)
	// Video数据库删除评论失败
	if err != nil {
		logger.Suger().Errorw("DeleteComment DeleteAComment failed", "err", err.Error())
		return nil, errorx.ErrInternalBusy
	}

	r, err := s.GetUserInfo(comment.UserID)
	// User数据库查找用户失败
	if err != nil {
		return nil, errorx.ErrUserDoesNotExists
	}

	rsp := CommentActionResponse{
		Response: r.Response,
		Comment: Comment{
			Id:         uint64(comment.Model.ID),
			User:       r.User,
			Content:    comment.Content,
			CreateDate: comment.Model.CreatedAt.Format("01-02"), //格式化为mm-dd
		},
	}
	return &rsp, nil
}

func (s *Service) GetCommentList(videoId uint64) (*CommentListResponse, error) {
	videoModel := model.NewVideoModel(s.db, s.rds)
	comments, err := videoModel.GetVideoComments(videoId)
	if err != nil {
		return nil, errorx.ErrCommentDoesNotExists
	}

	var commentList []Comment
	for i := 0; i < len(comments); i++ {
		commentA := comments[i]
		r, err := s.GetUserInfo(commentA.UserID)
		// User数据库查找用户失败
		if err != nil {
			return nil, errorx.ErrUserDoesNotExists
		}

		commentB := Comment{
			Id:         uint64(commentA.Model.ID),
			User:       r.User,
			Content:    commentA.Content,
			CreateDate: commentA.Model.CreatedAt.Format("01-02"),
		}
		commentList = append(commentList, commentB)
	}

	rsp := CommentListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "ok",
		},
		CommentList: commentList,
	}

	return &rsp, nil
}
