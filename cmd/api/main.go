package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"tiktok-backend/cmd/api/handlers"
	"tiktok-backend/cmd/api/rpc"
	"tiktok-backend/kitex_gen/user"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/errno"
	jwt2 "tiktok-backend/pkg/jwt"
	"tiktok-backend/pkg/tracer"
	"time"
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
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		IdentityKey: constants.IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims { // 登录时为 token 添加自定义负载信息的函数
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string { // jwt 校验流程产生错误, 对应 error 将以参数的形式传递给 HTTPStatusMessageFunc
			var errNo errno.ErrNo
			if errors.As(e, &errNo) {
				return errNo.ErrMsg
			} else {
				return e.Error()
			}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} { // 在登录成功后的每次请求中，用于设置从 token 提取用户信息的函数,存入请求上下文当中以备后续使用
			claims := jwt.ExtractClaims(ctx, c)
			return int64(int(claims[constants.IdentityKey].(float64))) // 在请求上下文中保存 id
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) { // 设置登录返回消息
			userId, err := jwt2.ParseToken(token)
			if err != nil {
				panic(err)
			}
			// 我服了：proto3 由于字段为默认值（比如0值、空串、false等），导致输出json对应字段被隐藏
			Err := errno.ConvertErr(errno.Success)
			c.JSON(http.StatusOK, map[string]interface{}{
				"status_code": Err.ErrCode,
				"status_msg":  Err.ErrMsg,
				"user_id":     userId,
				"token":      token,
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) { // 设置 jwt 授权失败后的响应函数，message从 HTTPStatusMessageFunc 来
			c.JSON(code, map[string]interface{}{
				"code":    errno.AuthorizationFailedErrCode,
				"message": message,
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) { // 配合 HertzJWTMiddleware.LoginHandler 使用，登录时触发，用于认证用户的登录信息。
			var loginVar user.DouyinUserLoginRequest
			if err := c.Bind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.Username) == 0 || len(loginVar.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return rpc.UserLogin(context.Background(), &loginVar)
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

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

	douyin := r.Group("/douyin")
	douyin.Use(authMiddleware.MiddlewareFunc())

	douyin.GET("/user", handlers.UserInfo) // 用户信息

	publish := douyin.Group("/publish")
	publish.POST("/action/", handlers.PublishActiion) // 视频投稿
	publish.GET("/list/", handlers.PublishList)       // 发布列表

	favorite := douyin.Group("/favorite")
	favorite.POST("/action/", handlers.FavoriteAction) // 点赞
	favorite.GET("/list/", handlers.FavoriteList)      // 喜欢列表

	comment := douyin.Group("/comment")
	comment.POST("/action/", handlers.CommentAction) // 评论
	comment.GET("/list/", handlers.CommentList)      // 视频评论列表

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
