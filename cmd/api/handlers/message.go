package handlers

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"tiktok-backend/cmd/api/pack"
	"tiktok-backend/cmd/api/rpc"
	"tiktok-backend/kitex_gen/message"
	"tiktok-backend/pkg/errno"
	"time"
)

// MessageAction 登录用户对消息的相关操作，目前只支持消息发送
func MessageAction(ctx context.Context, c *app.RequestContext) {
	tokenStr := c.Query("token")
	toUserIdStr := c.Query("to_user_id")
	actionTypeStr := c.Query("action_type") // 1-发送消息
	contentStr := c.Query("content")

	if len(tokenStr) == 0 {
		pack.SendMessageActionResponse(c, errno.ParamErr)
		return
	}

	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		pack.SendMessageActionResponse(c, err)
		return
	}

	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		pack.SendMessageActionResponse(c, err)
		return
	}
	if actionType != 1 {
		pack.SendMessageActionResponse(c, errors.New("actionType error"))
		return
	}

	if err = rpc.MessageAction(context.Background(), &message.DouyinMessageActionRequest{
		Token:      tokenStr,
		ToUserId:   toUserId,
		ActionType: int32(actionType),
		Content:    contentStr,
	}); err != nil {
		pack.SendMessageActionResponse(c, err)
		return
	}

	pack.SendMessageActionResponse(c, errno.Success)
}

// MessageChat 当前登录用户和其他指定用户的聊天消息记录
func MessageChat(ctx context.Context, c *app.RequestContext) {
	tokenStr := c.Query("token")
	toUserIdStr := c.Query("to_user_id")
	preMsgTimeStr := c.DefaultQuery("pre_msg_time", "0")
	preMsgTime, err := strconv.ParseInt(preMsgTimeStr, 10, 64)
	if err != nil {
		pack.SendMessageChatResponse(c, err, nil)
		return
	}
	if preMsgTime == 0 {
		preMsgTime = time.Now().UnixMilli()
	}

	if len(tokenStr) == 0 {
		pack.SendMessageChatResponse(c, errno.ParamErr, nil)
		return
	}

	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		pack.SendMessageChatResponse(c, err, nil)
		return
	}

	messageList, err := rpc.MessageChat(context.Background(), &message.DouyinMessageChatRequest{
		Token:      tokenStr,
		ToUserId:   toUserId,
		PreMsgTime: preMsgTime,
	})
	if err != nil {
		pack.SendMessageChatResponse(c, err, nil)
		return
	}

	pack.SendMessageChatResponse(c, errno.Success, messageList)
}
