package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"tiktok-backend/pkg/constants"
)

// Favorite Gorm Data Structures
type Favorite struct {
	gorm.Model
	UserId  int64 `gorm:"column:user_id;not null;index:idx_userid"`   // 用户id
	VideoId int64 `gorm:"column:video_id;not null;index:idx_videoid"` // 视频id
}

func (Favorite) TableName() string {
	return "favorite"
}

// MQueryFavoriteByIds 根据当前用户id和视频id获取点赞信息
func MQueryFavoriteByIds(ctx context.Context, currentId int64, videoIds []int64) (map[int64]struct{}, error) {
	res := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? AND video_id IN ?", currentId, videoIds).Find(&res).Error; err != nil {
		klog.Error("quert favorite record fail " + err.Error())
		return nil, err
	}
	favoriteMap := make(map[int64]struct{})
	for _, favorite := range res {
		favoriteMap[favorite.VideoId] = constants.BlankStruct{}  // 表示当前用户点赞了这个VideoId视频
	}
	return favoriteMap, nil
}
