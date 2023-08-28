package service

import (
	"context"
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
	claims, err := jwt.GetclaimsFromTokenStr(req.Token)
	if err != nil {
		return err
	}
	loginId := int64(int(claims[constants.IdentityKey].(float64)))

	// 新建一条消息记录
	mesId, err := db.CreateMessage(s.ctx, &db.Message{
		FromUserId: loginId,
		ToUserId:   req.ToUserId,
		Content:    req.Content,
	})
	if err != nil {
		return err
	}

	// 更新朋友关系中的messageId
	if err := db.UpdateFriendForMessageIdById(s.ctx, loginId, req.ToUserId, mesId); err != nil {
		return err
	}

	return nil
}
