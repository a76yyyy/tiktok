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
 * @Date: 2022-06-11 16:59:09
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:00:05
 * @FilePath: /tiktok/cmd/feed/handler.go
 * @Description: 定义 Feed RPC Server 端的相关接口
 */

package main

import (
	"context"

	"github.com/a76yyyy/tiktok/cmd/feed/command"
	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/feed"
	"github.com/a76yyyy/tiktok/pkg/errno"
)

// FeedSrvImpl implements the last service interface defined in the IDL.
type FeedSrvImpl struct{}

// GetUserFeed implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) GetUserFeed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	var uid int64 = 0
	if *req.Token != "" {
		claim, err := Jwt.ParseToken(*req.Token)
		if err != nil {
			resp = pack.BuildVideoResp(errno.ErrTokenInvalid)
			return resp, nil
		} else {
			uid = claim.Id
		}
	}

	vis, nextTime, err := command.NewGetUserFeedService(ctx).GetUserFeed(req, uid)
	if err != nil {
		resp = pack.BuildVideoResp(err)
		return resp, nil
	}

	resp = pack.BuildVideoResp(errno.Success)
	resp.VideoList = vis
	resp.NextTime = &nextTime
	return resp, nil
}

// GetVideoById implements the FeedSrvImpl interface.
// Deprecated: Never Use GetVideoById
func (s *FeedSrvImpl) GetVideoById(ctx context.Context, req *feed.VideoIdRequest) (resp *feed.Video, err error) {
	return
}
