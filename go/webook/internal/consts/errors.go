package consts

import "errors"

var (
	ErrCodeSendTooMany      = errors.New("验证码发送太频繁")
	ErrPhoneCodeSendTooMany = errors.New("手机验证码发送太频繁,请一小时后重试")
	ErrVerityTooMany        = errors.New("验证失败,并且验证次数已用完,请60s后再重试")
	ErrUnknown              = errors.New("未知错误")
	ErrCodeNotFound         = errors.New("验证码不存在或已过期")
)