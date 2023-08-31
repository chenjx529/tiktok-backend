package db

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	UserId        int64  `gorm:"column:user_id;not null;index:idx_userid"` // 视频作者信息,绑定用户id，索引
	Title         string `gorm:"column:title;type:varchar(128);not null"`  // 视频标题
	PlayUrl       string `gorm:"column:play_url;varchar(128);not null"`    // 视频播放地址
	CoverUrl      string `gorm:"column:cover_url;varchar(128);not null"`   // 视频封面地址
	FavoriteCount int64  `gorm:"column:favorite_count;default:0"`          // 视频的点赞总数
	CommentCount  int64  `gorm:"column:comment_count;default:0"`           // 视频的评论总数
}

func (v *Video) TableName() string {
	return "video"
}

// QueryVideoByLatestTime 通过LatestTime获取视频，倒序前30个
func QueryVideoByLatestTime(ctx context.Context, latestTime int64) ([]*Video, error) {
	var videos []*Video
	updatedTime := time.UnixMilli(latestTime)
	if err := DB.WithContext(ctx).Limit(30).Order("updated_at desc").Where(`updated_at < ?`, updatedTime).Find(&videos).Error; err != nil {
		klog.Error("QueryVideoByLatestTime find video error " + err.Error())
		return videos, err
	}
	return videos, nil
}

// MQueryVideoByVideoIds 通过视频id获取视频
func MQueryVideoByVideoIds(ctx context.Context, videoIds []int64) ([]*Video, error) {
	var videos []*Video
	if err := DB.WithContext(ctx).Where("id in (?)", videoIds).Find(&videos).Error; err != nil {
		klog.Error("QueryVideoByVideoIds error " + err.Error())
		return nil, err
	}
	return videos, nil
}

// QueryVideoByUserId 通过用户id获取视频
func QueryVideoByUserId(ctx context.Context, userId int64) ([]*Video, error) {
	var videos []*Video
	if err := DB.WithContext(ctx).Order("updated_at desc").Where("user_id = ?", userId).Find(&videos).Error; err != nil {
		klog.Error("QueryVideoByUserId find video error " + err.Error())
		return nil, err
	}
	return videos, nil
}

// CreateVideo creates a new video
// login用户的work_count++
func CreateVideo(ctx context.Context, video *Video, loginId int64) error {
	if err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// login用户的work_count++
		if err := tx.Model(&User{}).Where("id = ?", loginId).Update("work_count", gorm.Expr("work_count + ?", 1)).Error; err != nil {
			klog.Error("add user favorite_count fail " + err.Error())
			return err
		}
		// 添加一个视频记录
		if err := DB.WithContext(ctx).Create(video).Error; err != nil {
			klog.Error("create video fail " + err.Error())
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
