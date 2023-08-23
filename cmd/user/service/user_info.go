package service

import (
	"context"
	"errors"
	"tiktok-backend/cmd/user/pack"

	"tiktok-backend/cmd/user/dal/db"
	"tiktok-backend/kitex_gen/user"
)

type UserInfoService struct {
	ctx context.Context
}

// NewUserInfoService new UserInfoService
func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{
		ctx: ctx,
	}
}

// UserInfo get user info
func (s *UserInfoService) UserInfo(req *user.DouyinUserRequest) (*user.User, error) {
	userIds := []int64{req.UserId}
	users, err := db.QueryUserByIds(s.ctx, userIds)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("user not exist")
	}
	user1 := users[0]

	// TODO: 互动部分还没有实现

	userInfo := pack.UserInfo(user1, false)
	return userInfo, nil
}
