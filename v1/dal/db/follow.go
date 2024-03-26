package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"tiktok-backend/pkg/constants"
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

// CreateFollow 创建关注记录
// 增加当前用户的关注总数
// 增加其他用户的粉丝总数
// 这里需要写上事务
func CreateFollow(ctx context.Context, userId int64, toUserId int64) error {
	follow := &Follow{
		UserId:   userId,
		ToUserId: toUserId,
	}

	if err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 增加当前用户的关注总数
		if err := tx.Model(&User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			klog.Error("add user follow_count fail " + err.Error())
			return err
		}

		// 增加其他用户的粉丝总数
		if err := tx.Model(&User{}).Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			klog.Error("add user follower_count fail " + err.Error())
			return err
		}

		// 添加一条记录
		if err := tx.Create(follow).Error; err != nil {
			klog.Error("create relation record fail " + err.Error())
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

// DeleteFollow 删除关注记录
// 减少当前用户的关注总数
// 减少其他用户的粉丝总数
func DeleteFollow(ctx context.Context, userId int64, toUserId int64) error {
	if err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 增加当前用户的关注总数
		if err := tx.Model(&User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
			klog.Error("delete user follow_count fail " + err.Error())
			return err
		}

		// 增加其他用户的粉丝总数
		if err := tx.Model(&User{}).Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			klog.Error("delete user follower_count fail " + err.Error())
			return err
		}

		// 删除记录
		if err := tx.Where("user_id = ? AND to_user_id = ?", userId, toUserId).Delete(&Follow{}).Error; err != nil {
			klog.Error("delete relation record fail " + err.Error())
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

// MQueryFollowByUserIdAndToUserIds 根据当前用户id和目标用户id获取关注信息
func MQueryFollowByUserIdAndToUserIds(ctx context.Context, userId int64, toUserIds []int64) (map[int64]struct{}, error) {
	follows := make([]*Follow, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? AND to_user_id in (?)", userId, toUserIds).Find(&follows).Error; err != nil {
		klog.Error("query follow by userId and toUserId fail" + err.Error())
		return nil, err
	}
	followSet := make(map[int64]struct{})
	for _, follow := range follows {
		followSet[follow.ToUserId] = constants.BlankStruct{} // 表示当前用户userId关注了ToUserId
	}
	return followSet, nil
}

// QueryFollowByUserId 根据当前用户userId获取关注用户id
func QueryFollowByUserId(ctx context.Context, userId int64) ([]*Follow, error) {
	res := make([]*Follow, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&res).Error; err != nil {
		klog.Error("query follow by id fail " + err.Error())
		return nil, err
	}
	return res, nil
}


// QueryFollowerByToUserId 根据当前用户userId获取粉丝用户id
func QueryFollowerByToUserId(ctx context.Context, userId int64) ([]*Follow, error) {
	res := make([]*Follow, 0)
	if err := DB.WithContext(ctx).Where("to_user_id = ?", userId).Find(&res).Error; err != nil {
		klog.Error("query follower by id fail " + err.Error())
		return nil, err
	}
	return res, nil
}
