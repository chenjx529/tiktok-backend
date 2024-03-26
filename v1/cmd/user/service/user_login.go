package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"tiktok-backend/dal/db"

	"tiktok-backend/kitex_gen/user"
	"tiktok-backend/pkg/errno"
)

type UserLoginService struct {
	ctx context.Context
}

// NewUserLoginService new UserLoginService
func NewUserLoginService(ctx context.Context) *UserLoginService {
	return &UserLoginService{
		ctx: ctx,
	}
}

// UserLogin check user info
func (s *UserLoginService) UserLogin(req *user.DouyinUserLoginRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.Username
	users, err := db.QueryUserByName(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}
	u := users[0]
	if u.Password != passWord {
		return 0, errno.AuthorizationFailedErr
	}
	return int64(u.ID), nil
}
