package pack

import (
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/favorite"
)

func BuildVideoList(loginId int64, videoData []*db.Video, userMap map[int64]*db.User, favoriteSet map[int64]struct{}, followSet map[int64]struct{}) []*favorite.Video {
	videoList := make([]*favorite.Video, 0)
	for _, video := range videoData {
		// 视频用户
		user, _ := userMap[video.UserId]

		// 点赞 & 关注
		isFollow := false
		isFavorite := false
		if loginId != 0 {
			if _, ok := favoriteSet[int64(video.ID)]; ok {
				isFavorite = true
			}
			if _, ok := followSet[video.UserId]; ok {
				isFollow = true
			}
		}

		// 格式化
		videoList = append(videoList, videoInfo(video, userInfo(user, isFollow), isFavorite))
	}
	return videoList
}

func videoInfo(dbvideo *db.Video, author *favorite.User, isFavorite bool) *favorite.Video {
	return &favorite.Video{
		Id:            int64(dbvideo.ID),     // 视频唯一标识
		Author:        author,                // 视频作者信息
		PlayUrl:       dbvideo.PlayUrl,       // 视频播放地址
		CoverUrl:      dbvideo.CoverUrl,      // 视频封面地址
		FavoriteCount: dbvideo.FavoriteCount, // 视频的点赞总数
		CommentCount:  dbvideo.CommentCount,  // 视频的评论总数
		Title:         dbvideo.Title,         // 视频标题
		IsFavorite:    isFavorite,            // true-已点赞，false-未点赞
	}
}

func userInfo(dbuser *db.User, isFollow bool) *favorite.User {
	return &favorite.User{
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
