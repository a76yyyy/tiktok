package command

import (
	"context"

	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/relation"
	"github.com/a76yyyy/tiktok/kitex_gen/user"

	"github.com/a76yyyy/tiktok/dal/db"
)

type FollowingListService struct {
	ctx context.Context
}

func NewFollowingListService(ctx context.Context) *FollowingListService {
	return &FollowingListService{
		ctx: ctx,
	}
}

func (s *FollowingListService) FollowingList(req *relation.DouyinRelationFollowListRequest) ([]*user.User, error) {
	FollowingUser, err := db.FollowingList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.FollowingList(s.ctx, FollowingUser), nil
}

type FollowerListService struct {
	ctx context.Context
}

func NewFollowerListService(ctx context.Context) *FollowerListService {
	return &FollowerListService{
		ctx: ctx,
	}
}

func (s *FollowerListService) FollowerList(req *relation.DouyinRelationFollowerListRequest) ([]*user.User, error) {
	FollowerUser, err := db.FollowerList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return pack.FollowerList(s.ctx, FollowerUser), nil
}
