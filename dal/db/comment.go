package db

import (
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
