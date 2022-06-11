package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	// UpdatedAt     time.Time `gorm:"column:update_time;not null;index:idx_update" `
	Author        User `gorm:"ForeignKey:AuthorID;AssociationForeignKey:ID"`
	AuthorID      int
	PlayUrl       string `gorm:"type:varchar(255);not null"`
	CoverUrl      string `gorm:"type:varchar(255)"`
	FavoriteCount int    `gorm:"default:0"`
	CommentCount  int    `gorm:"default:0"`
	Title         string `gorm:"type:varchar(50);not null"`
}

// MGetVideoss multiple get list of videos info
func MGetVideos(ctx context.Context, latestTime int64) ([]*Video, error) {
	videos := make([]*Video, 0)

	if err := DB.WithContext(ctx).Limit(30).Order("update_time desc").Find(&videos, "update_time < ?", time.UnixMilli(latestTime)).Error; err != nil {
		return nil, err
	}
	return videos, nil
}
