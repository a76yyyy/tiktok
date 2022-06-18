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
 * @Date: 2022-06-11 22:57:06
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:37:39
 * @FilePath: /tiktok/dal/db/favorite.go
 * @Description: Favorite 数据库操作业务逻辑
 */

package db

import (
	"context"

	"github.com/a76yyyy/tiktok/pkg/errno"
	"gorm.io/gorm"
)

// GetFavoriteRelation get favorite video info
func GetFavoriteRelation(ctx context.Context, uid int64, vid int64) (*Video, error) {
	user := new(User)
	if err := DB.WithContext(ctx).First(user, uid).Error; err != nil {
		return nil, err
	}

	video := new(Video)
	// if err := DB.WithContext(ctx).First(&video, vid).Error; err != nil {
	// 	return nil, err
	// }

	if err := DB.WithContext(ctx).Model(&user).Association("FavoriteVideos").Find(&video, vid); err != nil {
		return nil, err
	}
	return video, nil
}

// Favorite new favorite data.
func Favorite(ctx context.Context, uid int64, vid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		//1. 新增点赞数据
		user := new(User)
		if err := tx.WithContext(ctx).First(user, uid).Error; err != nil {
			return err
		}

		video := new(Video)
		if err := tx.WithContext(ctx).First(video, vid).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Model(&user).Association("FavoriteVideos").Append(video); err != nil {
			return err
		}
		//2.改变 video 表中的 favorite count
		res := tx.Model(video).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.ErrDatabase
		}

		return nil
	})
	return err
}

// DisFavorite deletes the specified favorite from the database
func DisFavorite(ctx context.Context, uid int64, vid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		//1. 删除点赞数据
		user := new(User)
		if err := tx.WithContext(ctx).First(user, uid).Error; err != nil {
			return err
		}

		video, err := GetFavoriteRelation(ctx, uid, vid)
		if err != nil {
			return err
		}

		err = tx.Unscoped().WithContext(ctx).Model(&user).Association("FavoriteVideos").Delete(video)
		if err != nil {
			return err
		}

		//2.改变 video 表中的 favorite count
		res := tx.Model(video).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.ErrDatabase
		}

		return nil
	})
	return err
}

// FavoriteList returns a list of Favorite videos.
func FavoriteList(ctx context.Context, uid int64) ([]Video, error) {
	user := new(User)
	if err := DB.WithContext(ctx).First(user, uid).Error; err != nil {
		return nil, err
	}

	videos := []Video{}
	// if err := DB.WithContext(ctx).First(&video, vid).Error; err != nil {
	// 	return nil, err
	// }

	if err := DB.WithContext(ctx).Model(&user).Association("FavoriteVideos").Find(&videos); err != nil {
		return nil, err
	}
	return videos, nil
}
