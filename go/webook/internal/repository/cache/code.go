package cache

import (
	"context"
	_ "embed"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	ErrCodeSendTooMany = errors.New("发送太频繁")
	ErrVerityTooMany   = errors.New("验证次数太多或者验证码已过期")
	ErrUnknown         = errors.New("未知错误")
	ErrCodeNotFound    = errors.New("验证码不存在或已过期")
)

// 编译器在编译的时候,会把set_code的代码放进luaSetCode变量里
// go.embed lua/verify_code.lua
var luaVerityCode string

//go:embed lua/set_code.lua
var luaSetCode string

type CodeCache struct {
	client redis.Cmdable
}

func NewCodeCache(client redis.Cmdable) *CodeCache {
	return &CodeCache{
		client: client,
	}
}

// 生成 Hash 表的 field（biz + phone，如 "login_sms:138xxxx8888"）
func (c *CodeCache) field(biz, phone string) string {
	return fmt.Sprintf("%s:%s", biz, phone)
}

// 验证码 Hash 表名（固定）
func (c *CodeCache) codeHashKey() string {
	return "phone_code"
}

// 错误次数 Hash 表名（固定）
func (c *CodeCache) cntHashKey() string {
	return "phone_code:cnt"
}

func (c *CodeCache) Set(ctx context.Context, biz, phone, code string) error {
	field := c.field(biz, phone)
	codeHashKey := c.codeHashKey()
	cntHashKey := c.cntHashKey()

	_, err := c.client.Eval(ctx, luaSetCode,
		[]string{codeHashKey, cntHashKey},
		field, code, 600,
	).Int()

	fmt.Printf("Redis 验证：Hash表=%s，字段=%s；错误次数Hash表=%s，字段=%s\n",
		codeHashKey, field, cntHashKey, field)
	return err
}

func (c *CodeCache) Verify(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	// 生成 Hash 字段（biz:phone）
	field := c.field(biz, phone)
	// 验证码 Hash 表名、错误次数 Hash 表名
	codeHashKey := c.codeHashKey()
	cntHashKey := c.cntHashKey()

	fmt.Printf("Redis 验证：Hash表=%s，字段=%s；错误次数Hash表=%s，字段=%s\n",
		codeHashKey, field, cntHashKey, field)

	// 调用 Lua 脚本：KEYS 传 2 个 Hash 表名，ARGV 传 field 和 inputCode
	res, err := c.client.Eval(ctx, luaVerityCode,
		[]string{codeHashKey, cntHashKey}, // KEYS[1] = phone_code, KEYS[2] = phone_code:cnt
		field, inputCode,                  // ARGV[1] = field, ARGV[2] = 输入的验证码
	).Int()

	if err != nil {
		return false, fmt.Errorf("Redis 执行脚本失败：%w", err)
	}

	switch res {
	case -1:
		return false, ErrCodeSendTooMany
	case 0:
		return true, nil
	case -2:
		return false, nil
	case -3:
		return false, ErrCodeNotFound // 验证码不存在（过期）
	default:
		return false, ErrUnknown
	}
}
