/*
 * @Author: a76yyyy q981331502@163.com
 * @Date: 2022-06-11 01:01:53
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:08:03
 * @FilePath: /tiktok/cmd/user/command/getUserByID.go
 * @Description: 获取用户信息 操作业务逻辑
 */

package command

import (
	"context"
	"errors"

	"github.com/a76yyyy/tiktok/kitex_gen/user"
	"gorm.io/gorm"

	"github.com/a76yyyy/tiktok/dal/db"
	"github.com/a76yyyy/tiktok/dal/pack"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser get user info by userID
func (s *MGetUserService) MGetUser(req *user.DouyinUserRequest, fromID int64) (*user.User, error) {
	modelUser, err := db.GetUserByID(s.ctx, req.UserId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	user, err := pack.User(s.ctx, modelUser, fromID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
