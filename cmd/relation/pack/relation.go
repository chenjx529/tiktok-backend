package pack

import (
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/relation"
)

// BuildFollowList pack user list info
func BuildFollowList(relationUserList []*db.User, followSet map[int64]struct{}) []*relation.User {
	userList := make([]*relation.User, 0)
	for _, user := range relationUserList {
		isFollow := false
		if _, ok := followSet[int64(user.ID)]; ok {
			isFollow = true
		}
		userList = append(userList, buildFollowInfo(user, isFollow))
	}
	return userList
}

func buildFollowInfo(dbuser *db.User, isFollow bool) *relation.User {
	return &relation.User{
		Id:              int64(dbuser.ID),       // 用户id
		Name:            dbuser.Name,            // 用户名称
		FollowCount:     dbuser.FollowCount,     // 关注总数
		FollowerCount:   dbuser.FollowerCount,   // 粉丝总数
		Avatar:          dbuser.Avatar,          // 用户头像
		BackgroundImage: dbuser.BackgroundImage, // 用户个人页顶部大图
		Signature:       dbuser.Signature,       // 个人简介
		TotalFavorited:  dbuser.TotalFavorited,  // 获赞数量
		WorkCount:       dbuser.WorkCount,       // 作品数量
		FavoriteCount:   dbuser.FavoriteCount,   // 点赞数量
		IsFollow:        isFollow,               // true-已关注，false-未关注
	}
}

// BuildFriendUserList pack user list info
func BuildFriendUserList(relationUserList []*db.User, followSet map[int64]struct{}, messageMap map[int64]*db.Message) []*relation.FriendUser {
	userList := make([]*relation.FriendUser, 0)
	for _, user := range relationUserList {
		isFollow := false
		if _, ok := followSet[int64(user.ID)]; ok {
			isFollow = true
		}
		var message *db.Message
		if v, ok := messageMap[int64(user.ID)]; ok {
			message = v
		}
		userList = append(userList, buildFriendUserInfo(user, isFollow, message))
	}
	return userList
}

func buildFriendUserInfo(dbuser *db.User, isFollow bool, message *db.Message) *relation.FriendUser {
	var content string
	var msgType int64
	if message != nil {
		content = message.Content
		if int64(dbuser.ID) == message.ToUserId {
			msgType = 1 // 好友是接受方，自己就是发送方
		} else {
			msgType = 0
		}
	}
	return &relation.FriendUser{
		Id:              int64(dbuser.ID),       // 用户id
		Name:            dbuser.Name,            // 用户名称
		FollowCount:     dbuser.FollowCount,     // 关注总数
		FollowerCount:   dbuser.FollowerCount,   // 粉丝总数
		Avatar:          dbuser.Avatar,          // 用户头像
		BackgroundImage: dbuser.BackgroundImage, // 用户个人页顶部大图
		Signature:       dbuser.Signature,       // 个人简介
		TotalFavorited:  dbuser.TotalFavorited,  // 获赞数量
		WorkCount:       dbuser.WorkCount,       // 作品数量
		FavoriteCount:   dbuser.FavoriteCount,   // 点赞数量
		IsFollow:        isFollow,               // true-已关注，false-未关注
		Message:         content,                // 和该好友的最新聊天消息
		MsgType:         msgType,                // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
	}
}
