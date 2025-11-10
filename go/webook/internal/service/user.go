// user服务层: 调用逻辑层,解决请求,返回数据给业务层
package service

import (
	"context"

	"github.com/miver02/learn-program/go/webook/internal/consts"
	"github.com/miver02/learn-program/go/webook/internal/domain"
	"github.com/miver02/learn-program/go/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) Edit(ctx context.Context, domain_user domain.User) error {
	// 为密码加密
	if domain_user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(domain_user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		domain_user.Password = string(hash)
	}
	// 插入
	err := svc.repo.InsertUserInfo(ctx, domain_user)
	if err != nil {
		return err
	}
	return nil
}

func (svc *UserService) Login(ctx context.Context, new_udo domain.User) (domain.User, error) {
	// 找用户
	datas_u, err := svc.repo.FindByEmail(ctx, new_udo.Email)
	if err == consts.ErrUserNotFound {
		return datas_u, consts.ErrInvalidUserOrPassword
	}
	if err != nil {
		return datas_u, err
	}

	// 比较密码
	err = bcrypt.CompareHashAndPassword([]byte(datas_u.Password), []byte(new_udo.Password))
	if err != nil {
		// debug日志
		return datas_u, consts.ErrInvalidUserOrPassword
	}
	return datas_u, nil
}

func (svc *UserService) SignUp(ctx context.Context, domain_user domain.User) (domain.User, error) {
	// 考虑加密放在哪里
	hash, err := bcrypt.GenerateFromPassword([]byte(domain_user.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}
	domain_user.Password = string(hash)
	// 然后才存起来
	return svc.repo.Create(ctx, domain_user)
}

func (svc *UserService) Profile(ctx context.Context, id int64) (domain.User, error) {
	u, err := svc.repo.FindEmailById(ctx, id)
	if err != nil {

	}
	return u, err
}

func (svc *UserService) FindOrCreate(ctx context.Context, phone string) (domain.User, error) {
	domain_user, err := svc.repo.FindByPhone(ctx, phone)
	switch {
	case domain_user.Id != 0 && err == nil:
		// 用户存在,无报错
		return domain_user, nil
	case domain_user.Id == 0 && err != consts.ErrUserNotFound:
		// 用户存在, 但报错
		return domain.User{}, err
	}
	// 系统资源不足的时候,出发降级了,不创建用户
	// if ctx.Value("降级") == "true" {
	// 	return domain.User{}, errors.New("系统降级了")
	// }
	domain_user, err = svc.repo.Create(ctx, domain.User{
		Phone: phone,
	})
	if err != nil {
		return domain.User{}, err
	}
	return domain_user, nil
}
