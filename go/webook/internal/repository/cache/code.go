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
	ErrVerityTooMany   = errors.New("验证失败,并且验证次数已用完,请60s后再重试")
	ErrUnknown         = errors.New("未知错误")
	ErrCodeNotFound    = errors.New("验证码不存在或已过期")
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
	return "cnt"
}

func (c *CodeCache) Set(ctx context.Context, biz, phone, code string) error {
	// 生成 fieldPrefix：login_sms:12345678922（biz:phone）
	fieldPrefix := c.field(biz, phone)
	codeHashKey := c.codeHashKey() // "phone_code"（正确）
	cntHashKey := c.cntHashKey()   // "cnt"（已修正，与验证一致）

	// 执行存储脚本：参数顺序必须与脚本期望一致
	// KEYS：[phone_code, cnt]
	// ARGV：[fieldPrefix, code, expireSec]
	_, err := c.client.Eval(ctx, luaSetCode,
		[]string{codeHashKey, cntHashKey}, // KEYS[1]=phone_code, KEYS[2]=cnt
		fieldPrefix, code, 600,            // ARGV[1]=fieldPrefix, ARGV[2]=code, ARGV[3]=600
	).Int()

	if err != nil {
		fmt.Printf("存储脚本执行失败：%v\n", err)
		return err
	}
	return nil
}

// 验证验证码
func (c *CodeCache) Verify(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	fieldPrefix := c.field(biz, phone) // "login_sms:12345678922"
	codeHashKey := c.codeHashKey()     // "phone_code"（与存储一致）
	cntHashKey := c.cntHashKey()       // "cnt"（与存 储一致）

	// 执行验证脚本：参数顺序与脚本一致
	res, err := c.client.Eval(ctx, luaVerityCode,
		[]string{codeHashKey, cntHashKey}, // KEYS[1]=phone_code, KEYS[2]=cnt
		fieldPrefix, inputCode,            // ARGV[1]=fieldPrefix, ARGV[2]=inputCode
	).Int()

	if err != nil {
		return false, fmt.Errorf("Redis 执行脚本失败：%w", err)
	}

	switch res {
	case -1:
		return false, ErrVerityTooMany // 错误次数用尽
	case -3:
		return false, ErrCodeNotFound // 验证码不存在或输入错误
	case 0:
		return true, nil // 验证成功
	default:
		return false, ErrUnknown // 未知错误
	}
}
