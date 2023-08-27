package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"tiktok-backend/cmd/relation/pack"
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/relation"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/jwt"
)

type RelationFollowerListService struct {
	ctx context.Context
}

// NewRelationFollowerListService new RelationFollowListService
func NewRelationFollowerListService(ctx context.Context) *RelationFollowerListService {
	return &RelationFollowerListService{ctx: ctx}
}

// RelationFollowerList 获取粉丝列表，你不一定关注粉丝的
func (s *RelationFollowerListService) RelationFollowerList(req *relation.DouyinRelationFollowerListRequest) ([]*relation.User, error) {
	// 登录id
	claims, err := jwt.GetclaimsFromTokenStr(req.Token)
	if err != nil {
		return nil, err
	}
	loginId := int64(int(claims[constants.IdentityKey].(float64)))

	// 获取粉丝id
	followers, err := db.QueryFollowerByToUserId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	klog.Info(len(followers))
	followerIds := make([]int64, 0)
	for _, follower := range followers {
		followerIds = append(followerIds, follower.UserId) // 粉丝的id
	}

	// 利用toUserIds获取toUser信息
	followerUsers, err := db.MQueryUsersByIds(s.ctx, followerIds)
	if err != nil {
		return nil, err
	}

	// 获取粉丝Set
	followerSet, err := db.MQueryFollowByUserIdAndToUserIds(s.ctx, loginId, followerIds)
	if err != nil {
		return nil, err
	}

	userList := pack.BuildFollowList(followerUsers, followerSet)
	userList = nil
	return userList, nil
}
