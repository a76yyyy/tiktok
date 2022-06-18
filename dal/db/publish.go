/*
 * @Author: a76yyyy q981331502@163.com
 * @Date: 2022-06-12 12:41:13
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:35:35
 * @FilePath: /tiktok/dal/db/publish.go
 * @Description: Publish 数据库操作业务逻辑
 */

package db

import (
	"context"

	"gorm.io/gorm"
)

// CreateVideo creates a new video
func CreateVideo(ctx context.Context, video *Video) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(video).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// PublishList returns a list of videos with AuthorID.
func PublishList(ctx context.Context, authorId int64) ([]*Video, error) {
	var pubList []*Video
	err := DB.WithContext(ctx).Model(&Video{}).Where(&Video{AuthorID: int(authorId)}).Find(&pubList).Error
	if err != nil {
		return nil, err
	}
	return pubList, nil
}
