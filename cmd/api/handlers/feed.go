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
 * @Date: 2022-06-11 23:30:14
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-18 23:30:29
 * @FilePath: /tiktok/cmd/api/handlers/feed.go
 * @Description: 定义 Feed API 的 handler
 */

package handlers

import (
	"context"
	"strconv"

	"github.com/a76yyyy/tiktok/cmd/api/rpc"
	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/feed"
	"github.com/a76yyyy/tiktok/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

// 传递 获取用户视频流操作 的上下文至 Feed 服务的 RPC 客户端, 并获取相应的响应.
func GetUserFeed(ctx context.Context, c *app.RequestContext) {
	var feedVar FeedParam
	var laststTime int64
	var token string
	lastst_time := c.Query("latest_time")
	if len(lastst_time) != 0 {
		if latesttime, err := strconv.Atoi(lastst_time); err != nil {
			SendResponse(c, pack.BuildVideoResp(errno.ErrDecodingFailed))
			return
		} else {
			laststTime = int64(latesttime)
		}
	}

	feedVar.LatestTime = &laststTime

	token = c.Query("token")
	feedVar.Token = &token

	resp, err := rpc.GetUserFeed(ctx, &feed.DouyinFeedRequest{
		LatestTime: feedVar.LatestTime,
		Token:      feedVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildVideoResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
