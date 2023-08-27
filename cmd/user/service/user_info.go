package service

import (
	"context"
	"errors"
	"tiktok-backend/cmd/user/pack"
	"tiktok-backend/dal/db"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/jwt"

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
	// 登录
	claims, err := jwt.GetclaimsFromTokenStr(req.Token)
	if err != nil {
		return nil, err
	}
	loginId := int64(int(claims[constants.IdentityKey].(float64)))

	users, err := db.MQueryUsersByIds(s.ctx, []int64{req.UserId})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("user not exist")
	}
	user1 := users[0]

	// TODO: 是否关注
	followSet, err := db.MQueryFollowByUserIdAndToUserIds(s.ctx, loginId, []int64{req.UserId})
	if err != nil {
		return nil, err
	}

	// 之前没有关注过
	isFollow := false
	if _, ok := followSet[req.UserId]; ok {
		isFollow = true
	}

	userInfo := pack.UserInfo(user1, isFollow)
	return userInfo, nil
}
