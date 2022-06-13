package command

import (
	"context"

	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/comment"

	"github.com/a76yyyy/tiktok/dal/db"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{
		ctx: ctx,
	}
}

func (s *CommentListService) CommentList(req *comment.DouyinCommentListRequest) ([]*comment.Comment, error) {
	Comments, err := db.GetVideoComments(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	comments, err := pack.Comments(Comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
