package service

import (
	"context"
	"sync"
	"tiktok-backend/cmd/publish/pack"
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/publish"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/jwt"
)

type PublishListService struct {
	ctx context.Context
}

// NewPublishListService new PublishListService
func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{
		ctx: ctx,
	}
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishListService) PublishList(req *publish.DouyinPublishListRequest) ([]*publish.Video, error) {
	// 是否登录
	claims, err := jwt.GetclaimsFromTokenStr(req.Token)
	var login_id int64
	if err != nil {
		login_id = 0
	} else {
		login_id = int64(int(claims[constants.IdentityKey].(float64)))
	}

	// 利用用户名查找视频
	videoData, err := db.QueryVideoByUserId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	// 获取视频id和用户id
	videoIds := make([]int64, 0)
	userIds := []int64{req.UserId}
	for _, video := range videoData {
		videoIds = append(videoIds, int64(video.ID))
	}

	// 获取用户信息
	users, err := db.MQueryUsersByIds(s.ctx, userIds)
	if err != nil {
		return nil, err
	}
	userMap := make(map[int64]*db.User)
	for _, user := range users {
		userMap[int64(user.ID)] = user
	}

	// 视频点赞和用户关注
	var favoriteMap map[int64]*db.Favorite
	var relationMap map[int64]*db.Relation
	if login_id == 0 {
		favoriteMap = nil
		relationMap = nil
	} else {
		var wg sync.WaitGroup
		wg.Add(2)
		var favoriteErr, relationErr error

		//获取点赞信息
		go func() {
			defer wg.Done()
			favoriteMap, err = db.MQueryFavoriteByIds(s.ctx, login_id, videoIds)
			if err != nil {
				favoriteErr = err
				return
			}
		}()

		//获取关注信息
		go func() {
			defer wg.Done()
			relationMap, err = db.MQueryRelationByIds(s.ctx, login_id, userIds)
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
	videoListInfo := pack.VideoListInfo(login_id, videoData, userMap, favoriteMap, relationMap)
	return videoListInfo, nil
}