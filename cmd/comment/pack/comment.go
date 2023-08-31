package pack

import (
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/comment"
	"tiktok-backend/pkg/constants"
)

// BuildCommentList pack comment list info
func BuildCommentList(loginId int64, commentData []*db.Comment, userMap map[int64]*db.User, followSet map[int64]struct{}) []*comment.Comment {
	commentList := make([]*comment.Comment, 0)
	for _, com := range commentData {
		// 评论用户
		user, _ := userMap[com.UserId]

		// 关注
		isFollow := false
		if loginId != 0 {
			if _, ok := followSet[com.UserId]; ok {
				isFollow = true
			}
		}

		// 格式化
		commentList = append(commentList, BuildCommentInfo(com, user, isFollow))
	}
	return commentList
}

func BuildCommentInfo(com *db.Comment, dbuser *db.User, isFollow bool) *comment.Comment {

	return &comment.Comment{
		Id:         int64(com.ID),
		User:       buildCommentUserInfo(dbuser, isFollow),
		Content:    com.Contents,
		CreateDate: com.CreatedAt.Format(constants.TimeFormat),
	}
}

func buildCommentUserInfo(dbuser *db.User, isFollow bool) *comment.User {
	return &comment.User{
		Id:              int64(dbuser.ID),       // 用户id
		Name:            dbuser.Name,            // 用户名称
		FollowCount:     dbuser.FollowCount,     // 关注总数
		FollowerCount:   dbuser.FollowerCount,   // 粉丝总数
		Avatar:          dbuser.Avatar,          //用户头像
		BackgroundImage: dbuser.BackgroundImage, //用户个人页顶部大图
		Signature:       dbuser.Signature,       //个人简介
		TotalFavorited:  dbuser.TotalFavorited,  //获赞数量
		WorkCount:       dbuser.WorkCount,       //作品数量
		FavoriteCount:   dbuser.FavoriteCount,   //点赞数量
		IsFollow:        isFollow,               // true-已关注，false-未关注
	}
}
