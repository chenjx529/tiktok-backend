package handlers

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"tiktok-backend/cmd/api/pack"
	"tiktok-backend/cmd/api/rpc"
	"tiktok-backend/kitex_gen/favorite"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/errno"
)

// FavoriteAction 登录用户对视频的点赞和取消点赞操作
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	tokenStr := c.Query("token")
	videoIdStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")

	if len(tokenStr) == 0 {
		pack.SendFavoriteActionResponse(c, errno.ParamErr)
		return
	}

	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		pack.SendFavoriteActionResponse(c, err)
		return
	}

	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		pack.SendFavoriteActionResponse(c, err)
		return
	}
	if actionType != constants.Favorite && actionType != constants.UnFavorite {
		pack.SendFavoriteActionResponse(c, errors.New("actionType error"))
		return
	}

	if err = rpc.FavoriteAction(context.Background(), &favorite.DouyinFavoriteActionRequest{
		Token:      tokenStr,
		VideoId:    videoId,
		ActionType: int32(actionType),
	}); err != nil {
		pack.SendFavoriteActionResponse(c, err)
		return
	}

	pack.SendFavoriteActionResponse(c, errno.Success)
}

// FavoriteList 用户的所有点赞视频
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	userIdStr := c.Query("user_id")

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		pack.SendFavoriteListResponse(c, err, nil)
		return
	}

	userList, err := rpc.FavoriteList(context.Background(), &favorite.DouyinFavoriteListRequest{
		Token:  token,
		UserId: userId,
	})
	if err != nil {
		pack.SendFavoriteListResponse(c, err, nil)
		return
	}

	pack.SendFavoriteListResponse(c, errno.Success, userList)
}
