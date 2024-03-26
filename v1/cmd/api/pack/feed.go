package pack

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"tiktok-backend/kitex_gen/feed"
	"tiktok-backend/pkg/errno"
)

func SendFeedResponse(c *app.RequestContext, err error, videoList []*feed.Video, nextTime int64) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, FeedResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		VideoList:  buildFeedVideoListInfo(videoList),
		NextTime:   nextTime,
	})
}

// buildFeedVideoListInfo pack video list info
func buildFeedVideoListInfo(videoData []*feed.Video) []*Video {
	videoList := make([]*Video, 0)
	for _, video := range videoData {
		videoList = append(videoList, buildFeedVideoInfo(video, buildFeedUserInfo(video.Author)))
	}
	return videoList
}

func buildFeedUserInfo(kitex_user *feed.User) *User {
	return &User{
		Id:              kitex_user.Id,              // 用户id
		Name:            kitex_user.Name,            // 用户名称
		FollowCount:     kitex_user.FollowCount,     // 关注总数
		FollowerCount:   kitex_user.FollowerCount,   // 粉丝总数
		Avatar:          kitex_user.Avatar,          //用户头像
		BackgroundImage: kitex_user.BackgroundImage, //用户个人页顶部大图
		Signature:       kitex_user.Signature,       //个人简介
		TotalFavorited:  kitex_user.TotalFavorited,  //获赞数量
		WorkCount:       kitex_user.WorkCount,       //作品数量
		FavoriteCount:   kitex_user.FavoriteCount,   //点赞数量
		IsFollow:        kitex_user.IsFollow,        // true-已关注，false-未关注
	}
}

func buildFeedVideoInfo(kitex_video *feed.Video, author *User) *Video {
	return &Video{
		Id:            kitex_video.Id,            // 视频唯一标识
		Author:        author,                    // 视频作者信息
		PlayUrl:       kitex_video.PlayUrl,       // 视频播放地址
		CoverUrl:      kitex_video.CoverUrl,      // 视频封面地址
		FavoriteCount: kitex_video.FavoriteCount, // 视频的点赞总数
		CommentCount:  kitex_video.CommentCount,  // 视频的评论总数
		Title:         kitex_video.Title,         // 视频标题
		IsFavorite:    kitex_video.IsFavorite,    // true-已点赞，false-未点赞
	}
}
