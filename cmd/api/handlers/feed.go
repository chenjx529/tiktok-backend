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

// Feed 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
// 难点1：按投稿时间倒序的视频列表，单次最多30个。
// 难点2：视频是否点赞  is_favorite
// 难点3：当前用户时候关注  is_follow
func Feed(ctx context.Context, c *app.RequestContext) {
	token := c.DefaultQuery("token", "")
	latestTimeStr := c.DefaultQuery("latest_time", strconv.Itoa(int(time.Now().UnixMilli())))
	latestTime, err := strconv.ParseInt(latestTimeStr, 10, 64)
	if err != nil {
		pack.SendFeedResponse(c, err, nil, 0)
		return
	}

	req := &feed.DouyinFeedRequest{LatestTime: latestTime, Token: token}
	video, nextTime, err := rpc.Feed(context.Background(), req)
	if err != nil {
		pack.SendFeedResponse(c, err,nil, 0)
		return
	}

	pack.SendFeedResponse(c, errno.Success, video, nextTime)
}