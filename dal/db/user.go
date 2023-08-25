package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string `gorm:"column:name;index:idx_name,unique;type:varchar(32);not null"` // 用户名称 设置为索引，unique
	Password        string `gorm:"column:password;type:varchar(32);not null"`                   // 密码
	FollowCount     int64  `gorm:"column:follow_count;default:0"`                               // 关注总数
	FollowerCount   int64  `gorm:"column:follower_count;default:0"`                             // 粉丝总数
	Avatar          string `gorm:"column:avatar;type:varchar(32)"`                              // 用户头像
	BackgroundImage string `gorm:"column:background_image;type:varchar(32)"`                    // 用户个人页顶部大图
	Signature       string `gorm:"column:signature;type:varchar(32)"`                           // 个人简介
	TotalFavorited  int64  `gorm:"column:total_favorited;default:0"`                            // 获赞数量
	WorkCount       int64  `gorm:"column:work_count;default:0"`                                 // 作品数量
	FavoriteCount   int64  `gorm:"column:favorite_count;default:0"`                             // 点赞数量
}

func (User) TableName() string {
	return "user"
}

// QueryUserByIds 根据用户id获取用户信息
func QueryUserByIds(ctx context.Context, userIds []int64) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("id in (?)", userIds).Find(&res).Error; err != nil {
		klog.Error("query user by ids fail " + err.Error())
		return nil, err
	}
	return res, nil
}

// QueryUserByName 根据用户名获取用户信息
func QueryUserByName(ctx context.Context, name string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("name = ?", name).Find(&res).Error; err != nil {
		klog.Error("query user by name fail " + err.Error())
		return nil, err
	}
	return res, nil
}

// CreateUser 上传用户信息到数据库
func CreateUser(ctx context.Context, user *User) (int64, error) {
	if err := DB.WithContext(ctx).Create(user).Error; err != nil {
		klog.Error("create user fail " + err.Error())
		return 0, err
	}
	return int64(user.ID), nil
}
