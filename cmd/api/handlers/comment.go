package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"tiktok-backend/cmd/api/pack"
	"tiktok-backend/cmd/api/rpc"
	"tiktok-backend/kitex_gen/comment"
	"tiktok-backend/pkg/errno"
)

// CommentAction 已登录用户对视频进行评论
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var commentVar comment.DouyinCommentActionRequest
	if err := c.Bind(&commentVar); err != nil {
		pack.SendCommentActionResponse(c, err, -1)
		return
	}

	if commentVar.VideoId == 0 {
		pack.SendCommentActionResponse(c, errno.ParamErr, -1)
		return
	}

	commentData, err := rpc.CommentAction(context.Background(), &commentVar)
	if err != nil {
		pack.SendCommentActionResponse(c, err, -1)
	}
	pack.SendCommentActionResponse(c, errno.Success, commentData.Id)
}

// CommentList 查看某视频的评论列表
func CommentList(ctx context.Context, c *app.RequestContext) {
	videoIdStr := c.Query("video_id")
	token := c.DefaultQuery("token", "")
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		pack.SendCommentListResponse(c, err, nil)
	}

	req := &comment.DouyinCommentListRequest{
		Token:   token,
		VideoId: videoId,
	}

	commentList, err := rpc.CommentList(context.Background(), req)
	if err != nil {
		pack.SendCommentListResponse(c, err, nil)
	}
	pack.SendCommentListResponse(c, errno.Success, commentList)
}
