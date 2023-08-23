package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
	"tiktok-backend/cmd/api/rpc"
	"tiktok-backend/kitex_gen/user"
	"tiktok-backend/pkg/errno"
)

// UserRegister register uesr info
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var registerVar user.DouyinUserRegisterRequest
	if err := c.Bind(&registerVar); err != nil {
		sendUserRegisterResponse(c, errno.ConvertErr(err), -1, "")
		return
	}

	if len(registerVar.Username) == 0 || len(registerVar.Password) == 0 {
		sendUserRegisterResponse(c, errno.ParamErr, -1, "")
		return
	}

	uid, token, err := rpc.UserRegister(context.Background(), &registerVar)
	if err != nil {
		sendUserRegisterResponse(c, errno.ConvertErr(err), -1, "")
		return
	}
	sendUserRegisterResponse(c, errno.Success, uid, token)
}

func sendUserRegisterResponse(c *app.RequestContext, err error, userId int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, map[string]interface{}{
		"status_code": Err.ErrCode,
		"status_msg":  Err.ErrMsg,
		"user_id":     userId,
		"token":      token,
	})
}

// UserInfo get user info
func UserInfo(ctx context.Context, c *app.RequestContext) {
	userIdStr := c.Query("user_id")
	token := c.DefaultQuery("token", "")

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		sendUserInfoResponse(c, errno.ConvertErr(err), nil)
		return
	}

	userinfo, err := rpc.UserInfo(context.Background(), &user.DouyinUserRequest{
		UserId: userId,
		Token:  token,
	})
	if err != nil {
		sendUserInfoResponse(c, errno.ConvertErr(err), nil)
	}
	sendUserInfoResponse(c, errno.Success, userinfo)
}

func sendUserInfoResponse(c *app.RequestContext, err error, userinfo *user.User) {
	Err := errno.ConvertErr(err)

	c.JSON(http.StatusOK, map[string]interface{}{
		"status_code": Err.ErrCode,
		"status_msg":  Err.ErrMsg,
		"user":       userinfo,
	})
}
