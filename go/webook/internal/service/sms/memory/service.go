package memory

import (
	"context"
	"fmt"

	"github.com/miver02/Learn/go/webook/internal/service/sms"
)


type Service struct {
}

func NewService() sms.Service {
	return nil
}

// func NewService() *Service {
// 	return &Service{}
// }

func (s *Service) Send(ctx context.Context, tplId string, args []string, numbers...string) error {
	fmt.Println(args)
	return nil
}
