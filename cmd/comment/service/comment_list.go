package service

import (
	"context"
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/comment"
)

type CommentListService struct {
	ctx context.Context
}

// NewCommentListService new CommentListService
func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{
		ctx: ctx,
	}
}

// CommentList 获取评论列表
func (s *CommentListService) CommentList(req *comment.DouyinCommentListRequest) ([]*comment.Comment, error) {
	commentData, err := db.CommentList(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	var commentList []*comment.Comment
	for _, rawComment := range commentData {
		user, err := db.QueryUserById(s.ctx, rawComment.UserId)
		if err != nil {
			return nil, err
		}
		firstUser := user[0]
		commentUser := &comment.User{
			Id:              int64(firstUser.ID),
			Name:            firstUser.Name,
			FollowCount:     firstUser.FollowCount,
			FollowerCount:   firstUser.FollowerCount,
			Avatar:          firstUser.Avatar,
			BackgroundImage: firstUser.BackgroundImage,
			Signature:       firstUser.Signature,
			TotalFavorited:  firstUser.TotalFavorited,
			WorkCount:       firstUser.WorkCount,
			FavoriteCount:   firstUser.FavoriteCount,
		}
		newComment := &comment.Comment{
			Id:         int64(rawComment.ID),
			User:       commentUser,
			Content:    rawComment.Contents,
			CreateDate: rawComment.CreatedAt.String(),
		}
		commentList = append(commentList, newComment)
	}
	return commentList, nil
}
