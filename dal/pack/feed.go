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
 * @Date: 2022-06-11 19:25:03
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:41:22
 * @FilePath: /tiktok/dal/pack/feed.go
 * @Description: 封装 Videos 数据库数据为 RPC Server 端响应
 */

package pack

import (
	"context"
	"errors"

	"github.com/a76yyyy/tiktok/dal/db"
	"github.com/a76yyyy/tiktok/kitex_gen/feed"
	"gorm.io/gorm"
)

// Video pack feed info
func Video(ctx context.Context, v *db.Video, fromID int64) (*feed.Video, error) {
	if v == nil {
		return nil, nil
	}
	user, err := db.GetUserByID(ctx, int64(v.AuthorID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	author, err := User(ctx, user, fromID)
	if err != nil {
		return nil, err
	}
	favorite_count := int64(v.FavoriteCount)
	comment_count := int64(v.CommentCount)

	return &feed.Video{
		Id:            int64(v.ID),
		Author:        author,
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: favorite_count,
		CommentCount:  comment_count,
		Title:         v.Title,
	}, nil
}

// Videos pack list of video info
func Videos(ctx context.Context, vs []*db.Video, fromID *int64) ([]*feed.Video, error) {
	videos := make([]*feed.Video, 0)
	for _, v := range vs {
		video2, err := Video(ctx, v, *fromID)
		if err != nil {
			return nil, err
		}

		if video2 != nil {
			flag := false
			if *fromID != 0 {
				results, err := db.GetFavoriteRelation(ctx, *fromID, int64(v.ID))
				if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, err
				} else if errors.Is(err, gorm.ErrRecordNotFound) {
					flag = false
				} else if results != nil && results.AuthorID != 0 {
					flag = true
				}
			}
			video2.IsFavorite = flag
			videos = append(videos, video2)
		}
	}
	return videos, nil
}
