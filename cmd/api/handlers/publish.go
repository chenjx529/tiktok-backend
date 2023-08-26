package handlers

import (
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"io"
	"strconv"
	"tiktok-backend/cmd/api/pack"
	"tiktok-backend/cmd/api/rpc"
	"tiktok-backend/kitex_gen/publish"
	"tiktok-backend/pkg/errno"
)

// PublishAction 登录用户选择视频上传
// 难点1：处理流数据
func PublishAction(ctx context.Context, c *app.RequestContext) {
	token := c.PostForm("token")
	title := c.PostForm("title")

	// 文件读取
	fileHeader, err := c.FormFile("data")
	if err != nil {
		pack.SendPublishActionResponse(c, err)
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		pack.SendPublishActionResponse(c, errno.ConvertErr(err))
		return
	}
	defer file.Close()

	//处理视频数据
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		pack.SendPublishActionResponse(c, errno.ConvertErr(err))
		return
	}

	// rpc
	err = rpc.PublishAction(context.Background(), &publish.DouyinPublishActionRequest{
		Token: token,
		Title: title,
		Data: buf.Bytes(),
	})
	if err != nil {
		pack.SendPublishActionResponse(c, errno.ConvertErr(err))
		return
	}

	pack.SendPublishActionResponse(c, errno.Success)
}

// PublishList 用户的视频发布列表，直接列出用户所有投稿过的视频
func PublishList(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	userIdStr := c.Query("user_id")

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		pack.SendPublishListResponse(c, err, nil)
		return
	}

	// rpc
	videoList, err := rpc.PublishList(context.Background(), &publish.DouyinPublishListRequest{
		Token:  token,
		UserId: userId,
	})
	if err != nil {
		pack.SendPublishListResponse(c, err, nil)
		return
	}

	pack.SendPublishListResponse(c, errno.Success, videoList)
}
