package pack

import (
	"errors"
	"tiktok-backend/kitex_gen/favorite"
	"tiktok-backend/pkg/errno"
)

func BuildFavoriteActionResp(err error) *favorite.DouyinFavoriteActionResponse {
	if err == nil {
		return favoriteActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoriteActionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return favoriteActionResp(s)
}

func favoriteActionResp(err errno.ErrNo) *favorite.DouyinFavoriteActionResponse {
	return &favorite.DouyinFavoriteActionResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

// BuildFavoriteListResp 发送FavoriteListResponse
func BuildFavoriteListResp(err error) *favorite.DouyinFavoriteListResponse {
	if err == nil {
		return favoriteFollowListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoriteFollowListResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return favoriteFollowListResp(s)
}

func favoriteFollowListResp(err errno.ErrNo) *favorite.DouyinFavoriteListResponse {
	return &favorite.DouyinFavoriteListResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
