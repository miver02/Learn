// user服务层: 调用逻辑层,解决请求,返回数据给业务层
package service

import (
	"context"
	"errors"

	"github.com/miver02/Learn/go/webook/internal/domain"
	"github.com/miver02/Learn/go/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail
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


func (svc *UserService) Edit(ctx context.Context, new_udo domain.User) (error) {
	err := svc.repo.InsertUserInfo(ctx, new_udo)
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

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	// 考虑加密放在哪里
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	// 然后才存起来
	return svc.repo.Create(ctx, u)
}