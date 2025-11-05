package cache

import (
	"context"
	_ "embed"
	"github.com/miver02/Learn/go/webook/internal/consts"
	"fmt"
	"github.com/redis/go-redis/v9"
)

const (
	codeKeyPrefix  = "phone_code"
	cntKeyPrefix   = "cnt"
	blackKeyPrefix = "black"
	biz            = "login_sms"
)

// 编译器在编译的时候,会把set_code的代码放进luaSetCode变量里
//
//go:embed lua/verify_code.lua
var luaVerityCode string

//go:embed lua/set_code.lua
var luaSetCode string

type CodeCache struct {
	client redis.Cmdable
}

func NewCodeCache(client redis.Cmdable) *CodeCache {
	// 打印脚本内容，确认是否正确加载
	// fmt.Printf("验证脚本长度: %d\n", len(luaVerityCode))
	// fmt.Printf("设置脚本长度: %d\n", len(luaSetCode))

	return &CodeCache{
		client: client,
	}
}

// 验证码 key（固定）
func (c *CodeCache) codeKey(phone string) string {
	return fmt.Sprintf("%s:%s:%s", codeKeyPrefix, biz, phone)
}

// 验证码 key（固定）
func (c *CodeCache) cntKey(phone string) string {
	return fmt.Sprintf("%s:%s:%s", cntKeyPrefix, biz, phone)
}

// 验证码 key（固定）
func (c *CodeCache) blackKey(phone string) string {
	return fmt.Sprintf("%s:%s:%s", blackKeyPrefix, biz, phone)
}

func (c *CodeCache) Set(ctx context.Context, phone, code string) error {
	codeKey := c.codeKey(phone)   // "phone_code:login_sms:12345678922"
	cntKey := c.cntKey(phone)     // "cnt:login_sms:12345678922"
	blackKey := c.blackKey(phone) // "black:login_sms:12345678922"

	// 执行存储脚本：参数顺序必须与脚本期望一致
	// KEYS：[codeKey, cntKey, blackKey]
	// ARGV：[code, expireSec]
	res, err := c.client.Eval(ctx, luaSetCode,
		[]string{codeKey, cntKey, blackKey}, // KEYS[1]=codeKey, KEYS[2]=cntKey, KEYS[3]=blackKey
		code, 60,                            // ARGV[1]=code, ARGV[2]=60
	).Int()

	if err != nil {
		fmt.Printf("存储脚本执行失败：%v\n", err)
		return err
	}

	switch res {
	case 0:
		return nil // 成功
	case -1:
		return consts.ErrCodeSendTooMany // 发送过于频繁
	case -2:
		return consts.ErrPhoneCodeSendTooMany
	default:
		return consts.ErrUnknown
	}

}

// 验证验证码
func (c *CodeCache) Verify(ctx context.Context, phone, inputCode string) (bool, error) {
	codeKey := c.codeKey(phone)   // "phone_code:login_sms:12345678922"
	cntKey := c.cntKey(phone)     // "cnt:login_sms:12345678922"
	blackKey := c.blackKey(phone) // "black:login_sms:12345678922"

	// 执行验证脚本：参数顺序与脚本一致
	res, err := c.client.Eval(ctx, luaVerityCode,
		[]string{codeKey, cntKey, blackKey}, // KEYS[1]=codeKey, KEYS[2]=cntKey, KEYS[3]=blackKey
		inputCode,                           // ARGV[1]=inputCode
	).Int()

	if err != nil {
		return false, fmt.Errorf("Redis 执行脚本失败：%w", err)
	}

	switch res {
	case -1:
		return false, consts.ErrVerityTooMany // 错误次数用尽
	case -3:
		return false, consts.ErrCodeNotFound // 验证码不存在或输入错误
	case 0:
		return true, nil // 验证成功
	default:
		return false, consts.ErrUnknown // 未知错误
	}
}
