package service

import (
	"context"
	"errors"
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/message"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/jwt"
)

type MessageActionService struct {
	ctx context.Context
}

// NewMessageActionService new MessageActionService
func NewMessageActionService(ctx context.Context) *MessageActionService {
	return &MessageActionService{ctx: ctx}
}

func (s *MessageActionService) MessageAction(req *message.DouyinMessageActionRequest) error {
	// 登录id
	claims, err := jwt.GetClaimsFromTokenStr(req.Token)
	if err != nil {
		return err
	}
	loginId := int64(claims[constants.IdentityKey].(float64))

	// 先检查有没有这个朋友
	ok, err := db.QueryFriendShipByUserId(s.ctx, loginId, req.ToUserId)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("no this friend")
	}

	// 新建一条消息记录
	if err := db.CreateMessage(s.ctx, loginId, req.ToUserId, req.Content); err != nil {
		return err
	}

	return nil
}
