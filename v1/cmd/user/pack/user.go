package pack

import (
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/user"
)

func UserInfo(dbuser *db.User, isFollow bool) *user.User {
	return &user.User{
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
