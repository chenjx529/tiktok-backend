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

// CreateComment 添加一条comment记录
// video的comment_count++
func CreateComment(ctx context.Context, com *Comment) error {
	if err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		// Video的comment_count++
		if err := tx.Model(&Video{}).Where("id = ?", com.VideoId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			klog.Error("add video comment_count fail " + err.Error())
			return err
		}

		// 添加一条记录
		if err := tx.Create(com).Error; err != nil {
			klog.Error("create comment record fail " + err.Error())
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// DeleteComment 删除一条comment记录
// video的comment_count--
func DeleteComment(ctx context.Context, commentId int64, videoId int64) (*Comment, error) {
	var com *Comment
	if err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 获取comment
		if err := tx.Where("id = ?", commentId).First(&com).Error; err != nil{
			klog.Error("find comment fail ")
			return err
		}
		// Video的comment_count++
		if err := tx.Model(&Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			klog.Error("delete video comment_count fail " + err.Error())
			return err
		}
		// 删除一条记录
		if err := tx.Where("id = ?", commentId).Delete(&Comment{}).Error; err != nil {
			klog.Error("delete comment record fail " + err.Error())
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return com, nil
}

// QueryCommentByVideoId 获取评论列表
func QueryCommentByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := DB.WithContext(ctx).Order("updated_at desc").Where("video_id = ?", videoId).Find(&res).Error; err != nil {
		klog.Error("CommentList find comment failed" + err.Error())
		return nil, err
	}
	return res, nil
}
