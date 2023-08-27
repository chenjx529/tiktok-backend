package service

import (
	"context"
	"tiktok-backend/cmd/relation/pack"
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/relation"
	"tiktok-backend/pkg/constants"
)

type RelationFollowListService struct {
	ctx context.Context
}

// NewRelationFollowListService new RelationFollowListService
func NewRelationFollowListService(ctx context.Context) *RelationFollowListService {
	return &RelationFollowListService{ctx: ctx}
}

// RelationFollowList 获取关注列表
func (s *RelationFollowListService) RelationFollowList(req *relation.DouyinRelationFollowListRequest) ([]*relation.User, error) {
	// 登录id  所以这个token是没用的吧
	//claims, err := jwt.GetclaimsFromTokenStr(req.Token)
	//if err != nil {
	//	return nil, err
	//}
	//loginId := int64(int(claims[constants.IdentityKey].(float64)))

	// 获取toUserIds
	follows, err := db.QueryFollowByUserId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	toUserIds := make([]int64, 0)
	for _, follow := range follows {
		toUserIds = append(toUserIds, follow.ToUserId)  // 关注用户的id
	}


	// 利用toUserIds获取toUser信息
	toUsers, err := db.MQueryUsersByIds(s.ctx, toUserIds)
	if err != nil {
		return nil, err
	}

	// 获取关注Set
	followSet := make(map[int64]struct{})
	for _, follow := range follows {
		followSet[follow.ToUserId] = constants.BlankStruct{} // 表示当前用户userId关注了ToUserId
	}

	userList := pack.BuildFollowList(toUsers, followSet)
	return userList, nil
}
