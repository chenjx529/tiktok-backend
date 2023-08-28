package service

import (
	"context"
	"tiktok-backend/cmd/relation/pack"
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/relation"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/jwt"
)

type RelationFriendListService struct {
	ctx context.Context
}

// NewRelationFriendListService new RelationFriendListService
func NewRelationFriendListService(ctx context.Context) *RelationFriendListService {
	return &RelationFriendListService{ctx: ctx}
}

func (s *RelationFriendListService) RelationFriendList(req *relation.DouyinRelationFriendListRequest) ([]*relation.FriendUser, error) {
	// 判断登录
	// 登录id
	claims, err := jwt.GetclaimsFromTokenStr(req.Token)
	if err != nil {
		return nil, err
	}
	loginId := int64(int(claims[constants.IdentityKey].(float64)))

	// 找到自己的好友
	friends, err := db.QueryFriendByUserId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	toUserIds := make([]int64, 0)
	for _, follow := range friends {
		toUserIds = append(toUserIds, follow.ToUserId) // 粉丝的id
	}

	// 利用toUserIds获取toUser信息
	toUsers, err := db.MQueryUsersByIds(s.ctx, toUserIds)
	if err != nil {
		return nil, err
	}

	// 获取关注set
	followSet, err := db.MQueryFollowByUserIdAndToUserIds(s.ctx, loginId, toUserIds)
	if err != nil {
		return nil, err
	}

	// 通过好友的MessageId找到具体message
	// 这里需要一个meesageMap  map[friend_id]*meesage
	messageMap := make(map[int64]*db.Message)
	for _, friend := range friends {
		if friend.MessageId != 0 {
			message, err := db.QueryMessageById(s.ctx, friend.MessageId)
			if err != nil {
				return nil, err
			}
			messageMap[friend.ToUserId] = message[0]
		}
	}

	// 封装relation.FriendUser
	userList := pack.BuildFriendUserList(toUsers, followSet, messageMap)
	return userList, nil
}
