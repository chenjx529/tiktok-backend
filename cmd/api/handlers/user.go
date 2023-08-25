package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"tiktok-backend/cmd/api/pack"
	"tiktok-backend/cmd/api/rpc"
	"tiktok-backend/kitex_gen/user"
	"tiktok-backend/pkg/errno"
)

// UserRegister register uesr info
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var registerVar user.DouyinUserRegisterRequest
	if err := c.Bind(&registerVar); err != nil {
		pack.SendUserRegisterResponse(c, errno.ConvertErr(err), -1, "")
		return
	}

	if len(registerVar.Username) == 0 || len(registerVar.Password) == 0 {
		pack.SendUserRegisterResponse(c, errno.ParamErr, -1, "")
		return
	}

	uid, token, err := rpc.UserRegister(context.Background(), &registerVar)
	if err != nil {
		pack.SendUserRegisterResponse(c, errno.ConvertErr(err), -1, "")
		return
	}
	pack.SendUserRegisterResponse(c, errno.Success, uid, token)
}

// UserInfo get user info
func UserInfo(ctx context.Context, c *app.RequestContext) {
	userIdStr := c.Query("user_id")
	token := c.DefaultQuery("token", "")

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		pack.SendUserInfoResponse(c, errno.ConvertErr(err), nil)
		return
	}

	userinfo, err := rpc.UserInfo(context.Background(), &user.DouyinUserRequest{
		UserId: userId,
		Token:  token,
	})
	if err != nil {
		pack.SendUserInfoResponse(c, errno.ConvertErr(err), nil)
		return
	}
	pack.SendUserInfoResponse(c, errno.Success, userinfo)
}
