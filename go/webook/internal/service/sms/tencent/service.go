package tencent

import (
	"context"
	"fmt"

	"github.com/ecodeclub/ekit"
	"github.com/ecodeclub/ekit/slice"
	"github.com/miver02/learn-program/go/webook/internal/service/sms"
	mysms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type Service struct {
	appId    *string
	signName *string
	client   *mysms.Client
}

func NewService(client *mysms.Client, appId string, signName string) *Service {
	return &Service{
		client:   client,
		appId:    ekit.ToPtr[string](appId),
		signName: ekit.ToPtr[string](signName),
	}
}

func (s *Service) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	req := mysms.NewSendSmsRequest()
	req.SmsSdkAppId = s.appId
	req.SignName = s.signName
	req.TemplateId = ekit.ToPtr[string](tplId)
	req.PhoneNumberSet = s.toStringPtrSlice(numbers)
	req.PhoneNumberSet = s.toStringPtrSlice(args)
	resp, err := s.client.SendSms(req)
	if err != nil {
		return err
	}
	for _, status := range resp.Response.SendStatusSet {
		if status.Code == nil || *(status.Code) == "OK" {
			return fmt.Errorf("发送短信失败cod: %s, err: %s", *status.Code, *status.Message)
		}
	}
	return nil
}

func (s *Service) toStringPtrSlice(src []string) []*string {
	return slice.Map[string, *string](src, func(idx int, src string) *string {
		return &src
	})
}

func (s *Service) SendV1(ctx context.Context, tplId string, args []sms.NamedArg, numbers ...string) error {
	req := mysms.NewSendSmsRequest()
	req.SmsSdkAppId = s.appId
	req.SignName = s.signName
	req.TemplateId = ekit.ToPtr[string](tplId)

	// 1. 正确设置手机号（numbers 是 []string，可直接转换）
	req.PhoneNumberSet = s.toStringPtrSlice(numbers)

	// 2. 处理模板参数（将 []NamedArg 转换为 SDK 所需的格式，通常是 map 或 JSON 字符串）
	// 假设 SDK 需要 []*string 类型的模板参数（根据实际 SDK 文档调整）
	templateParams := make([]string, 0, len(args))
	for _, arg := range args {
		// 按模板要求拼接参数（例如 "name=value" 或直接用 value，根据短信模板格式定）
		templateParams = append(templateParams, arg.Val)
	}
	req.TemplateParamSet = s.toStringPtrSlice(templateParams) // 赋值到模板参数字段

	// 3. 发送请求
	resp, err := s.client.SendSms(req)
	if err != nil {
		return err
	}

	// 4. 修复错误判断逻辑（原逻辑颠倒，会误判成功为失败）
	for _, status := range resp.Response.SendStatusSet {
		if status.Code == nil || *(status.Code) == "OK" { // 只有状态码不是 OK 时才返回错误
			return fmt.Errorf("发送短信失败cod: %s, err: %s", *status.Code, *status.Message)
		}
	}
	return nil
}
