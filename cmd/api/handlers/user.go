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

// UserRegister 新用户注册时提供用户名，密码即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token
// 难点1：用户名需要保证唯一
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var registerVar user.DouyinUserRegisterRequest
	if err := c.Bind(&registerVar); err != nil {
		pack.SendUserRegisterResponse(c, err, -1, "")
		return
	}

	if len(registerVar.Username) == 0 || len(registerVar.Password) == 0 {
		pack.SendUserRegisterResponse(c, errno.ParamErr, -1, "")
		return
	}

	uid, token, err := rpc.UserRegister(context.Background(), &registerVar)
	if err != nil {
		pack.SendUserRegisterResponse(c, err, -1, "")
		return
	}
	pack.SendUserRegisterResponse(c, errno.Success, uid, token)
}

// UserInfo 获取用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数
// 难点1：返回关注数和粉丝数（这个可以直接存在数据库中）
// 难点2：login用户是否关注当前查询用户
func UserInfo(ctx context.Context, c *app.RequestContext) {
	userIdStr := c.Query("user_id")
	token := c.DefaultQuery("token", "")

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		pack.SendUserInfoResponse(c, err, nil)
		return
	}

	userinfo, err := rpc.UserInfo(context.Background(), &user.DouyinUserRequest{
		UserId: userId,
		Token:  token,
	})
	if err != nil {
		pack.SendUserInfoResponse(c, err, nil)
		return
	}
	pack.SendUserInfoResponse(c, errno.Success, userinfo)
}
