package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"tiktok-backend/kitex_gen/message"
	"tiktok-backend/kitex_gen/message/messageservice"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/middleware"
	"time"
)

var messageClient messageservice.Client

func initMessageRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress}) // 服务注册发现中心
	if err != nil {
		panic(err)
	}

	c, err := messageservice.NewClient(
		constants.MessageServiceName,
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
	messageClient = c
}

// MessageChat 当前登录用户和其他指定用户的聊天消息记录
func MessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) {

}

// MessageAction 登录用户对消息的相关操作，目前只支持消息发送
func MessageAction(ctx context.Context, req *message.DouyinMessageActionRequest) {

}