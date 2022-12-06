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
 * @Date: 2022-06-12 18:08:54
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:05:59
 * @FilePath: /tiktok/cmd/relation/command/relation_list.go
 * @Description: 获取关注列表 操作业务逻辑
 */

package command

import (
	"context"

	"github.com/a76yyyy/tiktok/dal/db"
	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/relation"
	"github.com/a76yyyy/tiktok/kitex_gen/user"
)

type FollowingListService struct {
	ctx context.Context
}

// NewFollowingListService creates a new FollowingListService
func NewFollowingListService(ctx context.Context) *FollowingListService {
	return &FollowingListService{
		ctx: ctx,
	}
}

// FollowingList returns the following lists
func (s *FollowingListService) FollowingList(req *relation.DouyinRelationFollowListRequest, fromID int64) ([]*user.User, error) {
	FollowingUser, err := db.FollowingList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.FollowingList(s.ctx, FollowingUser, fromID)
}

type FollowerListService struct {
	ctx context.Context
}

// NewFollowerListService creates a new FollowerListService
func NewFollowerListService(ctx context.Context) *FollowerListService {
	return &FollowerListService{
		ctx: ctx,
	}
}

// FollowerList returns the Follower Lists
func (s *FollowerListService) FollowerList(req *relation.DouyinRelationFollowerListRequest, fromID int64) ([]*user.User, error) {
	FollowerUser, err := db.FollowerList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return pack.FollowerList(s.ctx, FollowerUser, fromID)
}
