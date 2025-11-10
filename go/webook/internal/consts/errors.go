package consts

import (
	"errors"

	"gorm.io/gorm"
)

// 短信验证错误
var (
	ErrCodeSendTooMany      = errors.New("验证码发送太频繁")
	ErrPhoneCodeSendTooMany = errors.New("手机验证码发送太频繁,请一小时后重试")
	ErrVerityTooMany        = errors.New("验证失败,并且验证次数已用完,请60s后再重试")
	ErrCodeNotFound         = errors.New("验证码不存在或已过期")
)

var (
	ErrJwtTokenInvalid = errors.New("无效的token")
)

var (
	ErrUnknown = errors.New("未知错误")
	ErrSystem  = errors.New("系统错误")
)

var (
	ErrUserDuplicate = errors.New("邮箱或者手机号冲突")
	ErrUserNotFound  = gorm.ErrRecordNotFound
	ErrInvalidUserOrPassword = errors.New("邮箱或者密码不对")
)