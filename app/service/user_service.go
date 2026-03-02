package service

import (
	"context"
	"fmt"

	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/core/token"
	"iflow-lite/dao"
	"iflow-lite/type/input"
	"iflow-lite/type/model"

	"golang.org/x/crypto/bcrypt"
)

var DefaultUserService = NewUserService()

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) UserGet(ctx context.Context, in *input.UserGetInput) (interface{}, error) {
	return dao.DefaultUserDao.UserGet(ctx, in.ID)
}

func (this *UserService) UserAdd(ctx context.Context, in *input.UserAddInput) (interface{}, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.ServiceLogger.Errorf("bcrypt generate from password error: %+v", err)
		return nil, err
	}

	m := &model.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: string(pwd),
	}
	if err := dao.DefaultUserDao.UserAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *UserService) UserLogin(ctx context.Context, in *input.UserLoginInput) (interface{}, error) {
	user, err := dao.DefaultUserDao.UserTake(ctx, in.Email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return nil, err
	}
	jwt, err := token.GenerateJWT(fmt.Sprintf("%d", user.ID))
	return jwt, err
}
