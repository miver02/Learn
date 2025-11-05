// user逻辑层:实现请求,返回数据
package repository

import (
	"context"

	"github.com/miver02/learn-program/go/webook/internal/domain"
	"github.com/miver02/learn-program/go/webook/internal/repository/cache"
	"github.com/miver02/learn-program/go/webook/internal/repository/dao"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound       = dao.ErrUserNotFound
)

type UserRepository struct {
	dao   *dao.UserDAO
	cache *cache.UserCache
}

func NewUserRepository(dao *dao.UserDAO, c *cache.UserCache) *UserRepository {
	return &UserRepository{
		dao:   dao,
		cache: c,
	}
}

func (r *UserRepository) InsertUserInfo(ctx context.Context, new_ud domain.User) error {
	err := r.dao.InsertUserInfo(ctx, new_ud)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	ud, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       ud.Id,
		Email:    ud.Email,
		Password: ud.Password,
	}, nil
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (r *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	u, err := r.cache.Get(ctx, id)
	if err == nil {
		// 缓存有数据
		return u, nil
	}
	// 缓存没数据
	if err == cache.ErrUserNotExist {
		// 去数据库加载

	}

	ud, err := r.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, nil
	}

	u = domain.User{
		Id:       ud.Id,
		Email:    ud.Email,
		Password: ud.Password,
	}

	go func() {
		err = r.cache.Set(ctx, u)
		if err != nil {
			// 打印日志
		}
	}()

	return u, err
	// 缓存出错了

}
