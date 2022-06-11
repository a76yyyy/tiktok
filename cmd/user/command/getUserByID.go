package command

import (
	"context"

	"github.com/a76yyyy/tiktok/kitex_gen/user"

	"github.com/a76yyyy/tiktok/dal/db"
	"github.com/a76yyyy/tiktok/pack"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *MGetUserService) MGetUser(req *user.DouyinUserRequest) (*user.User, error) {
	modelUser, err := db.MGetUser(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.User(modelUser), nil
}
