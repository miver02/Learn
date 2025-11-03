package retryable

import (
	"context"

	"github.com/miver02/Learn/go/webook/internal/service/sms"
)


type Service struct {
	svc sms.Service

	// 重试
	retryCnt int
}

func New(svc sms.Service, retryCnt int) *Service {
	return &Service{
		svc:      svc,
		retryCnt: retryCnt,
	}
}

func (s *Service) Send(ctx context.Context, tpl string, args []string, numbers...string) error {
	var err error
	for i := 0; i < s.retryCnt; i++ {
		err = s.svc.Send(ctx, tpl, args, numbers...)
		if err == nil {
			return nil
		}
	}
	return err
}