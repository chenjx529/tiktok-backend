package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"tiktok-backend/kitex_gen/favorite"
	"tiktok-backend/kitex_gen/favorite/favoriteservice"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/errno"
	"tiktok-backend/pkg/middleware"
	"time"
)

var favoriteClient favoriteservice.Client

func initFavoriteRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress}) // 服务注册发现中心
	if err != nil {
		panic(err)
	}

	c, err := favoriteservice.NewClient(
		constants.FavoriteServiceName,
		client.WithMiddleware(middleware.CommonMiddleware), // 通用中间件
		client.WithInstanceMW(middleware.ClientMiddleware), // 客户端中间件
		client.WithMuxConnection(1),                        // 多路复用
		client.WithRPCTimeout(3*time.Second),               // 设置 rpc 调用超时时间
		client.WithConnectTimeout(50*time.Millisecond),     // 设置 rpc 连接超时时间
		client.WithFailureRetry(retry.NewFailurePolicy()),  // 重试，默认2次，可以设置重试次数，熔断
		client.WithSuite(trace.NewDefaultClientSuite()),    // 链路追踪，默认使用 OpenTracing GlobalTracer
		client.WithResolver(r),                             // 服务发现
	)
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

// FavoriteAction implement like and unlike operations
func FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) error {
	resp, err := favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return nil
}

// FavoriteList get favorite list info
func FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) ([]*favorite.Video, error) {
	resp, err := favoriteClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.VideoList, nil
}