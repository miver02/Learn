package service

import "context"


type CodeService struct {

}

// Send发送验证码
func (svc *CodeService) Send(ctx context.Context, biz string, code string, phone string) error {
	return nil
}

func (svc *CodeService) Verify(ctx context.Context, biz string, code string, phone string, inputCode string) (bool, error) {
	return true, nil
}