package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"tiktok-backend/kitex_gen/user"
	"tiktok-backend/kitex_gen/user/userservice"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/errno"
	"tiktok-backend/pkg/middleware"
	"time"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress}) // 服务注册发现中心
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
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
	userClient = c
}

// UserLogin check user info
func UserLogin(ctx context.Context, req *user.DouyinUserLoginRequest) (int64, error) {
	resp, err := userClient.UserLogin(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.UserId, nil
}

// UserRegister register user info
func UserRegister(ctx context.Context, req *user.DouyinUserRegisterRequest) (int64, string, error) {
	resp, err := userClient.UserRegister(ctx, req)
	if err != nil {
		return 0, "", err
	}
	if resp.StatusCode != 0 {
		return 0, "", errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.UserId, resp.Token, nil
}

// UserInfo get user info
func UserInfo(ctx context.Context, req *user.DouyinUserRequest) (*user.User, error) {
	resp, err := userClient.UserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.User, nil
}
