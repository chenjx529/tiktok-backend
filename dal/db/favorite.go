package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// Favorite Gorm Data Structures
type Favorite struct {
	gorm.Model
	UserId  int64 `gorm:"column:user_id;not null;index:idx_userid"`
	VideoId int64 `gorm:"column:video_id;not null;index:idx_videoid"`
}

func (Favorite) TableName() string {
	return "favorite"
}

// MQueryFavoriteByIds 根据当前用户id和视频id获取点赞信息
func MQueryFavoriteByIds(ctx context.Context, currentId int64, videoIds []int64) (map[int64]*Favorite, error) {
	var favorites []*Favorite
	if err := DB.WithContext(ctx).Where("user_id = ? AND video_id IN ?", currentId, videoIds).Find(&favorites).Error; err != nil {
		klog.Error("quert favorite record fail " + err.Error())
		return nil, err
	}
	favoriteMap := make(map[int64]*Favorite)
	for _, favorite := range favorites {
		favoriteMap[favorite.VideoId] = favorite
	}
	return favoriteMap, nil
}
