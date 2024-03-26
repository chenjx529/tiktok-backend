package pack

import (
	"errors"
	"tiktok-backend/kitex_gen/message"
	"tiktok-backend/pkg/errno"
)

// BuildMessageActionResp 发送MessageActionResponse
func BuildMessageActionResp(err error) *message.DouyinMessageActionResponse {
	if err == nil {
		return messageActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return messageActionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return messageActionResp(s)
}

func messageActionResp(err errno.ErrNo) *message.DouyinMessageActionResponse {
	return &message.DouyinMessageActionResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

// BuildMessageChatResp 发送MessageChatResponse
func BuildMessageChatResp(err error) *message.DouyinMessageChatResponse {
	if err == nil {
		return messageChatResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return messageChatResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return messageChatResp(s)
}

func messageChatResp(err errno.ErrNo) *message.DouyinMessageChatResponse {
	return &message.DouyinMessageChatResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
