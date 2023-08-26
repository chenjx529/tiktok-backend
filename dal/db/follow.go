package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// Follow Gorm Data Structures
type Follow struct {
	gorm.Model
	UserId   int64 `gorm:"column:user_id;not null;index:idx_userid"`      // 当前用户
	ToUserId int64 `gorm:"column:to_user_id;not null;index:idx_touserid"` // 关注用户
}

func (Follow) TableName() string {
	return "follow"
}

// MQueryFollowByIds 根据当前用户id和目标用户id获取关注信息
func MQueryFollowByIds(ctx context.Context, currentId int64, userIds []int64) (map[int64]*Follow, error) {
	var follows []*Follow
	if err := DB.WithContext(ctx).Where("user_id = ? AND to_user_id IN ?", currentId, userIds).Find(&follows).Error; err != nil {
		klog.Error("query follow by ids " + err.Error())
		return nil, err
	}
	followMap := make(map[int64]*Follow)
	for _, relation := range follows {
		followMap[relation.ToUserId] = relation
	}
	return followMap, nil
}
