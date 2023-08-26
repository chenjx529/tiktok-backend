package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// Relation Gorm Data Structures
type Relation struct {
	gorm.Model
	UserId   int64 `gorm:"column:user_id;not null;index:idx_userid"`
	ToUserId int64 `gorm:"column:to_user_id;not null;index:idx_touserid"`
}

func (Relation) TableName() string {
	return "relation"
}

// MQueryRelationByIds 根据当前用户id和目标用户id获取关注信息
func MQueryRelationByIds(ctx context.Context, currentId int64, userIds []int64) (map[int64]*Relation, error) {
	var relations []*Relation
	if err := DB.WithContext(ctx).Where("user_id = ? AND to_user_id IN ?", currentId, userIds).Find(&relations).Error; err != nil {
		klog.Error("query relation by ids " + err.Error())
		return nil, err
	}
	relationMap := make(map[int64]*Relation)
	for _, relation := range relations {
		relationMap[relation.ToUserId] = relation
	}
	return relationMap, nil
}
