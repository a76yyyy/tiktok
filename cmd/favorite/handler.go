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
 * @Date: 2022-06-12 00:02:35
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-18 23:58:11
 * @FilePath: /tiktok/cmd/favorite/handler.go
 * @Description: 定义 Favorite RPC Server 端的相关接口
 */

package main

import (
	"context"

	"github.com/a76yyyy/tiktok/cmd/favorite/command"
	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/favorite"
	"github.com/a76yyyy/tiktok/pkg/errno"
)

// FavoriteSrvImpl implements the last service interface defined in the IDL.
type FavoriteSrvImpl struct{}

// FavoriteAction implements the FavoriteSrvImpl interface.
func (s *FavoriteSrvImpl) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildFavoriteActionResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.UserId == 0 || claim.Id != 0 {
		req.UserId = claim.Id
	}

	if req.ActionType != 1 && req.ActionType != 2 || req.UserId == 0 || req.VideoId == 0 {
		resp = pack.BuildFavoriteActionResp(errno.ErrBind)
		return resp, nil
	}

	err = command.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp = pack.BuildFavoriteActionResp(err)
		return resp, nil
	}
	resp = pack.BuildFavoriteActionResp(err)
	return resp, nil
}

// FavoriteList implements the FavoriteSrvImpl interface.
func (s *FavoriteSrvImpl) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildFavoriteListResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.UserId == 0 || claim.Id == 0 {
		resp = pack.BuildFavoriteListResp(errno.ErrBind)
		return resp, nil
	}

	videos, err := command.NewFavoriteListService(ctx).FavoriteList(req)
	if err != nil {
		resp = pack.BuildFavoriteListResp(err)
		return resp, nil
	}
	resp = pack.BuildFavoriteListResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}
