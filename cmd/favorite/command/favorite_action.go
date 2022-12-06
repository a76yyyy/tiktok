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
 * @Date: 2022-06-12 16:03:40
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-18 23:55:36
 * @FilePath: /tiktok/cmd/favorite/command/favorite_action.go
 * @Description: 点赞 操作业务逻辑
 */

package command

import (
	"context"

	"github.com/a76yyyy/tiktok/dal/db"
	"github.com/a76yyyy/tiktok/kitex_gen/favorite"
	"github.com/a76yyyy/tiktok/pkg/errno"
)

type FavoriteActionService struct {
	ctx context.Context
}

// NewFavoriteActionService new FavoriteActionService
func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

// FavoriteAction action favorite.
func (s *FavoriteActionService) FavoriteAction(req *favorite.DouyinFavoriteActionRequest) error {
	// 1-点赞
	if req.ActionType == 1 {
		return db.Favorite(s.ctx, req.UserId, req.VideoId)
	}
	// 2-取消点赞
	if req.ActionType == 2 {
		return db.DisFavorite(s.ctx, req.UserId, req.VideoId)
	}
	return errno.ErrBind
}
