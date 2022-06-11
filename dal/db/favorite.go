package db

import (
	"context"

	"gorm.io/gorm"
)

type FavoriteVideo struct {
	gorm.Model
	User    User  `gorm:"foreignkey:UserID"`
	UserID  int   `gorm:"index,unique;not null"`
	Video   Video `gorm:"foreignkey:VideoID"`
	VideoID int   `gorm:"index,unique;not null"`
}

// MGetVideoss multiple get list of videos info
func GetFavoriteVideo(ctx context.Context, uid int, vid int) (FavoriteVideo, error) {
	video := FavoriteVideo{}

	if err := DB.WithContext(ctx).First(video, "user_id = ? and video_id = ?", uid, vid).Error; err != nil {
		return video, err
	}
	return video, nil
}
