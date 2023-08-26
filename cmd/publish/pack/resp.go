package pack

import (
	"errors"
	"tiktok-backend/kitex_gen/publish"
	"tiktok-backend/pkg/errno"
)

// BuildPublishActionResp 发送PublishActionResponse
func BuildPublishActionResp(err error) *publish.DouyinPublishActionResponse {
	if err == nil {
		return publishActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return publishActionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return publishActionResp(s)
}

func publishActionResp(err errno.ErrNo) *publish.DouyinPublishActionResponse {
	return &publish.DouyinPublishActionResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

// BuildPublishListResp 发送PublishActionResponse
func BuildPublishListResp(err error) *publish.DouyinPublishListResponse {
	if err == nil {
		return publishListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return publishListResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return publishListResp(s)
}

func publishListResp(err errno.ErrNo) *publish.DouyinPublishListResponse {
	return &publish.DouyinPublishListResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
