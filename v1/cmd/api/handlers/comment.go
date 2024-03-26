package handlers

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"tiktok-backend/cmd/api/pack"
	"tiktok-backend/cmd/api/rpc"
	"tiktok-backend/kitex_gen/comment"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/errno"
)

// CommentAction 已登录用户对视频进行评论
// 视频评论的总数需要增加
func CommentAction(ctx context.Context, c *app.RequestContext) {
	tokenStr := c.Query("token")
	videoIdStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")
	commentIdStr := c.DefaultQuery("comment_id", "0")
	commentTextStr := c.DefaultQuery("comment_text", "")

	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		pack.SendCommentActionResponse(c, err, nil)
		return
	}

	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		pack.SendCommentActionResponse(c, err, nil)
		return
	}

	commentId, err := strconv.ParseInt(commentIdStr, 10, 64)
	if err != nil {
		pack.SendCommentActionResponse(c, err, nil)
		return
	}

	if actionType != constants.Comment && actionType != constants.UnComment {
		pack.SendCommentActionResponse(c, errors.New("actionType error"), nil)
		return
	}

	commentData, err := rpc.CommentAction(context.Background(), &comment.DouyinCommentActionRequest{
		Token:      tokenStr,
		VideoId:    videoId,
		ActionType: int32(actionType),
		CommentId: commentId,
		CommentText: commentTextStr,
	})
	if err != nil {
		pack.SendCommentActionResponse(c, err, nil)
	}

	pack.SendCommentActionResponse(c, errno.Success, commentData)
}

// CommentList 查看某视频的评论列表
func CommentList(ctx context.Context, c *app.RequestContext) {
	videoIdStr := c.Query("video_id")
	token := c.Query("token")

	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		pack.SendCommentListResponse(c, err, nil)
	}

	commentList, err := rpc.CommentList(context.Background(), &comment.DouyinCommentListRequest{
		Token:   token,
		VideoId: videoId,
	})
	if err != nil {
		pack.SendCommentListResponse(c, err, nil)
	}

	pack.SendCommentListResponse(c, errno.Success, commentList)
}
