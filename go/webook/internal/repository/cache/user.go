package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/miver02/learn-program/go/webook/internal/domain"
	"github.com/redis/go-redis/v9"
)

var ErrUserNotExist = redis.Nil

type UserCache struct {
	client     redis.Cmdable
	expiration time.Duration
}

func NewUserCache(client redis.Cmdable) *UserCache {
	return &UserCache{
		client:     client,
		expiration: time.Minute * 15,
	}
}

func (c *UserCache) Get(ctx context.Context, id int64) (domain.User, error) {
	key := c.Key(id)
	val, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return domain.User{}, err
	}
	var domain_user domain.User
	err = json.Unmarshal(val, &domain_user)
	return domain_user, err
}

func (c *UserCache) Set(ctx context.Context, u domain.User) error {
	val, err := json.Marshal(u)
	if err != nil {
		return err
	}
	key := c.Key(u.Id)

	return c.client.Set(ctx, key, val, c.expiration).Err()
}

func (cache *UserCache) Key(id int64) string {
	return fmt.Sprintf("user:info:%d", id)
}
