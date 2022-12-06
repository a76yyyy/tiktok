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
 * @Date: 2022-06-12 17:02:26
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-18 23:57:17
 * @FilePath: /tiktok/cmd/favorite/command/favorite_list.go
 * @Description: 获取点赞列表 操作业务逻辑
 */

package command

import (
	"context"

	"github.com/a76yyyy/tiktok/dal/db"
	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/favorite"
	"github.com/a76yyyy/tiktok/kitex_gen/feed"
)

type FavoriteListService struct {
	ctx context.Context
}

// NewFavoriteListService creates a new FavoriteListService
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{
		ctx: ctx,
	}
}

// FavoriteList returns a Favorite List
func (s *FavoriteListService) FavoriteList(req *favorite.DouyinFavoriteListRequest) ([]*feed.Video, error) {
	FavoriteVideos, err := db.FavoriteList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	videos, err := pack.FavoriteVideos(s.ctx, FavoriteVideos, &req.UserId)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
