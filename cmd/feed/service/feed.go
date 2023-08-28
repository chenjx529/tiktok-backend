package service

import (
	"context"
	"sync"
	"tiktok-backend/cmd/feed/pack"
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/feed"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/jwt"
)

type FeedService struct {
	ctx context.Context
}

// NewFeedService new FeedService
func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

// Feed 不限制登录状态返回视频流，
// 难点1：按投稿时间倒序的视频列表，单次最多30个。
// 难点2：视频是否点赞  is_favorite
// 难点3：当前用户时候关注  is_follow
// 这个函数我们使用了多线程，同时获取is_favorite和is_follow
func (s *FeedService) Feed(req *feed.DouyinFeedRequest) ([]*feed.Video, int64, error) {
	// 是否登录
	claims, err := jwt.GetClaimsFromTokenStr(req.Token)
	var login_id int64
	if err == nil  {
		login_id = int64(claims[constants.IdentityKey].(float64))  // 这种写法，我是真的想骂人的
	}

	// 按投稿时间倒序的视频列表，单次最多30个。
	videoData, err := db.QueryVideoByLatestTime(s.ctx, req.LatestTime)
	if err != nil {
		return nil, 0, err
	}

	// 获取视频id和用户id
	videoIds := make([]int64, 0)
	userIds := make([]int64, 0)
	for _, video := range videoData {
		videoIds = append(videoIds, int64(video.ID))
		userIds = append(userIds, video.UserId)
	}

	// 获取用户信息
	users, err := db.MQueryUsersByIds(s.ctx, userIds)
	if err != nil {
		return nil, 0, err
	}
	userMap := make(map[int64]*db.User)
	for _, user := range users {
		userMap[int64(user.ID)] = user
	}

	// 视频点赞和用户关注
	var favoriteSet map[int64]struct{}
	var followSet map[int64]struct{}
	if login_id == 0 {
		favoriteSet = nil
		followSet = nil
	} else {
		var wg sync.WaitGroup
		wg.Add(2)
		var favoriteErr, relationErr error

		//获取点赞信息
		go func() {
			defer wg.Done()
			favoriteSet, err = db.MQueryFavoriteByIds(s.ctx, login_id, videoIds)
			if err != nil {
				favoriteErr = err
				return
			}
		}()

		//获取关注信息
		go func() {
			defer wg.Done()
			followSet, err = db.MQueryFollowByUserIdAndToUserIds(s.ctx, login_id, userIds)
			if err != nil {
				relationErr = err
				return
			}
		}()

		wg.Wait()
		if favoriteErr != nil {
			return nil, 0, favoriteErr
		}
		if relationErr != nil {
			return nil, 0, relationErr
		}
	}

	videoListInfo, nextTime := pack.VideoListInfo(login_id, videoData, userMap, favoriteSet, followSet)
	return videoListInfo, nextTime, nil
}
