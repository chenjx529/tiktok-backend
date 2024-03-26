package service

import (
	"context"
	"errors"
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/favorite"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/jwt"
)

type FavoriteActionService struct {
	ctx context.Context
}

// NewFavoriteActionService new FavoriteActionService
func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

// FavoriteAction 点赞
func (s *FavoriteActionService) FavoriteAction(req *favorite.DouyinFavoriteActionRequest) error {
	// 登录id
	claims, err := jwt.GetClaimsFromTokenStr(req.Token)
	if err != nil {
		return err
	}
	loginId := int64(claims[constants.IdentityKey].(float64))


	// 查找videos
	videos, err := db.MQueryVideoByVideoIds(s.ctx, []int64{req.VideoId})
	if err != nil {
		return err
	}
	if len(videos) == 0 {
		return errors.New("video not exist")
	}
	video := videos[0]

	// 点赞
	if req.ActionType == constants.Favorite {

		favoriteSet, err := db.MQueryFavoriteByUserIdAndVideoIds(s.ctx, loginId, []int64{req.VideoId})
		if err != nil {
			return err
		}

		// 之前没有点赞
		if _, ok := favoriteSet[req.VideoId]; !ok {
			err := db.CreateFavorite(s.ctx, loginId, req.VideoId, video.UserId)
			if err != nil {
				return err
			}
		}
		return nil
	}

	// 取消点赞
	if req.ActionType == constants.UnFavorite {

		// 确保之前点过赞
		favoriteSet, err := db.MQueryFavoriteByUserIdAndVideoIds(s.ctx, loginId, []int64{req.VideoId})
		if err != nil {
			return err
		}
		if _, ok := favoriteSet[req.VideoId]; ok {
			err := db.DeleteFavorite(s.ctx, loginId, req.VideoId, video.UserId)
			if err != nil {
				return err
			}
		}

		return nil
	}

	return errors.New("actionType error")
}