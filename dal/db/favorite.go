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

// MQueryFavoriteByUserIdAndVideoIds 根据当前用户id和视频id获取点赞信息
func MQueryFavoriteByUserIdAndVideoIds(ctx context.Context, userId int64, videoIds []int64) (map[int64]struct{}, error) {
	res := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? AND video_id in (?)", userId, videoIds).Find(&res).Error; err != nil {
		klog.Error("query favorite record by userId and videoId fail " + err.Error())
		return nil, err
	}
	favoriteMap := make(map[int64]struct{})
	for _, favorite := range res {
		favoriteMap[favorite.VideoId] = constants.BlankStruct{} // 表示当前用户点赞了这个VideoId视频
	}
	return favoriteMap, nil
}

// CreateFavorite 添加一条favorite记录
// login用户的FavoriteCount++
// Video的FavoriteCount++
// Video用户的TotalFavorited++
func CreateFavorite(ctx context.Context, userId int64, videoId int64, videoUserId int64) error {
	favorite := &Favorite{UserId: userId, VideoId: videoId}

	if err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// login用户的FavoriteCount++
		if err := tx.Model(&User{}).Where("id = ?", userId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			klog.Error("add user favorite_count fail " + err.Error())
			return err
		}

		// Video的FavoriteCount++
		if err := tx.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			klog.Error("add video favorite_count fail " + err.Error())
			return err
		}

		// Video用户的TotalFavorited++
		if err := tx.Model(&User{}).Where("id = ?", videoUserId).Update("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error; err != nil {
			klog.Error("add user total_favorited fail " + err.Error())
			return err
		}

		// 添加一条记录
		if err := tx.Create(favorite).Error; err != nil {
			klog.Error("create favorite record fail " + err.Error())
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// DeleteFavorite 删除一条favorite记录
// login用户的FavoriteCount--
// Video的FavoriteCount--
// Video用户的TotalFavorited--
func DeleteFavorite(ctx context.Context, userId int64, videoId int64, videoUserId int64) error {

	if err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// login用户的FavoriteCount--
		if err := tx.Model(&User{}).Where("id = ?", userId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
			klog.Error("delete user favorite_count fail " + err.Error())
			return err
		}

		// Video的FavoriteCount--
		if err := tx.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
			klog.Error("delete video favorite_count fail " + err.Error())
			return err
		}

		// Video用户的TotalFavorited--
		if err := tx.Model(&User{}).Where("id = ?", videoUserId).Update("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error; err != nil {
			klog.Error("delete user total_favorited fail " + err.Error())
			return err
		}

		// 删除一条记录
		if err := tx.Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&Favorite{}).Error; err != nil {
			klog.Error("delete favorite record fail " + err.Error())
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}


// QueryFavoriteByUserId 根据当前用户userId获取关注用户id
func QueryFavoriteByUserId(ctx context.Context, userId int64) ([]*Favorite, error) {
	res := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&res).Error; err != nil {
		klog.Error("query favorite by id fail " + err.Error())
		return nil, err
	}
	return res, nil
}
