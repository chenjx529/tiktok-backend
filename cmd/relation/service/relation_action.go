package service

import (
	"context"
	"tiktok-backend/kitex_gen/relation"
)

type RelationActionService struct {
	ctx context.Context
}

// NewRelationActionService new RelationActionService
func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

func (s *RelationActionService) RelationAction(req *relation.DouyinRelationActionRequest) error {
	return nil
}