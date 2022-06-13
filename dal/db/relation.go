package db

import (
	"context"

	"github.com/a76yyyy/tiktok/pkg/errno"
	"gorm.io/gorm"
)

type Relation struct {
	gorm.Model
	User     User `gorm:"foreignkey:UserID;"`
	UserID   int  `gorm:"index:idx_userid,unique;not null"`
	ToUser   User `gorm:"foreignkey:ToUserID;"`
	ToUserID int  `gorm:"index:idx_userid,unique;index:idx_userid_to;not null"`
}

func (Relation) TableName() string {
	return "relation"
}

// get relation info
func GetRelation(ctx context.Context, uid int64, tid int64) (*Relation, error) {
	relation := new(Relation)

	if err := DB.WithContext(ctx).First(&relation, "user_id = ? and to_user_id = ?", uid, tid).Error; err != nil {
		return relation, err
	}
	return relation, nil
}

func NewRelation(ctx context.Context, uid int64, tid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		//1. 新增关注数据
		err := tx.Create(&Relation{UserID: int(uid), ToUserID: int(tid)}).Error
		if err != nil {
			return err
		}

		//2.改变 user 表中的 following count
		res := tx.Model(new(User)).Where("ID = ?", uid).Update("following_count", gorm.Expr("following_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.ErrDatabase
		}

		//3.改变 user 表中的 follower count
		res = tx.Model(new(User)).Where("ID = ?", tid).Update("follower_count", gorm.Expr("follower_count + ?", 1))
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

func DisRelation(ctx context.Context, uid int64, tid int64) error {
	relation := Relation{}
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		tx.Where("user_id = ? AND to_user_id=?", uid, tid).First(&relation)
		//1. 删除关注数据
		err := tx.Unscoped().Delete(&relation).Error
		if err != nil {
			return err
		}
		//2.改变 user 表中的 following count
		res := tx.Model(new(User)).Where("ID = ?", uid).Update("following_count", gorm.Expr("following_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.ErrDatabase
		}

		//3.改变 user 表中的 follower count
		res = tx.Model(new(User)).Where("ID = ?", tid).Update("follower_count", gorm.Expr("follower_count - ?", 1))
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

func FollowingList(ctx context.Context, uid int64) ([]*Relation, error) {
	var RelationList []*Relation
	err := DB.WithContext(ctx).Where("user_id = ?", uid).Find(&RelationList).Error
	if err != nil {
		return nil, err
	}
	return RelationList, nil
}

func FollowerList(ctx context.Context, tid int64) ([]*Relation, error) {
	var RelationList []*Relation
	err := DB.WithContext(ctx).Where("to_user_id = ?", tid).Find(&RelationList).Error
	if err != nil {
		return nil, err
	}
	return RelationList, nil
}
