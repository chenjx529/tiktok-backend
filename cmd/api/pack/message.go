package pack

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"tiktok-backend/kitex_gen/message"
	"tiktok-backend/pkg/errno"
)

func SendMessageActionResponse(c *app.RequestContext, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, MessageActionResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
	})
}

func SendMessageChatResponse(c *app.RequestContext, err error, messageList []*message.Message) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, MessageChatResponse{
		StatusCode:  Err.ErrCode,
		StatusMsg:   Err.ErrMsg,
		MessageList: buildMessageList(messageList),
	})
}

func buildMessageList(kitexMessageList []*message.Message) []*Message {
	messageList := make([]*Message, 0)
	for _, mes := range kitexMessageList {
		messageList = append(messageList, buildMessageInfo(mes))
	}
	return messageList
}

func buildMessageInfo(mes *message.Message) *Message {
	return &Message{
		Id:         mes.Id,         // 消息id
		FromUserId: mes.FromUserId, // 该消息发送者的id
		ToUserId:   mes.ToUserId,   // 该消息接收者的id
		Content:    mes.Content,    // 消息内容
		CreateTime: mes.CreateTime, // 消息创建时间
	}
}

