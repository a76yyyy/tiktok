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
 * @Date: 2022-06-12 17:53:07
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:42:13
 * @FilePath: /tiktok/dal/pack/favorite.go
 * @Description: 封装 FavoriteVideos 数据库数据为 RPC Server 端响应
 */

package pack

import (
	"context"

	"github.com/a76yyyy/tiktok/dal/db"
	"github.com/a76yyyy/tiktok/kitex_gen/feed"
)

// FavoriteVideos pack favoriteVideos info.
func FavoriteVideos(ctx context.Context, vs []db.Video, uid *int64) ([]*feed.Video, error) {
	videos := make([]*db.Video, 0)
	for _, v := range vs {
		videos = append(videos, &v)
	}

	packVideos, err := Videos(ctx, videos, uid)
	if err != nil {
		return nil, err
	}

	return packVideos, nil
}
