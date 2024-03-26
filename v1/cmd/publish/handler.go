package main

import (
	"context"
	"tiktok-backend/cmd/publish/pack"
	"tiktok-backend/cmd/publish/service"
	publish "tiktok-backend/kitex_gen/publish"
	"tiktok-backend/pkg/errno"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {
	resp = new(publish.DouyinPublishActionResponse)

	if len(req.Token) == 0 || len(req.Title) == 0 || req.Data == nil {
		resp = pack.BuildPublishActionResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewPublishActionService(ctx).PublishAction(req)
	if err != nil {
		resp = pack.BuildPublishActionResp(err)
		return resp, nil
	}

	resp = pack.BuildPublishActionResp(errno.Success)
	return resp, nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	resp = new(publish.DouyinPublishListResponse)

	if req.UserId <= 0 {
		resp = pack.BuildPublishListResp(errno.ParamErr)
		return resp, nil
	}

	videoList, err := service.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		resp = pack.BuildPublishListResp(err)
		return resp, nil
	}

	resp = pack.BuildPublishListResp(errno.Success)
	resp.VideoList = videoList
	return resp, nil
}
