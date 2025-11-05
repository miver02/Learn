package repository

import (
	"context"

	"github.com/miver02/learn-program/go/webook/internal/consts"
	"github.com/miver02/learn-program/go/webook/internal/repository/cache"
)

var (
	ErrCodeSendTooMany = consts.ErrCodeSendTooMany
	ErrVerityTooMany   = consts.ErrVerityTooMany
)

type CodeRepository struct {
	cache *cache.CodeCache
}

func NewCodeRepository(codeCache *cache.CodeCache) *CodeRepository {
	return &CodeRepository{
		cache: codeCache,
	}
}

func (repo *CodeRepository) Store(ctx context.Context, phone, code string) error {
	return repo.cache.Set(ctx, phone, code)
}

func (repo *CodeRepository) Verify(ctx context.Context, phone, inputCode string) (bool, error) {
	return repo.cache.Verify(ctx, phone, inputCode)
}
