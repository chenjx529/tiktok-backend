package service

import (
	"context"
	"tiktok-backend/cmd/comment/pack"
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/comment"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/jwt"
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
	// 登录id
	claims, err := jwt.GetClaimsFromTokenStr(req.Token)
	if err != nil {
		return nil, err
	}
	loginId := int64(claims[constants.IdentityKey].(float64))

	// 根据videoId获取所有的评论
	commentDate, err := db.QueryCommentByVideoId(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	// 根据commentDate获取所有的userId
	commentUserIds := make([]int64, 0)
	for _, com := range commentDate {
		commentUserIds = append(commentUserIds, com.UserId)  // 点赞视频的id
	}

	// 利用commentUserIds获取所有的commentUser
	commentUsers, err := db.MQueryUsersByIds(s.ctx, commentUserIds)
	if err != nil {
		return nil, err
	}
	userMap := make(map[int64]*db.User)
	for _, user := range commentUsers {
		userMap[int64(user.ID)] = user
	}

	// 获取用户关注
	var followSet map[int64]struct{}
	if loginId != 0 {
		if followSet, err = db.MQueryFollowByUserIdAndToUserIds(s.ctx, loginId, commentUserIds); err != nil {
			return nil, err
		}
	}

	// 封装commentList
	commentList := pack.BuildCommentList(loginId, commentDate, userMap, followSet)
	return commentList, nil
}
