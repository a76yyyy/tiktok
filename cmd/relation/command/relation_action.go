/*
 * @Author: a76yyyy q981331502@163.com
 * @Date: 2022-06-12 18:08:29
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:04:11
 * @FilePath: /tiktok/cmd/relation/command/relation_action.go
 * @Description: 关注用户 操作业务逻辑
 */

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
