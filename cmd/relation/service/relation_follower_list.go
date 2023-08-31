package service

import (
	"context"
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
	// 登录用户的id是loginId
	claims, err := jwt.GetClaimsFromTokenStr(req.Token)
	if err != nil {
		return nil, err
	}
	loginId := int64(claims[constants.IdentityKey].(float64))

	// 通过UserId，获取目标人物的粉丝关系followers:   followerid -> UserId
	followers, err := db.QueryFollowerByToUserId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	followerIds := make([]int64, 0) // 粉丝的id
	for _, follower := range followers {
		followerIds = append(followerIds, follower.UserId)
	}

	// 利用粉丝的id找到粉丝用户的信息followerUsers
	followerUsers, err := db.MQueryUsersByIds(s.ctx, followerIds)
	if err != nil {
		return nil, err
	}

	// 获取粉丝Set  loginId -> followerIds
	// 当前login用户时候关注UserId用户的粉丝
	followerSet, err := db.MQueryFollowByUserIdAndToUserIds(s.ctx, loginId, followerIds)
	if err != nil {
		return nil, err
	}

	userList := pack.BuildFollowList(followerUsers, followerSet)
	return userList, nil
}
