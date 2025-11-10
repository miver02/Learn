// user逻辑层:实现请求,返回数据
package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/miver02/learn-program/go/webook/internal/domain"
	"github.com/miver02/learn-program/go/webook/internal/repository/cache"
	"github.com/miver02/learn-program/go/webook/internal/repository/dao"
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

func (r *UserRepository) InsertUserInfo(ctx context.Context, domain_user domain.User) error {
	err := r.dao.InsertUserInfo(ctx, r.domainToDao(domain_user))
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return r.daoToDomain(user), nil
}

func (r *UserRepository) Create(ctx context.Context, domain_user domain.User) (domain.User, error) {
	user, err := r.dao.Insert(ctx, r.domainToDao(domain_user))
	return r.daoToDomain(user), err
}

func (r *UserRepository) FindEmailById(ctx context.Context, id int64) (domain.User, error) {
	domain_user, err := r.cache.Get(ctx, id)
	if err == nil {
		// 缓存有数据
		return domain_user, nil
	}
	// 缓存没数据
	if err == cache.ErrUserNotExist {
		// 去数据库加载

	}

	user, err := r.dao.FindEmailById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	domain_user = r.daoToDomain(user)

	go func() {
		err = r.cache.Set(ctx, domain_user)
		if err != nil {
			// 打印日志
		}
	}()

	return domain_user, err
	// 缓存出错了

}

func (r *UserRepository) FindByPhone(ctx context.Context, phone string) (domain.User, error) {
	user, err := r.dao.FindByPhone(ctx, phone)
	if err != nil {
		return domain.User{}, err
	}
	return r.daoToDomain(user), nil
}

func (r *UserRepository) domainToDao(domain_user domain.User) dao.User {
	// 1. 处理 Email：string → sql.NullString
	var email, phone sql.NullString
	if domain_user.Email != "" {
		email.String = domain_user.Email // 赋值实际字符串
		email.Valid = true               // 标记为“非 NULL”
	} else {
		// 若 domain 的 Email 为空字符串，根据业务需求决定：
		// 方案1：设为 Valid=false（数据库存储为 NULL）→ 推荐（区分空字符串和 NULL）
		email.Valid = false
		// 方案2：设为 Valid=true 且 String=""（数据库存储为空字符串）
		// email.String = ""
		// email.Valid = true
	}

	if domain_user.Phone != "" {
		phone.String = domain_user.Phone
		phone.Valid = true
	} else {
		phone.Valid = false
	}

	// 2. 处理 Ctime：time.Time → int64（毫秒时间戳）
	// 注意：若 Ctime 是“零时间”（time.Time{}），需避免存储无效时间戳（可选处理）
	ctime := domain_user.Ctime.UnixMilli() // 转为毫秒级时间戳

	// 3. 其他字段直接赋值，返回 dao.User
	return dao.User{
		Id:           domain_user.Id,
		Name:         domain_user.Name,
		Introduction: domain_user.Introduction,
		Email:        email, // 已转换为 sql.NullString
		Phone:        phone,
		Password:     domain_user.Password,
		Ctime:        ctime, // 已转换为 int64 毫秒时间戳
	}
}
func (r *UserRepository) daoToDomain(user dao.User) domain.User {
	email, phone := "", ""
	if user.Email.Valid { // 只有数据库非 NULL 时，才取实际值
		email = user.Email.String
	}
	if user.Phone.Valid {
		phone = user.Phone.String
	}
	return domain.User{
		Id:           user.Id,
		Name:         user.Name,
		Introduction: user.Introduction,
		Email:        email,
		Phone:        phone,
		Password:     user.Password,
		Ctime:        time.UnixMilli(user.Ctime),
	}
}
