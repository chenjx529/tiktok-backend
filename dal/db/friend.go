package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
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


// QueryFriendByUserId 根据当前用户userId获取关注用户id
func QueryFriendByUserId(ctx context.Context, userId int64) ([]*Friend, error) {
	res := make([]*Friend, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&res).Error; err != nil {
		klog.Error("query friends by id fail " + err.Error())
		return nil, err
	}
	return res, nil
}
