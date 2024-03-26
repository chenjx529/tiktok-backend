package main

import (
	"context"
	"tiktok-backend/cmd/comment/pack"
	"tiktok-backend/cmd/comment/service"
	comment "tiktok-backend/kitex_gen/comment"
	"tiktok-backend/pkg/errno"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	resp = new(comment.DouyinCommentActionResponse)

	if len(req.Token) == 0 || req.VideoId == 0 {
		resp = pack.BuildCommentActionResp(errno.ParamErr)
		return resp, nil
	}

	commentData, err := service.NewCommentActionService(ctx).CommentAction(req)
	if err != nil {
		resp = pack.BuildCommentActionResp(err)
		return resp, nil
	}

	resp = pack.BuildCommentActionResp(errno.Success)
	resp.Comment = commentData
	return resp, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	resp = new(comment.DouyinCommentListResponse)

	if req.VideoId <= 0 {
		resp = pack.BuildCommentListResp(errno.ParamErr)
		return resp, nil
	}

	commentList, err := service.NewCommentListService(ctx).CommentList(req)
	if err != nil {
		resp = pack.BuildCommentListResp(err)
		return resp, nil
	}

	resp = pack.BuildCommentListResp(errno.Success)
	resp.CommentList = commentList
	return resp, nil
}
