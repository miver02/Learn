// user服务层: 调用逻辑层,解决请求,返回数据给业务层
package service

import (
	"context"
	"errors"

	"github.com/miver02/learn-program/go/webook/internal/domain"
	"github.com/miver02/learn-program/go/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserDuplicateEmail    = repository.ErrUserDuplicateEmail
	ErrInvalidUserOrPassword = errors.New("邮箱或者密码不对")
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
	err := svc.repo.InsertUserInfo(ctx, domain_user)
	if err != nil {
		return err
	}
	return nil
}

func (svc *UserService) Login(ctx context.Context, new_udo domain.User) (domain.User, error) {
	// 找用户
	datas_u, err := svc.repo.FindByEmail(ctx, new_udo.Email)
	if err == repository.ErrUserNotFound {
		return datas_u, ErrInvalidUserOrPassword
	}
	if err != nil {
		return datas_u, err
	}

	// 比较密码
	err = bcrypt.CompareHashAndPassword([]byte(datas_u.Password), []byte(new_udo.Password))
	if err != nil {
		// debug日志
		return datas_u, ErrInvalidUserOrPassword
	}
	return datas_u, nil
}

func (svc *UserService) SignUp(ctx context.Context, domain_user domain.User) error {
	// 考虑加密放在哪里
	hash, err := bcrypt.GenerateFromPassword([]byte(domain_user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
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
	if err != repository.ErrUserNotFound {
		return domain_user, err
	}
	err = svc.repo.Create(ctx, domain.User{
		Phone: phone,
	})
	if err != nil {
		return domain_user, err
	}
	return svc.repo.FindByPhone(ctx, phone)
}
