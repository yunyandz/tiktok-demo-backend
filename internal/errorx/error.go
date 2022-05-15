package errorx

import "errors"

var (
	ErrInternalBusy = errors.New("内部出错")
	ErrTokenMethod  = errors.New("token签名方法错误")
	ErrInvalidToken = errors.New("无效的token")

	ErrUserAlreadyExists = errors.New("用户已存在")
	ErrUserDoesNotExists = errors.New("用户不存在")
	ErrUserPassword      = errors.New("密码错误")
)