package main

import (
	"context"

	"github.com/a76yyyy/tiktok/cmd/publish/command"
	"github.com/a76yyyy/tiktok/kitex_gen/publish"
	"github.com/a76yyyy/tiktok/pack"
	"github.com/a76yyyy/tiktok/pkg/errno"
)

// PublishSrvImpl implements the last service interface defined in the IDL.
type PublishSrvImpl struct{}

// PublishAction implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) PublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {

	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildPublishResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if len(req.Data) == 0 || len(req.Title) == 0 {
		resp = pack.BuildPublishResp(errno.ErrBind)
		return resp, nil
	}

	err = command.NewPublishActionService(ctx).PublishAction(req, int(claim.Id), &Config)
	if err != nil {
		resp = pack.BuildPublishResp(err)
		return resp, nil
	}
	resp = pack.BuildPublishResp(errno.Success)
	return resp, nil
}

// PublishList implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) PublishList(ctx context.Context, req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {

	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildPublishListResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.UserId == 0 {
		req.UserId = claim.Id // 没有传入UserID，默认为自己
	}

	videos, err := command.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		resp = pack.BuildPublishListResp(err)
		return resp, nil
	}

	resp = pack.BuildPublishListResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}
