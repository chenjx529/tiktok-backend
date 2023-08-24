package main

import (
	"context"
	message "tiktok-backend/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) (resp *message.DouyinMessageChatResponse, err error) {
	// TODO: Your code here...
	return
}

// MessageActioin implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageActioin(ctx context.Context, req *message.DouyinRelationActionRequest) (resp *message.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}
