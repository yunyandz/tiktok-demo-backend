package errorx

import "errors"

var (
	ErrInternalBusy = errors.New("内部出错")
	ErrTokenMethod  = errors.New("token签名方法错误")
	ErrInvalidToken = errors.New("无效的token")

	ErrUserAlreadyExists = errors.New("用户已存在")
	ErrUserDoesNotExists = errors.New("用户不存在")
	ErrUserPassword      = errors.New("密码错误")

	ErrUserOffline          = errors.New("请先登录")
	ErrCommentDoesNotExists = errors.New("评论不存在")
	ErrReadVideo            = errors.New("读取视频失败")
	ErrNotVideo             = errors.New("不是视频文件")
	ErrUserAlreadyLikeVideo = errors.New("用户已经点赞过该视频")
	ErrUserNotLikeVideo     = errors.New("用户没有点赞过该视频")

	ErrPermissionDenied = errors.New("权限不足")
	ErrCommentTextEmpty = errors.New("评论内容不能为空")
	ErrCommentIDEmpty   = errors.New("评论ID不能为空")
)
