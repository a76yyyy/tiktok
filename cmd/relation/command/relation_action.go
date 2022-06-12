package command

import (
	"context"

	"github.com/a76yyyy/tiktok/kitex_gen/relation"
	"github.com/a76yyyy/tiktok/pkg/errno"

	"github.com/a76yyyy/tiktok/dal/db"
)

type RelationActionService struct {
	ctx context.Context
}

// NewRelationActionService new RelationActionService
func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

// RelationAction action favorite.
func (s *RelationActionService) RelationAction(req *relation.DouyinRelationActionRequest) error {
	// 1-关注
	if req.ActionType == 1 {
		return db.NewRelation(s.ctx, req.UserId, req.ToUserId)
	}
	// 2-取消关注
	if req.ActionType == 2 {
		return db.DisRelation(s.ctx, req.UserId, req.ToUserId)
	}
	return errno.ErrBind
}
