package pack

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"tiktok-backend/kitex_gen/relation"
	"tiktok-backend/pkg/errno"
)

func SendRelationActionResponse(c *app.RequestContext, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, RelationActionResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
	})
}

func SendRelationFollowListResponse(c *app.RequestContext, err error, userList []*relation.User) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, RelationFollowListResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserList:   buildRelationFollowList(userList),
	})
}

func SendRelationFollowerListResponse(c *app.RequestContext, err error, userList []*relation.User) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, RelationFollowerListResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserList:   buildRelationFollowList(userList),
	})
}

func SendRelationFriendListResponse(c *app.RequestContext, err error, userList []*relation.FriendUser) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, RelationFriendListResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserList:   buildRelationFriendList(userList),
	})
}

// buildRelationFollowList pack video list info
func buildRelationFollowList(relationUserList []*relation.User) []*User {
	userList := make([]*User, 0)
	for _, v := range relationUserList {
		userList = append(userList, buildRelationUserInfo(v))
	}
	return userList
}

func buildRelationUserInfo(kitex_user *relation.User) *User {
	return &User{
		Id:              kitex_user.Id,              // 用户id
		Name:            kitex_user.Name,            // 用户名称
		FollowCount:     kitex_user.FollowCount,     // 关注总数
		FollowerCount:   kitex_user.FollowerCount,   // 粉丝总数
		Avatar:          kitex_user.Avatar,          // 用户头像
		BackgroundImage: kitex_user.BackgroundImage, // 用户个人页顶部大图
		Signature:       kitex_user.Signature,       // 个人简介
		TotalFavorited:  kitex_user.TotalFavorited,  // 获赞数量
		WorkCount:       kitex_user.WorkCount,       // 作品数量
		FavoriteCount:   kitex_user.FavoriteCount,   // 点赞数量
		IsFollow:        kitex_user.IsFollow,        // true-已关注，false-未关注
	}
}

// buildRelationFriendList pack video list info
func buildRelationFriendList(userList []*relation.FriendUser) []*FriendUser {
	videoList := make([]*FriendUser, 0)
	for _, user := range userList {
		videoList = append(videoList, buildRelationFriendUserInfo(user))
	}
	return videoList
}

func buildRelationFriendUserInfo(kitex_user *relation.FriendUser) *FriendUser {
	return &FriendUser{
		Id:              kitex_user.Id,              // 用户id
		Name:            kitex_user.Name,            // 用户名称
		FollowCount:     kitex_user.FollowCount,     // 关注总数
		FollowerCount:   kitex_user.FollowerCount,   // 粉丝总数
		Avatar:          kitex_user.Avatar,          // 用户头像
		BackgroundImage: kitex_user.BackgroundImage, // 用户个人页顶部大图
		Signature:       kitex_user.Signature,       // 个人简介
		TotalFavorited:  kitex_user.TotalFavorited,  // 获赞数量
		WorkCount:       kitex_user.WorkCount,       // 作品数量
		FavoriteCount:   kitex_user.FavoriteCount,   // 点赞数量
		Message:         kitex_user.Message,         // 和该好友的最新聊天消息
		MsgType:         kitex_user.MsgType,         // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
		IsFollow:        kitex_user.IsFollow,        // true-已关注，false-未关注
	}
}
