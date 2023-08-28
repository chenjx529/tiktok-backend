package pack

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"tiktok-backend/kitex_gen/comment"
	"tiktok-backend/pkg/errno"
)

func SendCommentActionResponse(c *app.RequestContext, err error, commentId int64) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, CommentActionResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		CommentId:  commentId,
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
	for _, comment := range commentData {
		commentList = append(commentList, buildCommentInfo(comment))
	}
	return commentList
}

func buildCommentInfo(kitex_comment *comment.Comment) *Comment {
	user := &User{
		Id:              kitex_comment.User.Id,
		Name:            kitex_comment.User.Name,
		FollowCount:     kitex_comment.User.FollowCount,
		FollowerCount:   kitex_comment.User.FollowerCount,
		IsFollow:        kitex_comment.User.IsFollow,
		Avatar:          kitex_comment.User.Avatar,
		BackgroundImage: kitex_comment.User.BackgroundImage,
		Signature:       kitex_comment.User.Signature,
		TotalFavorited:  kitex_comment.User.TotalFavorited,
		WorkCount:       kitex_comment.User.WorkCount,
		FavoriteCount:   kitex_comment.User.FavoriteCount,
	}
	return &Comment{
		Id:         kitex_comment.Id,
		User:       user,
		Content:    kitex_comment.Content,
		CreateDate: kitex_comment.CreateDate,
	}
}
