package db

import (
	"context"

	"github.com/a76yyyy/tiktok/pkg/errno"
	"gorm.io/gorm"
)

type FavoriteRelation struct {
	gorm.Model
	User    User  `gorm:"foreignkey:UserID"`
	UserID  int   `gorm:"index:idx_userid_videoid,unique;not null"`
	Video   Video `gorm:"foreignkey:VideoID"`
	VideoID int   `gorm:"index:idx_userid_videoid,unique;index:idx_videoid;not null"`
}

func (FavoriteRelation) TableName() string {
	return "favorite_relation"
}

type FavoriteVideo struct {
	Video Video `gorm:"foreignkey:VideoID"`
}

// MGetVideoss multiple get list of videos info
func GetFavoriteRelation(ctx context.Context, uid int64, vid int64) (*FavoriteRelation, error) {
	video := new(FavoriteRelation)

	if err := DB.WithContext(ctx).Where("user_id = ? and video_id = ?", uid, vid).First(&video).Error; err != nil {
		return nil, err
	}
	return video, nil
}

func Favorite(ctx context.Context, uid int64, vid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		//1. 新增点赞数据
		err := tx.Create(&FavoriteRelation{UserID: int(uid), VideoID: int(vid)}).Error
		if err != nil {
			return err
		}
		//2.改变 video 表中的 favorite count
		res := tx.Model(new(Video)).Where("ID = ?", vid).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
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

func DisFavorite(ctx context.Context, uid int64, vid int64) error {
	favoriteRelation := FavoriteRelation{}
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		tx.First(&favoriteRelation, "user_id = ? and video_id = ?", uid, vid)
		//1. 删除点赞数据
		err := tx.Unscoped().Delete(&favoriteRelation).Error
		if err != nil {
			return err
		}
		//2.改变 video 表中的 favorite count
		res := tx.Model(new(Video)).Where("ID = ?", vid).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
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

func FavoriteList(ctx context.Context, uid int64) ([]*FavoriteVideo, error) {
	var favoriteList []*FavoriteVideo
	err := DB.WithContext(ctx).Model(&FavoriteRelation{}).Where(&FavoriteRelation{UserID: int(uid)}).Find(&favoriteList).Error
	if err != nil {
		return nil, err
	}
	return favoriteList, nil
}
