package service

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/miver02/Learn/go/webook/internal/repository"
	"github.com/miver02/Learn/go/webook/internal/service/sms"
)


const (
	codeTplId = "10086"
)

type CodeService struct {
	repo *repository.CodeRepository
	smsSvc sms.Service
	// tplId string
}

// Send发送验证码
func (svc *CodeService) Send(ctx context.Context, biz string, code string, phone string) error {
	// code := "1234"
	// serToRedis(code, key, time.Minute*10)
	// 两个步骤, 生成一个验证码
	code = svc.generateCode()
	// 塞进去 Redis
	err := svc.repo.Store(ctx, biz, phone, code)
	if err != nil {
		return nil
	}
	// 发送出去
	err = svc.smsSvc.Send(ctx, codeTplId, []string{code}, phone)
	return err
}

func (svc *CodeService) Verify(ctx context.Context, biz string, code string, phone string, inputCode string) (bool, error) {
	return svc.repo.Verify(ctx, biz, phone, inputCode)
}

func (svc *CodeService) generateCode() string {
	// 六位数: [0, 99999]
	num := rand.Intn(1000000)
	// 不够六位,加上前导0
	return fmt.Sprintf("%06d", num)
}