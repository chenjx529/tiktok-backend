package service

import (
	"context"
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/favorite"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/jwt"
)

type FavoriteListService struct {
	ctx context.Context
}

// NewFavoriteListService new RelationFollowListService
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

func (s *FavoriteListService) FavoriteList(req *favorite.DouyinFavoriteListRequest) ([]*favorite.Video, error) {
	// 登录id
	claims, err := jwt.GetClaimsFromTokenStr(req.Token)
	if err != nil {
		return nil, err
	}
	loginId := int64(claims[constants.IdentityKey].(float64))

	// 获取userId的favorite
	favoriteForUserId, err := db.QueryFavoriteByUserId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	favoriteVideoIds := make([]int64, 0)
	for _, v := range favoriteForUserId {
		favoriteVideoIds = append(favoriteVideoIds, v.VideoId)  // 点赞视频的id
	}

	// 获取点赞视频的信息
	videos, err := db.MQueryVideoByVideoIds(s.ctx, favoriteVideoIds)
	if err != nil {
		return err
	}
}
