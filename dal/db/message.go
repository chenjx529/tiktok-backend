package db

import (
	"gorm.io/gorm"
)

// Message Gorm Data Structures
type Message struct {
	gorm.Model
	ToUserId   int64  `gorm:"column:to_user_id;not null;index:idx_touserid"`     // 该消息接收者的id
	FromUserId int64  `gorm:"column:from_user_id;not null;index:idx_fromuserid"` // 该消息发送者的id
	Content    string `grom:"column:contents;type:varchar(255);not null"`        // 消息内容
}

func (Message) TableName() string {
	return "message"
}
