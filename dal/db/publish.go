package db

import (
	"context"
)

func CreateVideo(ctx context.Context, video *Video) error {
	return DB.WithContext(ctx).Create(video).Error
}

func PublishList(ctx context.Context, authorId int64) ([]*Video, error) {
	var pubList []*Video
	err := DB.WithContext(ctx).Model(&Video{}).Where(&Video{AuthorID: int(authorId)}).Find(&pubList).Error
	if err != nil {
		return nil, err
	}
	return pubList, nil
}
