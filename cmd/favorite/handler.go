package main

import (
	"context"
	"tiktok-backend/cmd/favorite/pack"
	"tiktok-backend/cmd/favorite/service"
	favorite "tiktok-backend/kitex_gen/favorite"
	"tiktok-backend/pkg/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	resp = new(favorite.DouyinFavoriteActionResponse)

	if len(req.Token) == 0 || req.VideoId == 0 || req.ActionType == 0 {
		resp = pack.BuildFavoriteActionResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp = pack.BuildFavoriteActionResp(err)
		return resp, nil
	}

	resp = pack.BuildFavoriteActionResp(errno.Success)
	return resp, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	resp = new(favorite.DouyinFavoriteListResponse)

	if req.UserId == 0 {
		resp = pack.BuildFavoriteListResp(errno.ParamErr)
		return resp, nil
	}

	videoList, err := service.NewFavoriteListService(ctx).FavoriteList(req)
	if err != nil {
		resp = pack.BuildFavoriteListResp(err)
		return resp, nil
	}

	resp = pack.BuildFavoriteListResp(errno.Success)
	resp.VideoList = videoList
	return resp, nil
}
