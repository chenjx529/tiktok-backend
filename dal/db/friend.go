package db

import (
	"gorm.io/gorm"
)

// Friend Gorm Data Structures
type Friend struct {
	gorm.Model
	UserId    int64 `gorm:"column:user_id;not null;index:idx_userid"`       // 当前用户id
	ToUserId  int64 `gorm:"column:to_user_id;not null;index:idx_touserid"`  // 好友id
	MessageId int64 `gorm:"column:to_user_id;not null;index:idx_messageid"` // 消息
}

func (Friend) TableName() string {
	return "friend"
}
