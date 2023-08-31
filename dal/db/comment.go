package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// Comment Gorm Data Structures
type Comment struct {
	gorm.Model
	UserId   int64  `gorm:"column:user_id;not null;index:idx_userid"`   // 用户id
	VideoId  int64  `gorm:"column:video_id;not null;index:idx_videoid"` // 视频id
	Contents string `grom:"column:contents;type:varchar(255);not null"` // 内容
}

func (Comment) TableName() string {
	return "comment"
}

// CommentAction 评论操作
func CommentAction(ctx context.Context, comment *Comment) (int64, error) {
	if err := DB.WithContext(ctx).Create(comment).Error; err != nil {
		klog.Error("create comment fail" + err.Error())
		return 0, nil
	}
	return int64(comment.ID), nil
}

// CommentList 获取评论列表
func CommentList(ctx context.Context, videoId int64) ([]*Comment, error) {
	var commentList []*Comment
	if err := DB.WithContext(ctx).Order("updated_at desc").Where("video_id = ?", videoId).Find(&commentList).Error; err != nil {
		klog.Error("CommentList find comment failed" + err.Error())
		return nil, err
	}
	return commentList, nil
}
