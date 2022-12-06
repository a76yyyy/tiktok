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
 * @Date: 2022-06-12 19:27:00
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-18 23:57:04
 * @FilePath: /tiktok/cmd/comment/command/command_list.go
 * @Description: 获取评论列表 操作业务逻辑
 */

package command

import (
	"context"

	"github.com/a76yyyy/tiktok/dal/db"
	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/comment"
)

type CommentListService struct {
	ctx context.Context
}

// NewCommentActionService new CommentActionService
func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{
		ctx: ctx,
	}
}

// CommentList return comment list
func (s *CommentListService) CommentList(req *comment.DouyinCommentListRequest, fromID int64) ([]*comment.Comment, error) {
	Comments, err := db.GetVideoComments(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	comments, err := pack.Comments(s.ctx, Comments, fromID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
