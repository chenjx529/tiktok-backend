package service

import (
	"context"
	"sync"
	"tiktok-backend/cmd/favorite/pack"
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
	favorites, err := db.QueryFavoriteByUserId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	favoriteVideoIds := make([]int64, 0)
	for _, v := range favorites {
		favoriteVideoIds = append(favoriteVideoIds, v.VideoId)  // 点赞视频的id
	}

	// 获取点赞视频的信息
	favoriteVideos, err := db.MQueryVideoByVideoIds(s.ctx, favoriteVideoIds)
	if err != nil {
		return nil, err
	}

	// 获取点赞视频的用户id
	videoUserIds := make([]int64, 0)
	for _, video := range favoriteVideos {
		videoUserIds = append(videoUserIds, video.UserId)
	}

	// 获取点赞视频的用户信息
	videoUsers, err := db.MQueryUsersByIds(s.ctx, videoUserIds)
	if err != nil {
		return nil, err
	}
	userMap := make(map[int64]*db.User)
	for _, user := range videoUsers {
		userMap[int64(user.ID)] = user
	}

	// 视频点赞和用户关注
	var favoriteSet map[int64]struct{}
	var followSet map[int64]struct{}
	if loginId != 0 {
		var wg sync.WaitGroup
		wg.Add(2)
		var favoriteErr, relationErr error

		//获取点赞信息
		go func() {
			defer wg.Done()
			favoriteSet, err = db.MQueryFavoriteByUserIdAndVideoIds(s.ctx, loginId, favoriteVideoIds)
			if err != nil {
				favoriteErr = err
				return
			}
		}()

		//获取关注信息
		go func() {
			defer wg.Done()
			followSet, err = db.MQueryFollowByUserIdAndToUserIds(s.ctx, loginId, videoUserIds)
			if err != nil {
				relationErr = err
				return
			}
		}()

		wg.Wait()
		if favoriteErr != nil {
			return nil, favoriteErr
		}
		if relationErr != nil {
			return nil, relationErr
		}
	}

	// 封装db数据到response
	videoList := pack.BuildVideoList(loginId, favoriteVideos, userMap, favoriteSet, followSet)
	return videoList, nil
}
