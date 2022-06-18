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
 * @Date: 2022-06-11 23:52:53
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:01:03
 * @FilePath: /tiktok/cmd/relation/handler.go
 * @Description: 定义 Relation RPC Server 端的相关接口
 */

package main

import (
	"context"

	"github.com/a76yyyy/tiktok/cmd/relation/command"
	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/relation"
	"github.com/a76yyyy/tiktok/pkg/errno"
)

// RelationSrvImpl implements the last service interface defined in the IDL.
type RelationSrvImpl struct{}

// RelationAction implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildRelationActionResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.UserId == 0 || claim.Id != 0 {
		req.UserId = claim.Id
	}

	if req.ActionType < 1 || req.ActionType > 2 {
		resp = pack.BuildRelationActionResp(errno.ErrBind)
		return resp, nil
	}
	err = command.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		resp = pack.BuildRelationActionResp(err)
		return resp, nil
	}
	resp = pack.BuildRelationActionResp(errno.Success)
	return resp, nil
}

// RelationFollowList implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildFollowingListResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.UserId == 0 || claim.Id != 0 {
		req.UserId = claim.Id // 没有传入UserID，默认为自己
	}

	users, err := command.NewFollowingListService(ctx).FollowingList(req, claim.Id)
	if err != nil {
		resp = pack.BuildFollowingListResp(err)
		return resp, nil
	}

	resp = pack.BuildFollowingListResp(errno.Success)
	resp.UserList = users
	return resp, nil
}

// RelationFollowerList implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildFollowerListResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.UserId == 0 {
		req.UserId = claim.Id // 没有传入UserID，默认为自己
	}

	users, err := command.NewFollowerListService(ctx).FollowerList(req, claim.Id)
	if err != nil {
		resp = pack.BuildFollowerListResp(err)
		return resp, nil
	}

	resp = pack.BuildFollowerListResp(errno.Success)
	resp.UserList = users
	return resp, nil
}
