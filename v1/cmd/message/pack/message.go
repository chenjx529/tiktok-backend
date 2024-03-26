package pack

import (
	"strconv"
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/message"
)

func BuildMessageList(dbMessageList []*db.Message) []*message.Message {
	messageList := make([]*message.Message, 0)
	for _, mes := range dbMessageList {
		messageList = append(messageList, buildMessageInfo(mes))
	}
	return messageList
}

func buildMessageInfo(mes *db.Message) *message.Message {
	return &message.Message{
		Id:         int64(mes.ID),                         // 消息id
		FromUserId: mes.FromUserId,                        // 该消息发送者的id
		ToUserId:   mes.ToUserId,                          // 该消息接收者的id
		Content:    mes.Content,                           // 消息内容
		CreateTime: strconv.FormatInt(mes.CreateTime, 10), // 消息创建时间
	}
}
