package pack

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"tiktok-backend/kitex_gen/user"
	"tiktok-backend/pkg/errno"
)

func BuildUserInfo(kitex_user *user.User) *User {
	return &User{
		Id:              kitex_user.Id,
		Name:            kitex_user.Name,
		FollowCount:     kitex_user.FollowCount,
		FollowerCount:   kitex_user.FollowerCount,
		IsFollow:        kitex_user.IsFollow, // true-已关注，false-未关注
		Avatar:          kitex_user.Avatar,
		BackgroundImage: kitex_user.BackgroundImage,
		Signature:       kitex_user.Signature,
		TotalFavorited:  kitex_user.TotalFavorited,
		WorkCount:       kitex_user.WorkCount,
		FavoriteCount:   kitex_user.FavoriteCount,
	}
}

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
	c.JSON(http.StatusOK, UserInfoResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		User:       BuildUserInfo(userinfo),
	})
}
