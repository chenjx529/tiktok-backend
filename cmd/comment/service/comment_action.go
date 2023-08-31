package service

import (
	"context"
	"errors"
	"sync"
	"tiktok-backend/cmd/comment/pack"
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

func (s *CommentActionService) CommentAction(req *comment.DouyinCommentActionRequest) (*comment.Comment, error) {
	// 登录id
	claims, err := jwt.GetClaimsFromTokenStr(req.Token)
	if err != nil {
		return nil, err
	}
	loginId := int64(claims[constants.IdentityKey].(float64))

	// 添加评论
	if req.ActionType == 1 {

		commentData := &db.Comment{
			UserId:   loginId,
			VideoId:  req.VideoId,
			Contents: req.CommentText,
		}

		var wg sync.WaitGroup
		wg.Add(2)
		var user *db.User
		var commentErr, userErr error

		// 添加记录
		go func() {
			defer wg.Done()
			if err := db.CreateComment(s.ctx, commentData); err != nil {
				commentErr = err
				return
			}
		}()

		//获取当前用户信息
		go func() {
			defer wg.Done()
			users, err := db.MQueryUsersByIds(s.ctx, []int64{loginId})
			if err != nil {
				userErr = err
				return
			}
			user = users[0]
		}()

		wg.Wait()
		if commentErr != nil {
			return nil, commentErr
		}
		if userErr != nil {
			return nil, userErr
		}

		return pack.BuildCommentInfo(commentData, user, false), nil

	}

	// 删除评论
	if req.ActionType == 2 {
		var wg sync.WaitGroup
		wg.Add(2)
		var user *db.User
		var commentData *db.Comment
		var commentErr, userErr error

		// 删除评论
		go func() {
			defer wg.Done()
			if commentData, err = db.DeleteComment(s.ctx, req.CommentId, req.VideoId); err != nil {
				commentErr = err
				return
			}
		}()

		// 获取当前用户信息
		go func() {
			defer wg.Done()
			users, err := db.MQueryUsersByIds(s.ctx, []int64{loginId})
			if err != nil {
				userErr = err
				return
			}
			user = users[0]
		}()

		wg.Wait()
		if commentErr != nil {
			return nil, commentErr
		}
		if userErr != nil {
			return nil, userErr
		}

		return pack.BuildCommentInfo(commentData, user, false), nil
	}

	return nil, errors.New("ActionType error")
}
