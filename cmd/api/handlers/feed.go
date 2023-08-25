package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"tiktok-backend/cmd/api/pack"
	"tiktok-backend/cmd/api/rpc"
	"tiktok-backend/kitex_gen/feed"
	"tiktok-backend/pkg/errno"
	"time"
)

func Feed(ctx context.Context, c *app.RequestContext) {
	token := c.DefaultQuery("token", "")
	defaultTimeStr := strconv.Itoa(int(time.Now().UnixMilli()))
	latestTimeStr := c.DefaultQuery("latest_time", defaultTimeStr)
	latestTime, err := strconv.ParseInt(latestTimeStr, 10, 64)
	if err != nil {
		pack.SendFeedResponse(c, err, nil, 0)
		return
	}

	// 我就写一个注释
	req := &feed.DouyinFeedRequest{LatestTime: latestTime, Token: token}
	video, nextTime, err := rpc.Feed(context.Background(), req)
	if err != nil {
		pack.SendFeedResponse(c, err,nil, 0)
		return
	}

	pack.SendFeedResponse(c, errno.Success, video, nextTime)
}