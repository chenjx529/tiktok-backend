package main

import (
	"context"
	"tiktok-backend/cmd/message/pack"
	"tiktok-backend/cmd/message/service"
	message "tiktok-backend/kitex_gen/message"
	"tiktok-backend/pkg/errno"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.DouyinMessageActionRequest) (resp *message.DouyinMessageActionResponse, err error) {
	resp = new(message.DouyinMessageActionResponse)

	if len(req.Token) == 0 || req.ToUserId == 0 || req.ActionType != 1 || len(req.Content) == 0 {
		resp = pack.BuildMessageActionResp(errno.ParamErr)
		return resp, nil
	}

	if err := service.NewMessageActionService(ctx).MessageAction(req); err != nil {
		resp = pack.BuildMessageActionResp(err)
		return resp, nil
	}

	resp = pack.BuildMessageActionResp(errno.Success)
	return resp, nil
}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) (resp *message.DouyinMessageChatResponse, err error) {
	resp = new(message.DouyinMessageChatResponse)

	if len(req.Token) == 0 || req.ToUserId == 0 {
		resp = pack.BuildMessageChatResp(errno.ParamErr)
		return resp, nil
	}

	messageList, err := service.NewMessageChatService(ctx).MessageChat(req)
	if err != nil {
		resp = pack.BuildMessageChatResp(err)
		return resp, nil
	}

	resp = pack.BuildMessageChatResp(errno.Success)
	resp.MessageList = messageList
	return resp, nil
}
