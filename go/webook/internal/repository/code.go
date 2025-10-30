package repository

import (
	"context"

	"github.com/miver02/Learn/go/webook/internal/repository/cache"
)


var (
	ErrCodeSendTooMany = cache.ErrCodeSendTooMany
)
	

type CodeRepository struct {
	cache *cache.CodeCache
}

func NewCodeRepository() *CodeRepository {
	return nil
}

func (repo *CodeRepository) Store(ctx context.Context, biz, phone, code string) error {
	return nil
}