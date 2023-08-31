package pack

import (
	"errors"
	"tiktok-backend/kitex_gen/comment"
	"tiktok-backend/pkg/errno"
)

// BuildCommentActionResp 发送 CommentActionResponse
func BuildCommentActionResp(err error) *comment.DouyinCommentActionResponse {
	if err == nil {
		return commentActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentActionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return commentActionResp(s)
}

func commentActionResp(err errno.ErrNo) *comment.DouyinCommentActionResponse {
	return &comment.DouyinCommentActionResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

// BuildCommentListResp 发送 CommentActionResponse
func BuildCommentListResp(err error) *comment.DouyinCommentListResponse {
	if err == nil {
		return commentListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentListResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return commentListResp(s)
}

func commentListResp(err errno.ErrNo) *comment.DouyinCommentListResponse {
	return &comment.DouyinCommentListResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
