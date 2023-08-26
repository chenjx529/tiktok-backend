package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"tiktok-backend/cmd/api/handlers"
	"tiktok-backend/cmd/api/rpc"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/errno"
	"tiktok-backend/pkg/jwt"
	"tiktok-backend/pkg/tracer"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

func main() {
	Init()
	r := server.New(
		server.WithHostPorts("127.0.0.1"+constants.ApiServicePort),
		server.WithHandleMethodNotAllowed(true), // 全局处理 HTTP 404 与 405 请求
	)

	authMiddleware, err := jwt.NewJwtMiddleware()

	if err != nil {
		hlog.Fatal("JWT Error:" + err.Error())
	}

	// 默认的 panic 处理函数
	r.Use(recovery.Recovery(recovery.WithRecoveryHandler(
		func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
			hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
			c.JSON(consts.StatusInternalServerError, map[string]interface{}{
				"code":    errno.ServiceErrCode,
				"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
			})
		})))

	// 不需要token
	r.GET("/douyin/feed/", handlers.Feed)                      // 视频流
	r.POST("/douyin/user/login/", authMiddleware.LoginHandler) // 用户登录
	r.POST("/douyin/user/register/", handlers.UserRegister)    // 用户注册
	r.GET("/douyin/comment/list/", handlers.CommentList)       // 视频评论列表

	douyin := r.Group("/douyin")
	douyin.Use(authMiddleware.MiddlewareFunc())

	douyin.GET("/user", handlers.UserInfo) // 用户信息

	publish := douyin.Group("/publish")
	publish.POST("/action/", handlers.PublishAction) // 视频投稿
	publish.GET("/list/", handlers.PublishList)       // 发布列表

	favorite := douyin.Group("/favorite")
	favorite.POST("/action/", handlers.FavoriteAction) // 点赞
	favorite.GET("/list/", handlers.FavoriteList)      // 喜欢列表

	comment := douyin.Group("/comment")
	comment.POST("/action/", handlers.CommentAction) // 评论

	relation := douyin.Group("/relation")
	relation.POST("/action/", handlers.RelationAction)             // 关注
	relation.GET("/follow/list/", handlers.RelationFollowList)     // 用户关注列表
	relation.GET("/follower/list/", handlers.RelationFollowerList) // 用户粉丝列表
	relation.GET("/friend/list/", handlers.RelationFriendList)     // 用户好友

	message := douyin.Group("/message")
	message.GET("/chat/", handlers.MessageChat)      // 聊天记录
	message.POST("/action/", handlers.MessageAction) // 消息发送

	// 全局处理 HTTP 404 与 405 请求
	r.NoRoute(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no route")
	})
	r.NoMethod(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no method")
	})
	r.Spin()
}
