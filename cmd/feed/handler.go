package main

import (
	"context"
	"tiktok-backend/cmd/feed/pack"
	"tiktok-backend/cmd/feed/service"
	feed "tiktok-backend/kitex_gen/feed"
	"tiktok-backend/pkg/errno"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	resp = new(feed.DouyinFeedResponse)

	if req.LatestTime <= 0 {
		resp = pack.BuildFeedResp(errno.ParamErr)
		return resp, nil
	}

	videos, nextTime, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		resp = pack.BuildFeedResp(err)
		return resp, nil
	}

	resp = pack.BuildFeedResp(errno.Success)
	resp.VideoList = videos
	resp.NextTime = nextTime
	return resp, nil
}
