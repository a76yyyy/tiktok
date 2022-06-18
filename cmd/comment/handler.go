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
 * @Date: 2022-06-12 00:03:00
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-18 23:52:17
 * @FilePath: /tiktok/cmd/comment/handler.go
 * @Description: 定义 Comment RPC Server 端的相关接口
 */

package main

import (
	"context"

	"github.com/a76yyyy/tiktok/cmd/comment/command"
	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/comment"
	"github.com/a76yyyy/tiktok/pkg/errno"
)

// CommentSrvImpl implements the last service interface defined in the IDL.
type CommentSrvImpl struct{}

// CommentAction implements the CommentSrvImpl interface.
func (s *CommentSrvImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildCommentActionResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.UserId == 0 || claim.Id != 0 {
		req.UserId = claim.Id
	}

	if req.ActionType != 1 && req.ActionType != 2 || req.UserId == 0 || req.VideoId == 0 {
		resp = pack.BuildCommentActionResp(errno.ErrBind)
		return resp, nil
	}

	err = command.NewCommentActionService(ctx).CommentAction(req)
	if err != nil {
		resp = pack.BuildCommentActionResp(err)
		return resp, nil
	}
	resp = pack.BuildCommentActionResp(err)
	return resp, nil
}

// CommentList implements the CommentSrvImpl interface.
func (s *CommentSrvImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildCommentListResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.VideoId == 0 || claim.Id == 0 {
		resp = pack.BuildCommentListResp(errno.ErrBind)
		return resp, nil
	}

	comments, err := command.NewCommentListService(ctx).CommentList(req, claim.Id)
	if err != nil {
		resp = pack.BuildCommentListResp(err)
		return resp, nil
	}
	resp = pack.BuildCommentListResp(errno.Success)
	resp.CommentList = comments
	return resp, nil
}
