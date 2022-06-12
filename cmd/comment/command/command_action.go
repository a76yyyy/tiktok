package command

import (
	"context"

	"github.com/a76yyyy/tiktok/kitex_gen/comment"
	"github.com/a76yyyy/tiktok/pkg/errno"

	"github.com/a76yyyy/tiktok/dal/db"
)

type CommentActionService struct {
	ctx context.Context
}

// NewFavoriteActionService new FavoriteActionService
func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

// CommentAction action comment.
func (s *CommentActionService) CommentAction(req *comment.DouyinCommentActionRequest) error {
	// 1-评论
	if req.ActionType == 1 {
		return db.NewComment(s.ctx, &db.Comment{
			UserID:  int(req.UserId),
			VideoID: int(req.VideoId),
			Content: *req.CommentText,
		})
	}
	// 2-删除评论
	if req.ActionType == 2 {
		return db.DelComment(s.ctx, *req.CommentId, req.VideoId)
	}
	return errno.ErrBind
}
