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
 * @Date: 2022-06-10 14:47:34
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:01:30
 * @FilePath: /tiktok/cmd/user/handler.go
 * @Description: 定义 User RPC Server 端的相关接口
 */

package main

import (
	"context"

	"github.com/a76yyyy/tiktok/cmd/user/command"
	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/user"
	"github.com/a76yyyy/tiktok/pkg/errno"
	"github.com/a76yyyy/tiktok/pkg/jwt"
)

// UserSrvImpl implements the user service interface defined in the IDL.
type UserSrvImpl struct{}

// Register implements the UserSrvImpl interface.
func (s *UserSrvImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	if len(req.Username) == 0 || len(req.Password) == 0 {

		resp = pack.BuilduserRegisterResp(errno.ErrBind)
		return resp, nil
	}

	err = command.NewCreateUserService(ctx).CreateUser(req, Argon2Config)
	if err != nil {
		resp = pack.BuilduserRegisterResp(err)
		return resp, nil
	}

	// 新用户注册成功后直接登录
	uid, err := command.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp = pack.BuilduserRegisterResp(err)
		return resp, nil
	}

	token, err := Jwt.CreateToken(jwt.CustomClaims{
		Id: int64(uid),
	})
	if err != nil {
		resp = pack.BuilduserRegisterResp(errno.ErrSignatureInvalid)
		return resp, nil
	}

	resp = pack.BuilduserRegisterResp(errno.Success)
	resp.UserId = uid
	resp.Token = token
	return resp, nil
}

// Login implements the UserSrvImpl interface.
func (s *UserSrvImpl) Login(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuilduserRegisterResp(errno.ErrBind)
		return resp, nil
	}

	uid, err := command.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp = pack.BuilduserRegisterResp(err)
		return resp, nil
	}

	token, err := Jwt.CreateToken(jwt.CustomClaims{
		Id: int64(uid),
	})
	if err != nil {
		resp = pack.BuilduserRegisterResp(errno.ErrSignatureInvalid)
		return resp, nil
	}

	resp = pack.BuilduserRegisterResp(errno.Success)
	resp.UserId = uid
	resp.Token = token
	return resp, nil
}

// GetUserById implements the UserSrvImpl interface.
func (s *UserSrvImpl) GetUserById(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuilduserUserResp(errno.ErrTokenInvalid)
		return resp, nil
	}
	// else if claim.Id != int64(req.UserId) {
	// 	resp = pack.BuilduserUserResp(errno.ErrValidation)
	// 	return resp, nil
	// }

	if req.UserId < 0 {
		resp = pack.BuilduserUserResp(errno.ErrBind)
		return resp, nil
	}

	user, err := command.NewMGetUserService(ctx).MGetUser(req, claim.Id)
	if err != nil {
		resp = pack.BuilduserUserResp(err)
		return resp, nil
	}

	if claim.Id == req.UserId {
		user.IsFollow = true
	} else {
		// TODO 获取claim.id 是否已关注 req.userid
		user.IsFollow = false
	}

	resp = pack.BuilduserUserResp(errno.Success)
	resp.User = user
	return resp, nil
}
