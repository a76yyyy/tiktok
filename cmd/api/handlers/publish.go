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
 * @Date: 2022-06-12 22:09:45
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-18 23:32:57
 * @FilePath: /tiktok/cmd/api/handlers/publish.go
 * @Description: 定义 Publish API 的 handler
 */

package handlers

import (
	"bytes"
	"context"
	"io"
	"strconv"

	"github.com/a76yyyy/tiktok/cmd/api/rpc"
	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/publish"
	"github.com/a76yyyy/tiktok/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

// 传递 发布视频操作 的上下文至 Publish 服务的 RPC 客户端, 并获取相应的响应.
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var paramVar PublishActionParam
	token := c.PostForm("token")
	title := c.PostForm("title")

	fileHeader, err := c.Request.FormFile("data")
	if err != nil {
		SendResponse(c, pack.BuildPublishResp(errno.ErrDecodingFailed))
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		SendResponse(c, pack.BuildPublishResp(errno.ErrDecodingFailed))
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		SendResponse(c, pack.BuildPublishResp(err))
		return
	}

	paramVar.Token = token
	paramVar.Title = title

	resp, err := rpc.PublishAction(ctx, &publish.DouyinPublishActionRequest{
		Title: paramVar.Title,
		Token: paramVar.Token,
		Data:  buf.Bytes(),
	})
	if err != nil {
		SendResponse(c, pack.BuildPublishResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

// 传递 获取视频列表操作 的上下文至 Publish 服务的 RPC 客户端, 并获取相应的响应.
func PublishList(ctx context.Context, c *app.RequestContext) {
	var paramVar UserParam
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, pack.BuildPublishListResp(errno.ErrBind))
		return
	}
	paramVar.UserId = int64(userid)
	paramVar.Token = c.Query("token")

	if len(paramVar.Token) == 0 || paramVar.UserId < 0 {
		SendResponse(c, pack.BuildPublishListResp(errno.ErrBind))
		return
	}

	resp, err := rpc.PublishList(ctx, &publish.DouyinPublishListRequest{
		UserId: paramVar.UserId,
		Token:  paramVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildPublishListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
