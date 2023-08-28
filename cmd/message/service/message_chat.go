package service

import (
	"context"
	"tiktok-backend/cmd/message/pack"
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/message"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/jwt"
)

type MessageChatService struct {
	ctx context.Context
}

// NewMessageChatService new MessageChatService
func NewMessageChatService(ctx context.Context) *MessageChatService {
	return &MessageChatService{ctx: ctx}
}

func (s *MessageChatService) MessageChat(req *message.DouyinMessageChatRequest) ([]*message.Message, error) {
	// 登录id
	claims, err := jwt.GetClaimsFromTokenStr(req.Token)
	if err != nil {
		return nil, err
	}
	loginId := int64(claims[constants.IdentityKey].(float64))

	// 查询聊天记录
	messageList, err := db.QueryMessageByUserId(s.ctx, loginId, req.ToUserId)
	if err != nil {
		return nil, err
	}

	return pack.BuildMessageList(messageList), nil
}
