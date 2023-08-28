package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
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

// QueryMessageByUserId 根据当前用户userId获取关注用户id
func QueryMessageByUserId(ctx context.Context, fromUserId int64, toUserId int64) ([]*Message, error) {
	res := make([]*Message, 0)
	querySql := "(to_user_id = ? and from_user_id = ?) or (to_user_id = ? and from_user_id = ?)"
	if err := DB.WithContext(ctx).Order("created_at desc").Where(querySql, toUserId, fromUserId, fromUserId, toUserId).Find(&res).Error; err != nil {
		klog.Error("query Message by userId fail " + err.Error())
		return nil, err
	}
	return res, nil
}

// QueryMessageById 根据id查找
func QueryMessageById(ctx context.Context, userId int64) ([]*Message, error) {
	res := make([]*Message, 0)
	if err := DB.WithContext(ctx).Where("id = ?", userId).Find(&res).Error; err != nil {
		klog.Error("query Message by id fail " + err.Error())
		return nil, err
	}
	return res, nil
}

// CreateMessage 新建一条聊天记录
func CreateMessage(ctx context.Context, mes *Message) (int64, error) {
	if err := DB.WithContext(ctx).Create(mes).Error; err != nil {
		klog.Error("create user fail " + err.Error())
		return 0, err
	}
	return int64(mes.ID), nil
}
