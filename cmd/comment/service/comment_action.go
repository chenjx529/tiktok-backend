package service

import (
	"context"
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/comment"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/jwt"
)

type CommentActionService struct {
	ctx context.Context
}

// NewCommentActionService new CommentActionService
func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{
		ctx: ctx,
	}
}

func (s *CommentActionService) CommentAction(req *comment.DouyinCommentActionRequest) (int64, error) {
	// 是否登录
	claims, err := jwt.GetclaimsFromTokenStr(req.Token)
	var login_id int64
	if err != nil {
		login_id = 0
	} else {
		login_id = int64(int(claims[constants.IdentityKey].(float64))) // 这种写法，我是真的想骂人的
	}

	// 进行评论操作
	commentId, err := db.CommentAction(s.ctx, &db.Comment{
		UserId:   login_id,
		VideoId:  req.VideoId,
		Contents: req.CommentText,
	})

	return commentId, nil
}
