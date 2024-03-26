package pack

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"tiktok-backend/kitex_gen/user"
	"tiktok-backend/pkg/errno"
)

func SendUserRegisterResponse(c *app.RequestContext, err error, userId int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserRegisterResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserId:     userId,
		Token:      token,
	})
}

func SendUserInfoResponse(c *app.RequestContext, err error, userinfo *user.User) {
	Err := errno.ConvertErr(err)
	var tmp *User
	if userinfo != nil {
		tmp = buildUserInfo(userinfo)
	}
	c.JSON(http.StatusOK, UserInfoResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		User:       tmp,
	})
}

func buildUserInfo(kitex_user *user.User) *User {
	return &User{
		Id:              kitex_user.Id,              // 用户id
		Name:            kitex_user.Name,            // 用户名称
		FollowCount:     kitex_user.FollowCount,     // 关注总数
		FollowerCount:   kitex_user.FollowerCount,   // 粉丝总数
		IsFollow:        kitex_user.IsFollow,        // true-已关注，false-未关注
		Avatar:          kitex_user.Avatar,          // 用户头像
		BackgroundImage: kitex_user.BackgroundImage, // 用户个人页顶部大图
		Signature:       kitex_user.Signature,       // 个人简介
		TotalFavorited:  kitex_user.TotalFavorited,  // 获赞数量
		WorkCount:       kitex_user.WorkCount,       // 作品数量
		FavoriteCount:   kitex_user.FavoriteCount,   // 点赞数量
	}
}
