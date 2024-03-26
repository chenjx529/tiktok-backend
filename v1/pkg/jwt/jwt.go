package jwt

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"tiktok-backend/cmd/api/rpc"
	"tiktok-backend/kitex_gen/user"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/errno"
	"time"
)

var jwtMiddleware *jwt.HertzJWTMiddleware

func InitJwtMiddleware() {
	mw, err := jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
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
		//IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} { // 在登录成功后的每次请求中，用于设置从 token 提取用户信息的函数,存入请求上下文当中以备后续使用
		//	claims := jwt.ExtractClaims(ctx, c)
		//	return int64(claims[constants.IdentityKey].(float64)) // 在请求上下文中保存 id
		//},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, tokenStr string, expire time.Time) { // 设置登录返回消息
			token, err := jwt2.Parse(
				tokenStr,
				func(t *jwt2.Token) (interface{}, error) {
					if jwt2.GetSigningMethod("HS256") != t.Method {
						return nil, errors.New("invalid signing algorithm")
					}
					return []byte(constants.SecretKey), nil
				})
			if err != nil {
				panic(err)
			}
			claims := jwt.ExtractClaimsFromToken(token)
			userId := claims[constants.IdentityKey]
			// 我服了：proto3 由于字段为默认值（比如0值、空串、false等），导致输出json对应字段被隐藏
			Err := errno.ConvertErr(errno.Success)
			c.JSON(http.StatusOK, map[string]interface{}{
				"status_code": Err.ErrCode,
				"status_msg":  Err.ErrMsg,
				"user_id":     userId,
				"token":       tokenStr,
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) { // 设置 jwt 授权失败后的响应函数，message从 HTTPStatusMessageFunc 来
			c.JSON(code, map[string]interface{}{
				"status_code": errno.AuthorizationFailedErrCode,
				"status_msg":  message,
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) { // 配合 HertzJWTMiddleware.LoginHandler 使用，登录时触发，用于认证用户的登录信息。
			//var loginVar user.DouyinUserLoginRequest
			//if err := c.Bind(&loginVar); err != nil {
			//	return "", jwt.ErrMissingLoginValues
			//}

			userName := c.Query("username")
			passWord := c.Query("password")

			if len(userName) == 0 || len(passWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return rpc.UserLogin(context.Background(), &user.DouyinUserLoginRequest{
				Username: userName,
				Password: passWord,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		panic("JWT Error:" + err.Error())
	}

	jwtMiddleware = mw
}

func CreateToken(uid int64) (string, error) {
	token := jwt2.New(jwt2.GetSigningMethod(jwtMiddleware.SigningAlgorithm))
	claims := token.Claims.(jwt2.MapClaims)

	expire := jwtMiddleware.TimeFunc().UTC().Add(jwtMiddleware.Timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = jwtMiddleware.TimeFunc().Unix()
	claims[constants.IdentityKey] = uid

	tokenString, err := token.SignedString(jwtMiddleware.Key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetClaimsFromTokenStr(tokenStr string) (map[string]interface{}, error) {
	token, err := jwtMiddleware.ParseTokenString(tokenStr)
	if err != nil {
		return nil, err
	}
	return jwt.ExtractClaimsFromToken(token), nil
}

func GetJwtMiddleware() *jwt.HertzJWTMiddleware {
	return jwtMiddleware
}
