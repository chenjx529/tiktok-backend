package pack

import (
	"errors"
	"tiktok-backend/kitex_gen/feed"
	"tiktok-backend/pkg/errno"
)

// BuildFeedResp build baseResp from error
func BuildFeedResp(err error) *feed.DouyinFeedResponse {
	if err == nil {
		return FeedResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return FeedResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return FeedResp(s)
}

func FeedResp(err errno.ErrNo) *feed.DouyinFeedResponse {
	return &feed.DouyinFeedResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
