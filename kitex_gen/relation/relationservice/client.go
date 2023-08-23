// Code generated by Kitex v0.7.0. DO NOT EDIT.

package relationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	relation "tiktok-backend/kitex_gen/relation"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RelationAction(ctx context.Context, Req *relation.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *relation.DouyinRelationActionResponse, err error)
	RelationFollowList(ctx context.Context, Req *relation.DouyinRelationFollowListRequest, callOptions ...callopt.Option) (r *relation.DouyinRelationFollowListResponse, err error)
	RelationFollowerList(ctx context.Context, Req *relation.DouyinRelationFollowerListRequest, callOptions ...callopt.Option) (r *relation.DouyinRelationFollowerListResponse, err error)
	RelationFriendList(ctx context.Context, Req *relation.DouyinRelationFriendListRequest, callOptions ...callopt.Option) (r *relation.DouyinRelationFriendListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kRelationServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) RelationAction(ctx context.Context, Req *relation.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *relation.DouyinRelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationAction(ctx, Req)
}

func (p *kRelationServiceClient) RelationFollowList(ctx context.Context, Req *relation.DouyinRelationFollowListRequest, callOptions ...callopt.Option) (r *relation.DouyinRelationFollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationFollowList(ctx, Req)
}

func (p *kRelationServiceClient) RelationFollowerList(ctx context.Context, Req *relation.DouyinRelationFollowerListRequest, callOptions ...callopt.Option) (r *relation.DouyinRelationFollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationFollowerList(ctx, Req)
}

func (p *kRelationServiceClient) RelationFriendList(ctx context.Context, Req *relation.DouyinRelationFriendListRequest, callOptions ...callopt.Option) (r *relation.DouyinRelationFriendListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationFriendList(ctx, Req)
}
