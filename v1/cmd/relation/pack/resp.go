package pack

import (
	"errors"
	"tiktok-backend/kitex_gen/relation"
	"tiktok-backend/pkg/errno"
)

// BuildRelationActionResp 发送RelationActionResponse
func BuildRelationActionResp(err error) *relation.DouyinRelationActionResponse {
	if err == nil {
		return relationActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return relationActionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return relationActionResp(s)
}

func relationActionResp(err errno.ErrNo) *relation.DouyinRelationActionResponse {
	return &relation.DouyinRelationActionResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

// BuildRelationFollowListResp 发送RelationFollowListResponse
func BuildRelationFollowListResp(err error) *relation.DouyinRelationFollowListResponse {
	if err == nil {
		return relationFollowListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return relationFollowListResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return relationFollowListResp(s)
}

func relationFollowListResp(err errno.ErrNo) *relation.DouyinRelationFollowListResponse {
	return &relation.DouyinRelationFollowListResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

// BuildRelationFollowerListResp 发送RelationFollowerListResponse
func BuildRelationFollowerListResp(err error) *relation.DouyinRelationFollowerListResponse {
	if err == nil {
		return relationFollowerListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return relationFollowerListResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return relationFollowerListResp(s)
}

func relationFollowerListResp(err errno.ErrNo) *relation.DouyinRelationFollowerListResponse {
	return &relation.DouyinRelationFollowerListResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

// BuildRelationFriendListResp 发送RelationFriendListResponse
func BuildRelationFriendListResp(err error) *relation.DouyinRelationFriendListResponse {
	if err == nil {
		return relationFriendListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return relationFriendListResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return relationFriendListResp(s)
}

func relationFriendListResp(err errno.ErrNo) *relation.DouyinRelationFriendListResponse {
	return &relation.DouyinRelationFriendListResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
