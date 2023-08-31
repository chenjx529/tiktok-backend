package pack

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"tiktok-backend/kitex_gen/comment"
	"tiktok-backend/pkg/errno"
)

func SendCommentActionResponse(c *app.RequestContext, err error, com *comment.Comment) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, CommentActionResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		Comment:  buildCommentInfo(com, buildCommentUserInfo(com.User)),
	})
}

func SendCommentListResponse(c *app.RequestContext, err error, commentList []*comment.Comment) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, CommentListResponse{
		StatusCode:  Err.ErrCode,
		StatusMsg:   Err.ErrMsg,
		CommentList: buildCommentList(commentList),
	})
}

// buildCommentListInfo pack comment list info
func buildCommentList(commentData []*comment.Comment) []*Comment {
	commentList := make([]*Comment, 0)
	for _, v := range commentData {
		commentList = append(commentList, buildCommentInfo(v, buildCommentUserInfo(v.User)))
	}
	return commentList
}

func buildCommentInfo(kitex_comment *comment.Comment, user *User) *Comment {
	return &Comment{
		Id:         kitex_comment.Id,
		User:       user,
		Content:    kitex_comment.Content,
		CreateDate: kitex_comment.CreateDate,
	}
}

func buildCommentUserInfo(kitex_user *comment.User) *User {
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
