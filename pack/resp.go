// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package pack

import (
	"errors"

	"github.com/a76yyyy/tiktok/pkg/errno"

	"github.com/a76yyyy/tiktok/kitex_gen/feed"
	"github.com/a76yyyy/tiktok/kitex_gen/user"
)

// BuilduserRegisterResp build userRegisterResp from error
func BuilduserRegisterResp(err error) *user.DouyinUserRegisterResponse {
	if err == nil {
		return userRegisterResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userRegisterResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return userRegisterResp(s)
}

func userRegisterResp(err errno.ErrNo) *user.DouyinUserRegisterResponse {
	return &user.DouyinUserRegisterResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// BuilduserResp build userResp from error
func BuilduserUserResp(err error) *user.DouyinUserResponse {
	if err == nil {
		return userResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return userResp(s)
}

func userResp(err errno.ErrNo) *user.DouyinUserResponse {
	return &user.DouyinUserResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// BuildVideoResp build VideoResp from error
func BuildVideoResp(err error) *feed.DouyinFeedResponse {
	if err == nil {
		return videoResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return videoResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return videoResp(s)
}

func videoResp(err errno.ErrNo) *feed.DouyinFeedResponse {
	return &feed.DouyinFeedResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}
