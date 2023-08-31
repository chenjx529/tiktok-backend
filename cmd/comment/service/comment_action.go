package service

import (
	"context"
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

func (s *CommentActionService) CommentAction(req *comment.DouyinCommentActionRequest) (*comment.Comment, error) {
	// 登录id
	claims, err := jwt.GetClaimsFromTokenStr(req.Token)
	if err != nil {
		return nil, err
	}
	loginId := int64(claims[constants.IdentityKey].(float64))

	// 找到评论的用户，看loginId有没有关注他



}
