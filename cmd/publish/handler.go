// Copyright 2022 a76yyyy && CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * @Author: a76yyyy q981331502@163.com
 * @Date: 2022-06-12 00:00:46
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:00:44
 * @FilePath: /tiktok/cmd/publish/handler.go
 * @Description: 定义 Publish RPC Server 端的相关接口
 */

package main

import (
	"context"

	"github.com/a76yyyy/tiktok/cmd/publish/command"
	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/publish"
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
