package service

import (
	"context"
	"errors"
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/relation"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/jwt"
)

type RelationActionService struct {
	ctx context.Context
}

// NewRelationActionService new RelationActionService
func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

// RelationAction 关注
// actionType=1 关注
// actionType=2 取消关注
func (s *RelationActionService) RelationAction(req *relation.DouyinRelationActionRequest) error {
	// 登录id
	claims, err := jwt.GetClaimsFromTokenStr(req.Token)
	if err != nil {
		return err
	}
	loginId := int64(claims[constants.IdentityKey].(float64))


	// 查找toUsers用户
	toUsers, err := db.MQueryUsersByIds(s.ctx, []int64{req.ToUserId})
	if err != nil {
		return err
	}
	if len(toUsers) == 0 {
		return errors.New("toUserId not exist")
	}

	// 关注
	if req.ActionType == constants.Follow {

		followSet, err := db.MQueryFollowByUserIdAndToUserIds(s.ctx, loginId, []int64{req.ToUserId})
		if err != nil {
			return err
		}

		// 之前没有关注过
		if _, ok := followSet[req.ToUserId]; !ok {
			err := db.CreateFollow(s.ctx, loginId, req.ToUserId)
			if err != nil {
				return err
			}
		}
		return nil
	}

	// 取消关注
	if req.ActionType == constants.UnFollow {

		// 确保这连个人之前关注过
		followSet, err := db.MQueryFollowByUserIdAndToUserIds(s.ctx, loginId, []int64{req.ToUserId})
		if err != nil {
			return err
		}
		if _, ok := followSet[req.ToUserId]; ok {
			err := db.DeleteFollow(s.ctx, loginId, req.ToUserId)
			if err != nil {
				return err
			}
		}

		return nil
	}

	return errors.New("actionType error")
}