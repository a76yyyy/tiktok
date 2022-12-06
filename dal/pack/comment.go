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
 * @Date: 2022-06-12 20:27:06
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:42:26
 * @FilePath: /tiktok/dal/pack/comment.go
 * @Description: 封装 Comments 数据库数据为 RPC Server 端响应
 */

package pack

import (
	"context"
	"errors"

	"github.com/a76yyyy/tiktok/dal/db"
	"github.com/a76yyyy/tiktok/kitex_gen/comment"
	"gorm.io/gorm"
)

// Comment pack Comments info.
func Comments(ctx context.Context, vs []*db.Comment, fromID int64) ([]*comment.Comment, error) {
	comments := make([]*comment.Comment, 0)
	for _, v := range vs {
		user, err := db.GetUserByID(ctx, int64(v.UserID))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		packUser, err := User(ctx, user, fromID)
		if err != nil {
			return nil, err
		}

		comments = append(comments, &comment.Comment{
			Id:         int64(v.ID),
			User:       packUser,
			Content:    v.Content,
			CreateDate: v.CreatedAt.Format("01-02"),
		})
	}
	return comments, nil
}
