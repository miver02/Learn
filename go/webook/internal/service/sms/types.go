package sms

import "context"


type Service interface {
	Send(ctx context.Context, tpl string, args []string, numbers...string) error
	SendV1(ctx context.Context, tpl string, args []NamedArg, numbers...string) error
}

type NamedArg struct {
	Val string
	Name string
}