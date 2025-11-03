package repository

import (
	"context"

	"github.com/miver02/Learn/go/webook/internal/repository/cache"
)


var (
	ErrCodeSendTooMany = cache.ErrCodeSendTooMany
	ErrVerityTooMany = cache.ErrVerityTooMany
)
	

type CodeRepository struct {
	cache *cache.CodeCache
}

func NewCodeRepository() *CodeRepository {
	return nil
}

func (repo *CodeRepository) Store(ctx context.Context, biz, phone, code string) error {
	return repo.cache.Set(ctx, biz, phone, code)
}

func (repo *CodeRepository) Verify(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	return repo.cache.Verify(ctx, biz, phone, inputCode)
}