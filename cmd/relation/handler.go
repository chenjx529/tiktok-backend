package main

import (
	"context"
	"tiktok-backend/cmd/relation/pack"
	"tiktok-backend/cmd/relation/service"
	relation "tiktok-backend/kitex_gen/relation"
	"tiktok-backend/pkg/errno"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	resp = new(relation.DouyinRelationActionResponse)

	if len(req.Token) == 0 || req.ToUserId == 0 || req.ActionType == 0 {
		resp = pack.BuildRelationActionResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		resp = pack.BuildRelationActionResp(err)
		return resp, nil
	}
	resp = pack.BuildRelationActionResp(errno.Success)
	return resp, nil
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	resp = new(relation.DouyinRelationFollowListResponse)

	if req.UserId == 0 {
		resp = pack.BuildRelationFollowListResp(errno.ParamErr)
		return resp, nil
	}

	userList, err := service.NewRelationFollowListService(ctx).RelationFollowList(req)
	if err != nil {
		resp = pack.BuildRelationFollowListResp(err)
		return resp, nil
	}
	resp = pack.BuildRelationFollowListResp(errno.Success)
	resp.UserList = userList
	return resp, nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	resp = new(relation.DouyinRelationFollowerListResponse)

	if req.UserId == 0 {
		resp = pack.BuildRelationFollowerListResp(errno.ParamErr)
		return resp, nil
	}

	userList, err := service.NewRelationFollowerListService(ctx).RelationFollowerList(req)
	if err != nil {
		resp = pack.BuildRelationFollowerListResp(err)
		return resp, nil
	}
	resp = pack.BuildRelationFollowerListResp(errno.Success)
	resp.UserList = userList
	return resp, nil
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (resp *relation.DouyinRelationFriendListResponse, err error) {
	resp = new(relation.DouyinRelationFriendListResponse)

	if req.UserId == 0 {
		resp = pack.BuildRelationFriendListResp(errno.ParamErr)
		return resp, nil
	}

	userList, err := service.NewRelationFriendListService(ctx).RelationFriendList(req)
	if err != nil {
		resp = pack.BuildRelationFriendListResp(err)
		return resp, nil
	}
	resp = pack.BuildRelationFriendListResp(errno.Success)
	resp.UserList = userList
	return resp, nil
}
