package pack

import (
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/user"
)

func UserInfo(dbuser *db.User, isFollow bool) *user.User {
	userInfo := &user.User{
		Id:              int64(dbuser.ID),
		Name:            dbuser.Name,
		FollowCount:     dbuser.FollowCount,
		FollowerCount:   dbuser.FollowerCount,
		IsFollow:        isFollow,  // true-已关注，false-未关注
		Avatar:          dbuser.Avatar,
		BackgroundImage: dbuser.BackgroundImage,
		Signature:       dbuser.Signature,
		TotalFavorited:  dbuser.TotalFavorited,
		WorkCount:       dbuser.WorkCount,
		FavoriteCount:   dbuser.FavoriteCount,
	}
	return userInfo
}
