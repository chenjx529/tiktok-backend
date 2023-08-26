package service

import (
	"context"
	"tiktok-backend/kitex_gen/publish"
)

type PublishActionService struct {
	ctx context.Context
}

// NewPublishActionService new PublishActionService
func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{
		ctx: ctx,
	}
}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishActionService) PublishAction(req *publish.DouyinPublishActionRequest) error {
	// TODO: Your code here...
	return nil
}

