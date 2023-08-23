package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"tiktok-backend/cmd/user/dal/db"
	"tiktok-backend/kitex_gen/user"
	"tiktok-backend/pkg/errno"
)

type UserRegisterService struct {
	ctx context.Context
}

// NewUserRegisterService new UserRegisterService
func NewUserRegisterService(ctx context.Context) *UserRegisterService {
	return &UserRegisterService{
		ctx: ctx,
	}
}

// UserRegister register user info
func (s *UserRegisterService) UserRegister(req *user.DouyinUserRegisterRequest) (int64, error) {
	users, err := db.QueryUserByName(s.ctx, req.Username)
	if err != nil {
		return 0, err
	}
	if len(users) != 0 {
		return 0, errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))

	userId, err := db.CreateUser(s.ctx, &db.User{
		Name:     req.Username,
		Password: password,
	})
	if err != nil {
		return 0, err
	}

	return userId, nil

}
