package jwt

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"testing"
	"tiktok-backend/cmd/api/rpc"
	"tiktok-backend/kitex_gen/user"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/errno"
	"time"
)

func TestNewJWT(t *testing.T) {
	authMiddleware, _ := jwt.New(&jwt.HertzJWTMiddleware{
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
			//userId, err := jwt2.ParseToken(token)
			//if err != nil {
			//	panic(err)
			//}
			//// 我服了：proto3 由于字段为默认值（比如0值、空串、false等），导致输出json对应字段被隐藏
			//Err := errno.ConvertErr(errno.Success)
			//c.JSON(http.StatusOK, map[string]interface{}{
			//	"status_code": Err.ErrCode,
			//	"status_msg":  Err.ErrMsg,
			//	"user_id":     userId,
			//	"token":      token,
			//})
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

	tokenString, _, _ := authMiddleware.TokenGenerator(jwt.MapClaims{
		constants.IdentityKey: 3,
	})

	fmt.Println(tokenString)

	id, _ := ParseToken(tokenString)
	fmt.Println(id)

}
